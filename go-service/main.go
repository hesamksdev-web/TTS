package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	PYTHON_SERVICE_URL         = "http://python-service:5000"
	voiceCloneWorkerIdleSleep  = 2 * time.Second
	voiceCloneWorkerErrorSleep = 5 * time.Second
)

var (
	db         *gorm.DB
	httpClient = newHTTPClient()
	jwtSecret  string
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
	Role     string // "admin" or "user"
	Approved bool   `gorm:"default:false"` // Admin approval required
}

type VoiceCloneJob struct {
	gorm.Model
	UserID       uint
	Text         string
	Language     string
	Speed        float64
	Emotion      string
	ModelType    string
	Status       string
	InputPath    string
	OutputPath   string
	ErrorMessage string
}

type Notification struct {
	gorm.Model
	UserID  uint
	JobID   uint
	Message string
	Read    bool `gorm:"default:false"`
}

type contextKey string

const (
	ctxKeyUserID   contextKey = "userID"
	ctxKeyUserRole contextKey = "userRole"
)

const (
	jobStatusPending    = "pending"
	jobStatusProcessing = "processing"
	jobStatusCompleted  = "completed"
	jobStatusFailed     = "failed"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type synthesizeRequest struct {
	Text      string `json:"text"`
	SpeakerID string `json:"speaker_id"`
}

type synthesizeTrainedRequest struct {
	Text  string `json:"text"`
	JobID string `json:"job_id"`
}

type finetuneRequest struct {
	DatasetName string `json:"dataset_name"`
	Epochs      int    `json:"epochs"`
	BatchSize   int    `json:"batch_size"`
}

func initDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Waiting for Database...", err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&User{}, &VoiceCloneJob{}, &Notification{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	var count int64
	db.Model(&User{}).Where("email = ?", "admin@tts.com").Count(&count)
	if count == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 14)
		db.Create(&User{Email: "admin@tts.com", Password: string(hash), Role: "admin"})
		fmt.Println("Default Admin Created: admin@tts.com / admin123")
	}
}

type corsMiddleware struct {
	handler http.Handler
}

func (c *corsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	c.handler.ServeHTTP(w, r)
}

func main() {
	initDB()
	jwtSecret = getEnv("JWT_SECRET", "my_super_secret_key_2025")
	startVoiceCloneWorker()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/auth/register", handleRegister)
	mux.HandleFunc("/api/v1/auth/login", handleLogin)

	mux.HandleFunc("/api/v1/dataset/upload", authMiddleware(datasetUploadHandler))
	mux.HandleFunc("/api/v1/finetune/start", authMiddleware(finetuneHandler))
	mux.HandleFunc("/api/v1/synthesize", authMiddleware(synthesizeHandler))
	mux.HandleFunc("/api/v1/synthesize-trained", authMiddleware(synthesizeTrainedHandler))
	mux.HandleFunc("/api/v1/trained-models", authMiddleware(listTrainedModelsHandler))
	mux.HandleFunc("/api/v1/voice-clone", authMiddleware(voiceCloneHandler))
	mux.HandleFunc("/api/v1/voice-clone/jobs", authMiddleware(listVoiceCloneJobsHandler))
	mux.HandleFunc("/api/v1/voice-clone/job", authMiddleware(getVoiceCloneJobHandler))
	mux.HandleFunc("/api/v1/voice-clone/job/download", authMiddleware(downloadVoiceCloneJobHandler))
	mux.HandleFunc("/api/v1/notifications", authMiddleware(listNotificationsHandler))
	mux.HandleFunc("/api/v1/notifications/read", authMiddleware(markNotificationsReadHandler))

	mux.HandleFunc("/api/v1/admin/users", adminMiddleware(handleListUsers))
	mux.HandleFunc("/api/v1/admin/approve", adminMiddleware(handleApproveUser))
	mux.HandleFunc("/api/v1/admin/test", adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"message": "Admin access granted"})
	}))
	mux.HandleFunc("/api/v1/admin/training/upload", adminMiddleware(adminTrainingUploadHandler))
	mux.HandleFunc("/api/v1/admin/training/start", adminMiddleware(adminTrainingStartHandler))
	mux.HandleFunc("/api/v1/admin/training/datasets", adminMiddleware(adminListDatasetsHandler))
	mux.HandleFunc("/api/v1/admin/training/jobs", adminMiddleware(adminListTrainingJobsHandler))

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	handler := &corsMiddleware{handler: mux}

	port := getEnv("PORT", "8080")
	log.Printf("Go Service (SaaS Edition) listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	var req RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	user := User{Email: req.Email, Password: string(hashedPassword), Role: "user"}

	if result := db.Create(&user); result.Error != nil {
		writeError(w, http.StatusBadRequest, "User already exists")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"message": "User registered"})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	var user User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		writeError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		writeError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Check if user is approved
	if !user.Approved && user.Role != "admin" {
		writeError(w, http.StatusForbidden, "Account pending admin approval")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(jwtSecret))

	writeJSON(w, http.StatusOK, AuthResponse{Token: tokenString, Role: user.Role, Email: user.Email})
}

func handleListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	type UserResponse struct {
		ID        uint   `json:"id"`
		Email     string `json:"email"`
		Role      string `json:"role"`
		Approved  bool   `json:"approved"`
		CreatedAt string `json:"created_at"`
	}

	var users []User
	if result := db.Select("id, email, role, approved, created_at").Find(&users); result.Error != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	if users == nil {
		users = []User{}
	}

	// Convert to response format
	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Role:      user.Role,
			Approved:  user.Approved,
			CreatedAt: user.CreatedAt.Format("2006-01-02"),
		})
	}

	writeJSON(w, http.StatusOK, response)
}

func handleApproveUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	type ApproveRequest struct {
		UserID uint `json:"user_id"`
	}

	var req ApproveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if result := db.Model(&User{}).Where("id = ?", req.UserID).Update("approved", true); result.Error != nil {
		writeError(w, http.StatusInternalServerError, "Failed to approve user")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "User approved"})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rWithCtx, err := injectUserContextFromRequest(r)
		if err != nil {
			writeError(w, http.StatusUnauthorized, "Invalid or missing token")
			return
		}
		next(w, rWithCtx)
	}
}

func injectUserContextFromRequest(r *http.Request) (*http.Request, error) {
	authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	if authHeader == "" {
		return nil, errors.New("missing Authorization header")
	}

	tokenString := authHeader
	if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
		tokenString = strings.TrimSpace(tokenString[7:])
	}
	if tokenString == "" {
		return nil, errors.New("missing bearer token")
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	userIDClaim, ok := claims["sub"]
	if !ok {
		return nil, errors.New("token missing subject")
	}
	userID, err := claimValueToUint(userIDClaim)
	if err != nil {
		return nil, err
	}

	role, _ := claims["role"].(string)

	ctx := context.WithValue(r.Context(), ctxKeyUserID, userID)
	ctx = context.WithValue(ctx, ctxKeyUserRole, role)
	return r.WithContext(ctx), nil
}

func getUserIDFromContext(r *http.Request) (uint, error) {
	val := r.Context().Value(ctxKeyUserID)
	if id, ok := val.(uint); ok && id > 0 {
		return id, nil
	}
	return 0, errors.New("missing user context")
}

func getUserRoleFromContext(r *http.Request) (string, error) {
	val := r.Context().Value(ctxKeyUserRole)
	if role, ok := val.(string); ok && role != "" {
		return role, nil
	}
	return "", errors.New("missing role context")
}

func claimValueToUint(value any) (uint, error) {
	switch v := value.(type) {
	case float64:
		if v < 0 {
			return 0, errors.New("negative user id")
		}
		return uint(v), nil
	case int:
		if v < 0 {
			return 0, errors.New("negative user id")
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, errors.New("negative user id")
		}
		return uint(v), nil
	case json.Number:
		parsed, err := v.Int64()
		if err != nil || parsed < 0 {
			return 0, errors.New("invalid user id")
		}
		return uint(parsed), nil
	case string:
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, errors.New("invalid user id string")
		}
		return uint(parsed), nil
	default:
		return 0, errors.New("unsupported user id type")
	}
}

func adminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rWithCtx, err := injectUserContextFromRequest(r)
		if err != nil {
			writeError(w, http.StatusUnauthorized, "Invalid or missing token")
			return
		}

		role, err := getUserRoleFromContext(rWithCtx)
		if err != nil || role != "admin" {
			writeError(w, http.StatusForbidden, "Forbidden: Admin access required")
			return
		}

		next(w, rWithCtx)
	}
}

func synthesizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req synthesizeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if req.Text == "" || req.SpeakerID == "" {
		writeError(w, http.StatusBadRequest, "text and speaker_id are required")
		return
	}

	proxyJSON(w, PYTHON_SERVICE_URL+"/tts", req)
}

func synthesizeTrainedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req synthesizeTrainedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if req.Text == "" || req.JobID == "" {
		writeError(w, http.StatusBadRequest, "text and job_id are required")
		return
	}

	proxyJSON(w, PYTHON_SERVICE_URL+"/tts-trained", req)
}

func listTrainedModelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	proxyRequest, err := http.NewRequest(http.MethodGet, PYTHON_SERVICE_URL+"/trained-models", nil)
	if err != nil {
		log.Printf("failed to create request: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp, err := httpClient.Do(proxyRequest)
	if err != nil {
		log.Printf("python service request failed: %v", err)
		writeError(w, http.StatusBadGateway, "python service unavailable")
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if _, err := w.Write(respBody); err != nil {
		log.Printf("failed to write response body: %v", err)
	}
}

func voiceCloneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid user context")
		return
	}

	if err := r.ParseMultipartForm(100 << 20); err != nil {
		writeError(w, http.StatusBadRequest, "failed to parse form data")
		return
	}

	text := strings.TrimSpace(r.FormValue("text"))
	language := strings.TrimSpace(r.FormValue("language"))
	if text == "" || language == "" {
		writeError(w, http.StatusBadRequest, "text and language are required")
		return
	}

	speed := 1.0
	if speedStr := strings.TrimSpace(r.FormValue("speed")); speedStr != "" {
		if parsed, err := strconv.ParseFloat(speedStr, 64); err == nil && parsed > 0 {
			speed = parsed
		} else {
			writeError(w, http.StatusBadRequest, "invalid speed value")
			return
		}
	}

	emotion := strings.TrimSpace(r.FormValue("emotion"))
	if emotion == "" {
		emotion = "neutral"
	}

	modelType := strings.TrimSpace(r.FormValue("model_type"))
	if modelType == "" {
		modelType = "your_tts"
	}

	file, handler, err := r.FormFile("audio")
	if err != nil {
		writeError(w, http.StatusBadRequest, "audio file is required")
		return
	}
	defer file.Close()

	voiceCloneRoot := getEnv("VOICE_CLONE_ROOT", "/app/data/voice_clone")
	if err := os.MkdirAll(voiceCloneRoot, 0o755); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to prepare storage directory")
		return
	}

	jobDir := filepath.Join(voiceCloneRoot, fmt.Sprintf("job_%d_%d", userID, time.Now().UnixNano()))
	if err := os.MkdirAll(jobDir, 0o755); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to prepare job directory")
		return
	}

	inputExt := strings.ToLower(filepath.Ext(handler.Filename))
	if inputExt == "" {
		inputExt = ".wav"
	}
	inputPath := filepath.Join(jobDir, fmt.Sprintf("input%s", inputExt))
	outputPath := filepath.Join(jobDir, "output.wav")

	dest, err := os.Create(inputPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to store audio file")
		return
	}
	if _, err := io.Copy(dest, file); err != nil {
		dest.Close()
		writeError(w, http.StatusInternalServerError, "failed to save audio file")
		return
	}
	dest.Close()

	job := VoiceCloneJob{
		UserID:     userID,
		Text:       text,
		Language:   language,
		Speed:      speed,
		Emotion:    emotion,
		ModelType:  modelType,
		Status:     jobStatusPending,
		InputPath:  inputPath,
		OutputPath: outputPath,
	}

	if err := db.Create(&job).Error; err != nil {
		writeError(w, http.StatusInternalServerError, "failed to queue job")
		return
	}

	writeJSON(w, http.StatusAccepted, map[string]any{
		"job_id":  job.ID,
		"status":  job.Status,
		"message": "Voice clone job queued",
	})
}

func startVoiceCloneWorker() {
	go func() {
		log.Println("Voice clone worker started")
		for {
			job, err := claimNextVoiceCloneJob()
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					time.Sleep(voiceCloneWorkerIdleSleep)
					continue
				}
				log.Printf("voice clone worker: failed to claim job: %v", err)
				time.Sleep(voiceCloneWorkerErrorSleep)
				continue
			}

			if err := processVoiceCloneJob(job); err != nil {
				log.Printf("voice clone worker: job %d failed: %v", job.ID, err)
			}
		}
	}()
}

func claimNextVoiceCloneJob() (*VoiceCloneJob, error) {
	var job VoiceCloneJob

	tx := db.Begin()
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
		Where("status = ?", jobStatusPending).
		Order("created_at ASC").
		Take(&job).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&job).Updates(map[string]any{
		"status":        jobStatusProcessing,
		"error_message": "",
		"updated_at":    time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &job, nil
}

func processVoiceCloneJob(job *VoiceCloneJob) error {
	err := runVoiceCloneJob(job)

	status := jobStatusCompleted
	errorMessage := ""
	if err != nil {
		status = jobStatusFailed
		errorMessage = err.Error()
	}

	updatePayload := map[string]any{
		"status":        status,
		"error_message": errorMessage,
		"updated_at":    time.Now(),
	}

	if status == jobStatusCompleted {
		updatePayload["output_path"] = job.OutputPath
	}

	if err2 := db.Model(&VoiceCloneJob{}).Where("id = ?", job.ID).Updates(updatePayload).Error; err2 != nil {
		log.Printf("voice clone worker: failed to update job %d status: %v", job.ID, err2)
	}

	if status == jobStatusCompleted {
		createNotification(job.UserID, job.ID, "Voice clone job completed")
	} else {
		createNotification(job.UserID, job.ID, fmt.Sprintf("Voice clone job failed: %s", errorMessage))
	}

	return err
}

func runVoiceCloneJob(job *VoiceCloneJob) error {
	audioFile, err := os.Open(job.InputPath)
	if err != nil {
		return fmt.Errorf("failed to open input audio: %w", err)
	}
	defer audioFile.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile("audio", filepath.Base(job.InputPath))
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(formFile, audioFile); err != nil {
		return fmt.Errorf("failed to copy audio data: %w", err)
	}

	writer.WriteField("text", job.Text)
	writer.WriteField("language", job.Language)
	writer.WriteField("speed", fmt.Sprintf("%.2f", job.Speed))
	writer.WriteField("emotion", job.Emotion)
	writer.WriteField("model_type", job.ModelType)

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to finalize multipart payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, PYTHON_SERVICE_URL+"/voice-clone", body)
	if err != nil {
		return fmt.Errorf("failed to create python request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("python service request failed: %w", err)
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read python response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("python service returned %d: %s", resp.StatusCode, string(respData))
	}

	// Verify response is actually audio data (should be binary WAV)
	if len(respData) < 44 {
		log.Printf("WARNING: Response data too small (%d bytes) - may not be valid WAV file", len(respData))
		return fmt.Errorf("python service returned invalid audio data (size: %d bytes)", len(respData))
	}

	// Check for WAV header
	if string(respData[:4]) != "RIFF" {
		log.Printf("WARNING: Response does not start with RIFF header - first 100 bytes: %s", string(respData[:min(100, len(respData))]))
		return fmt.Errorf("python service returned non-WAV data")
	}

	if err := os.MkdirAll(filepath.Dir(job.OutputPath), 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(job.OutputPath, respData, 0o644); err != nil {
		return fmt.Errorf("failed to write output audio: %w", err)
	}

	log.Printf("Voice clone job %d: successfully wrote %d bytes to %s", job.ID, len(respData), job.OutputPath)
	return nil
}

func listVoiceCloneJobsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid user context")
		return
	}

	var jobs []VoiceCloneJob
	if err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&jobs).Error; err != nil {
		writeError(w, http.StatusInternalServerError, "failed to load jobs")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"jobs": jobs,
	})
}

func getVoiceCloneJobHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid user context")
		return
	}

	jobIDStr := r.URL.Query().Get("job_id")
	if jobIDStr == "" {
		writeError(w, http.StatusBadRequest, "job_id is required")
		return
	}
	jobID, err := strconv.ParseUint(jobIDStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid job_id")
		return
	}

	var job VoiceCloneJob
	if err := db.Where("id = ? AND user_id = ?", jobID, userID).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "job not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to load job")
		return
	}

	// Format response with proper JSON tags
	type JobResponse struct {
		ID           uint      `json:"id"`
		Status       string    `json:"status"`
		Text         string    `json:"text"`
		Language     string    `json:"language"`
		Speed        float64   `json:"speed"`
		Emotion      string    `json:"emotion"`
		ModelType    string    `json:"model_type"`
		InputPath    string    `json:"input_path"`
		OutputPath   string    `json:"output_path"`
		ErrorMessage string    `json:"error_message"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	response := JobResponse{
		ID:           job.ID,
		Status:       job.Status,
		Text:         job.Text,
		Language:     job.Language,
		Speed:        job.Speed,
		Emotion:      job.Emotion,
		ModelType:    job.ModelType,
		InputPath:    job.InputPath,
		OutputPath:   job.OutputPath,
		ErrorMessage: job.ErrorMessage,
		CreatedAt:    job.CreatedAt,
		UpdatedAt:    job.UpdatedAt,
	}

	writeJSON(w, http.StatusOK, response)
}

func getFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func downloadVoiceCloneJobHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid user context")
		return
	}

	jobIDStr := r.URL.Query().Get("job_id")
	if jobIDStr == "" {
		writeError(w, http.StatusBadRequest, "job_id is required")
		return
	}
	jobID, err := strconv.ParseUint(jobIDStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid job_id")
		return
	}

	var job VoiceCloneJob
	if err := db.Where("id = ? AND user_id = ?", jobID, userID).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "job not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to load job")
		return
	}

	if job.Status != jobStatusCompleted {
		writeError(w, http.StatusBadRequest, "job not completed yet")
		return
	}

	if _, err := os.Stat(job.OutputPath); err != nil {
		log.Printf("Output file not found at: %s (error: %v)", job.OutputPath, err)
		writeError(w, http.StatusInternalServerError, "output file not found")
		return
	}

	log.Printf("Serving voice clone file: %s (job_id: %d, size: %d bytes)", job.OutputPath, job.ID, getFileSize(job.OutputPath))
	w.Header().Set("Content-Type", "audio/wav")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"voice_clone_%d.wav\"", job.ID))
	http.ServeFile(w, r, job.OutputPath)
}

func listNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid user context")
		return
	}

	var notifications []Notification
	if err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error; err != nil {
		writeError(w, http.StatusInternalServerError, "failed to load notifications")
		return
	}

	var unreadCount int64
	db.Model(&Notification{}).Where("user_id = ? AND read = false", userID).Count(&unreadCount)

	// Format notifications for JSON response
	type NotificationResponse struct {
		ID        uint      `json:"id"`
		Message   string    `json:"message"`
		Read      bool      `json:"read"`
		CreatedAt time.Time `json:"created_at"`
	}

	var formattedNotifications []NotificationResponse
	for _, n := range notifications {
		formattedNotifications = append(formattedNotifications, NotificationResponse{
			ID:        n.ID,
			Message:   n.Message,
			Read:      n.Read,
			CreatedAt: n.CreatedAt,
		})
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"notifications": formattedNotifications,
		"unread_count":  unreadCount,
	})
}

func markNotificationsReadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid user context")
		return
	}

	type markRequest struct {
		NotificationIDs []uint `json:"notification_ids"`
	}

	var req markRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err != io.EOF {
		writeError(w, http.StatusBadRequest, "invalid payload")
		return
	}

	query := db.Model(&Notification{}).Where("user_id = ?", userID).Where("read = ?", false)
	if len(req.NotificationIDs) > 0 {
		query = query.Where("id IN ?", req.NotificationIDs)
	}

	if err := query.Update("read", true).Error; err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update notifications")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

func createNotification(userID uint, jobID uint, message string) {
	notification := Notification{
		UserID:  userID,
		JobID:   jobID,
		Message: message,
		Read:    false,
	}

	if err := db.Create(&notification).Error; err != nil {
		log.Printf("failed to create notification: %v", err)
	}
}

func finetuneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req finetuneRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if req.DatasetName == "" || req.Epochs <= 0 || req.BatchSize <= 0 {
		writeError(w, http.StatusBadRequest, "dataset_name, epochs and batch_size are required")
		return
	}

	proxyJSON(w, PYTHON_SERVICE_URL+"/train", req)
}

func datasetUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	if err := r.ParseMultipartForm(200 << 20); err != nil {
		writeError(w, http.StatusBadRequest, "failed to parse form data")
		return
	}

	datasetName := sanitizeName(r.FormValue("dataset_name"))
	transcript := strings.TrimSpace(r.FormValue("transcript"))
	if datasetName == "" || transcript == "" {
		writeError(w, http.StatusBadRequest, "dataset_name and transcript are required")
		return
	}

	file, handler, err := r.FormFile("audio")
	if err != nil {
		writeError(w, http.StatusBadRequest, "audio file is required")
		return
	}
	defer file.Close()

	sourceExt := strings.ToLower(filepath.Ext(handler.Filename))
	if sourceExt == "" {
		sourceExt = ".mp3"
	}

	baseName := sanitizeName(strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename)))
	if baseName == "" {
		baseName = fmt.Sprintf("clip_%d", time.Now().Unix())
	}
	wavName := fmt.Sprintf("%s_%d.wav", baseName, time.Now().Unix())

	dataRoot := getEnv("DATA_ROOT", "/app/data")
	datasetDir := filepath.Join(dataRoot, datasetName)
	wavDir := filepath.Join(datasetDir, "wavs")
	tmpDir := filepath.Join(datasetDir, "tmp")
	if err := os.MkdirAll(wavDir, 0o755); err != nil {
		log.Printf("failed to create wav dir: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to prepare dataset directory")
		return
	}
	if err := os.MkdirAll(tmpDir, 0o755); err != nil {
		log.Printf("failed to create tmp dir: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to prepare dataset directory")
		return
	}

	mp3Path := filepath.Join(tmpDir, fmt.Sprintf("%s%s", baseName, sourceExt))
	dest, err := os.Create(mp3Path)
	if err != nil {
		log.Printf("failed to create temp file: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to store audio")
		return
	}
	if _, err := io.Copy(dest, file); err != nil {
		dest.Close()
		log.Printf("failed to copy audio: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to store audio")
		return
	}
	dest.Close()

	wavPath := filepath.Join(wavDir, wavName)
	if err := convertToWav(mp3Path, wavPath); err != nil {
		log.Printf("ffmpeg conversion failed: %v", err)
		writeError(w, http.StatusInternalServerError, "audio conversion failed")
		return
	}
	_ = os.Remove(mp3Path)

	if err := appendMetadata(datasetDir, wavName, transcript); err != nil {
		log.Printf("failed to update metadata: %v", err)
		writeError(w, http.StatusInternalServerError, "unable to update metadata")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"status":        "stored",
		"dataset":       datasetName,
		"wav_filename":  wavName,
		"metadata_path": filepath.Join(datasetDir, "metadata.csv"),
	})
}

func proxyJSON(w http.ResponseWriter, targetURL string, payload any) {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal payload: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	proxyReq, err := http.NewRequest(http.MethodPost, targetURL, bytes.NewReader(body))
	if err != nil {
		log.Printf("failed to create request: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	proxyReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(proxyReq)
	if err != nil {
		log.Printf("python service request failed: %v", err)
		writeError(w, http.StatusBadGateway, "python service unavailable")
		return
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	// Log response for debugging
	log.Printf("Python service response: status=%d, body=%s", resp.StatusCode, string(respBody))

	// Normalize error responses from FastAPI (convert "detail" to "error")
	if resp.StatusCode >= 400 {
		var errResp map[string]interface{}
		if err := json.Unmarshal(respBody, &errResp); err == nil {
			if detail, ok := errResp["detail"]; ok && errResp["error"] == nil {
				errResp["error"] = detail
				delete(errResp, "detail")
				if normalizedBody, err := json.Marshal(errResp); err == nil {
					respBody = normalizedBody
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if _, err := w.Write(respBody); err != nil {
		log.Printf("failed to write response body: %v", err)
	}
}

func convertToWav(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-acodec", "pcm_s16le", "-ar", "22050", outputPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w", err)
	}
	return nil
}

func appendMetadata(datasetDir, wavName, transcript string) error {
	metaPath := filepath.Join(datasetDir, "metadata.csv")
	f, err := os.OpenFile(metaPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s|%s\n", wavName, transcript))
	return err
}

func sanitizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	return name
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func newHTTPClient() *http.Client {
	timeoutStr := getEnv("HTTP_CLIENT_TIMEOUT_SECONDS", "1800")
	timeoutSeconds, err := strconv.Atoi(timeoutStr)
	if err != nil || timeoutSeconds <= 0 {
		log.Printf("invalid HTTP_CLIENT_TIMEOUT_SECONDS='%s', falling back to 1800 seconds", timeoutStr)
		timeoutSeconds = 1800
	}
	return &http.Client{Timeout: time.Duration(timeoutSeconds) * time.Second}
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func adminTrainingUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	if err := r.ParseMultipartForm(200 << 20); err != nil {
		writeError(w, http.StatusBadRequest, "failed to parse form data")
		return
	}

	datasetName := sanitizeName(r.FormValue("dataset_name"))
	transcript := strings.TrimSpace(r.FormValue("transcript"))
	if datasetName == "" || transcript == "" {
		writeError(w, http.StatusBadRequest, "dataset_name and transcript are required")
		return
	}

	file, handler, err := r.FormFile("audio")
	if err != nil {
		writeError(w, http.StatusBadRequest, "audio file is required")
		return
	}
	defer file.Close()

	sourceExt := strings.ToLower(filepath.Ext(handler.Filename))
	if sourceExt == "" {
		sourceExt = ".mp3"
	}

	baseName := sanitizeName(strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename)))
	if baseName == "" {
		baseName = fmt.Sprintf("clip_%d", time.Now().Unix())
	}
	wavName := fmt.Sprintf("%s_%d.wav", baseName, time.Now().Unix())

	dataRoot := getEnv("DATA_ROOT", "/app/data")
	datasetDir := filepath.Join(dataRoot, datasetName)
	wavDir := filepath.Join(datasetDir, "wavs")
	tmpDir := filepath.Join(datasetDir, "tmp")
	if err := os.MkdirAll(wavDir, 0o755); err != nil {
		log.Printf("failed to create wav dir: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to prepare dataset directory")
		return
	}
	if err := os.MkdirAll(tmpDir, 0o755); err != nil {
		log.Printf("failed to create tmp dir: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to prepare dataset directory")
		return
	}

	mp3Path := filepath.Join(tmpDir, fmt.Sprintf("%s%s", baseName, sourceExt))
	dest, err := os.Create(mp3Path)
	if err != nil {
		log.Printf("failed to create temp file: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to store audio")
		return
	}
	if _, err := io.Copy(dest, file); err != nil {
		dest.Close()
		log.Printf("failed to copy audio: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to store audio")
		return
	}
	dest.Close()

	wavPath := filepath.Join(wavDir, wavName)
	if err := convertToWav(mp3Path, wavPath); err != nil {
		log.Printf("ffmpeg conversion failed: %v", err)
		writeError(w, http.StatusInternalServerError, "audio conversion failed")
		return
	}
	_ = os.Remove(mp3Path)

	if err := appendMetadata(datasetDir, wavName, transcript); err != nil {
		log.Printf("failed to update metadata: %v", err)
		writeError(w, http.StatusInternalServerError, "unable to update metadata")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"status":        "stored",
		"dataset":       datasetName,
		"wav_filename":  wavName,
		"metadata_path": filepath.Join(datasetDir, "metadata.csv"),
	})
}

func adminTrainingStartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req finetuneRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if req.DatasetName == "" || req.Epochs <= 0 || req.BatchSize <= 0 {
		writeError(w, http.StatusBadRequest, "dataset_name, epochs and batch_size are required")
		return
	}

	proxyJSON(w, PYTHON_SERVICE_URL+"/train", req)
}

func adminListDatasetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	dataRoot := getEnv("DATA_ROOT", "/app/data")

	type DatasetInfo struct {
		Name       string `json:"name"`
		AudioCount int    `json:"audio_count"`
		CreatedAt  string `json:"created_at"`
	}

	var datasets []DatasetInfo

	entries, err := os.ReadDir(dataRoot)
	if err != nil {
		if os.IsNotExist(err) {
			writeJSON(w, http.StatusOK, map[string]interface{}{
				"status":   "ok",
				"datasets": []DatasetInfo{},
			})
			return
		}
		log.Printf("failed to read data root: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to list datasets")
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		wavDir := filepath.Join(dataRoot, entry.Name(), "wavs")
		audioCount := 0
		if wavEntries, err := os.ReadDir(wavDir); err == nil {
			audioCount = len(wavEntries)
		}

		datasetPath := filepath.Join(dataRoot, entry.Name())
		info, err := os.Stat(datasetPath)
		createdAt := ""
		if err == nil {
			createdAt = info.ModTime().Format("2006-01-02 15:04:05")
		}

		datasets = append(datasets, DatasetInfo{
			Name:       entry.Name(),
			AudioCount: audioCount,
			CreatedAt:  createdAt,
		})
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status":   "ok",
		"datasets": datasets,
	})
}

func adminListTrainingJobsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	proxyRequest, err := http.NewRequest(http.MethodGet, PYTHON_SERVICE_URL+"/training-jobs", nil)
	if err != nil {
		log.Printf("failed to create request: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp, err := httpClient.Do(proxyRequest)
	if err != nil {
		log.Printf("python service request failed: %v", err)
		writeError(w, http.StatusBadGateway, "python service unavailable")
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if _, err := w.Write(respBody); err != nil {
		log.Printf("failed to write response body: %v", err)
	}
}
