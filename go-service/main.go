package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PYTHON_SERVICE_URL = "http://python-service:5000"
	JWT_SECRET         = "my_super_secret_key_2025"
)

var db *gorm.DB
var httpClient = &http.Client{Timeout: 60 * time.Second}

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
	Role     string // "admin" or "user"
	Approved bool   `gorm:"default:false"` // Admin approval required
}

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

	db.AutoMigrate(&User{})

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

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/auth/register", handleRegister)
	mux.HandleFunc("/api/v1/auth/login", handleLogin)

	mux.HandleFunc("/api/v1/dataset/upload", authMiddleware(datasetUploadHandler))
	mux.HandleFunc("/api/v1/finetune/start", authMiddleware(finetuneHandler))
	mux.HandleFunc("/api/v1/synthesize", authMiddleware(synthesizeHandler))

	mux.HandleFunc("/api/v1/admin/users", adminMiddleware(handleListUsers))
	mux.HandleFunc("/api/v1/admin/approve", adminMiddleware(handleApproveUser))
	mux.HandleFunc("/api/v1/admin/test", adminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"message": "Admin access granted"})
	}))

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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	user := User{Email: req.Email, Password: string(hashedPassword), Role: "user"}

	if result := db.Create(&user); result.Error != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	var user User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Check if user is approved
	if !user.Approved && user.Role != "admin" {
		http.Error(w, "Account pending admin approval", http.StatusForbidden)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(JWT_SECRET))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Token: tokenString, Role: user.Role, Email: user.Email})
}

func handleListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleApproveUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type ApproveRequest struct {
		UserID uint `json:"user_id"`
	}

	var req ApproveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if result := db.Model(&User{}).Where("id = ?", req.UserID).Update("approved", true); result.Error != nil {
		http.Error(w, "Failed to approve user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "User approved"})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func adminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

func synthesizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req synthesizeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}

	if req.Text == "" || req.SpeakerID == "" {
		http.Error(w, "text and speaker_id are required", http.StatusBadRequest)
		return
	}

	proxyJSON(w, PYTHON_SERVICE_URL+"/tts", req)
}

func finetuneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req finetuneRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}

	if req.DatasetName == "" || req.Epochs <= 0 || req.BatchSize <= 0 {
		http.Error(w, "dataset_name, epochs and batch_size are required", http.StatusBadRequest)
		return
	}

	proxyJSON(w, PYTHON_SERVICE_URL+"/train", req)
}

func datasetUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(200 << 20); err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	datasetName := sanitizeName(r.FormValue("dataset_name"))
	transcript := strings.TrimSpace(r.FormValue("transcript"))
	if datasetName == "" || transcript == "" {
		http.Error(w, "dataset_name and transcript are required", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("audio")
	if err != nil {
		http.Error(w, "audio file is required", http.StatusBadRequest)
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
		http.Error(w, "failed to prepare dataset directory", http.StatusInternalServerError)
		return
	}
	if err := os.MkdirAll(tmpDir, 0o755); err != nil {
		log.Printf("failed to create tmp dir: %v", err)
		http.Error(w, "failed to prepare dataset directory", http.StatusInternalServerError)
		return
	}

	mp3Path := filepath.Join(tmpDir, fmt.Sprintf("%s%s", baseName, sourceExt))
	dest, err := os.Create(mp3Path)
	if err != nil {
		log.Printf("failed to create temp file: %v", err)
		http.Error(w, "failed to store audio", http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(dest, file); err != nil {
		dest.Close()
		log.Printf("failed to copy audio: %v", err)
		http.Error(w, "failed to store audio", http.StatusInternalServerError)
		return
	}
	dest.Close()

	wavPath := filepath.Join(wavDir, wavName)
	if err := convertToWav(mp3Path, wavPath); err != nil {
		log.Printf("ffmpeg conversion failed: %v", err)
		http.Error(w, "audio conversion failed", http.StatusInternalServerError)
		return
	}
	_ = os.Remove(mp3Path)

	if err := appendMetadata(datasetDir, wavName, transcript); err != nil {
		log.Printf("failed to update metadata: %v", err)
		http.Error(w, "unable to update metadata", http.StatusInternalServerError)
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
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	proxyReq, err := http.NewRequest(http.MethodPost, targetURL, bytes.NewReader(body))
	if err != nil {
		log.Printf("failed to create request: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	proxyReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(proxyReq)
	if err != nil {
		log.Printf("python service request failed: %v", err)
		http.Error(w, "python service unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("failed to copy response body: %v", err)
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

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
