#!/usr/bin/env python3
"""
Test script for voice clone text splitting functionality.
Tests the split_text_into_chunks function with various languages and text lengths.
"""

import sys
from pathlib import Path

# Add the python-service app to the path
sys.path.insert(0, str(Path(__file__).parent / "python-service"))

from app.main import split_text_into_chunks


def test_english_text():
    """Test English text splitting"""
    print("\n" + "="*80)
    print("TEST 1: English Text Splitting")
    print("="*80)
    
    text = "This is a long sentence that exceeds the character limit and needs to be split into multiple chunks for synthesis with XTTS v2 voice cloning technology. The system should intelligently split at sentence boundaries first, then at phrase boundaries if needed."
    
    chunks = split_text_into_chunks(text, max_chars=240, language="en")
    
    print(f"Original text length: {len(text)} characters")
    print(f"Number of chunks: {len(chunks)}\n")
    
    for i, chunk in enumerate(chunks, 1):
        print(f"Chunk {i} ({len(chunk)} chars):")
        print(f"  {chunk}\n")


def test_farsi_text():
    """Test Farsi text splitting"""
    print("\n" + "="*80)
    print("TEST 2: Farsi Text Splitting")
    print("="*80)
    
    text = "این یک متن طولانی است که از حد مجاز کاراکتر تجاوز می‌کند و باید به چند قسمت تقسیم شود تا بتوان آن را با مدل XTTS v2 برای تولید صدای کلون شده سنتز کرد۔ سیستم باید به طور هوشمندانه در مرزهای جملات اول تقسیم کند۔"
    
    chunks = split_text_into_chunks(text, max_chars=240, language="fa")
    
    print(f"Original text length: {len(text)} characters")
    print(f"Number of chunks: {len(chunks)}\n")
    
    for i, chunk in enumerate(chunks, 1):
        print(f"Chunk {i} ({len(chunk)} chars):")
        print(f"  {chunk}\n")


def test_german_text():
    """Test German text splitting"""
    print("\n" + "="*80)
    print("TEST 3: German Text Splitting")
    print("="*80)
    
    text = "Dies ist ein langer Satz, der die Zeichenbegrenzung überschreitet und in mehrere Chunks aufgeteilt werden muss, um mit dem XTTS v2-Modell für die Sprachsynthese mit geklonter Stimme verarbeitet zu werden. Das System sollte intelligent an Satzgrenzen aufteilen."
    
    chunks = split_text_into_chunks(text, max_chars=240, language="de")
    
    print(f"Original text length: {len(text)} characters")
    print(f"Number of chunks: {len(chunks)}\n")
    
    for i, chunk in enumerate(chunks, 1):
        print(f"Chunk {i} ({len(chunk)} chars):")
        print(f"  {chunk}\n")


def test_short_text():
    """Test short text that doesn't need splitting"""
    print("\n" + "="*80)
    print("TEST 4: Short Text (No Splitting Needed)")
    print("="*80)
    
    text = "Hello, this is a short text."
    
    chunks = split_text_into_chunks(text, max_chars=240, language="en")
    
    print(f"Original text length: {len(text)} characters")
    print(f"Number of chunks: {len(chunks)}\n")
    
    for i, chunk in enumerate(chunks, 1):
        print(f"Chunk {i} ({len(chunk)} chars):")
        print(f"  {chunk}\n")


def test_multiple_sentences():
    """Test text with multiple sentences"""
    print("\n" + "="*80)
    print("TEST 5: Multiple Sentences")
    print("="*80)
    
    text = "First sentence. Second sentence! Third sentence? Fourth sentence with more content that might need splitting."
    
    chunks = split_text_into_chunks(text, max_chars=100, language="en")
    
    print(f"Original text length: {len(text)} characters")
    print(f"Number of chunks: {len(chunks)}\n")
    
    for i, chunk in enumerate(chunks, 1):
        print(f"Chunk {i} ({len(chunk)} chars):")
        print(f"  {chunk}\n")


def test_edge_cases():
    """Test edge cases"""
    print("\n" + "="*80)
    print("TEST 6: Edge Cases")
    print("="*80)
    
    # Test 1: Empty text
    print("Test 6a: Empty text")
    chunks = split_text_into_chunks("", max_chars=240, language="en")
    print(f"  Chunks: {chunks}\n")
    
    # Test 2: Single word longer than limit
    print("Test 6b: Single very long word")
    text = "Supercalifragilisticexpialidocious" * 10
    chunks = split_text_into_chunks(text, max_chars=50, language="en")
    print(f"  Original length: {len(text)} chars")
    print(f"  Number of chunks: {len(chunks)}")
    print(f"  Chunk 1: {chunks[0][:50]}...\n")
    
    # Test 3: Text with only whitespace
    print("Test 6c: Whitespace only")
    chunks = split_text_into_chunks("   \n\n   ", max_chars=240, language="en")
    print(f"  Chunks: {chunks}\n")


def test_character_limits():
    """Verify all chunks respect the character limit"""
    print("\n" + "="*80)
    print("TEST 7: Character Limit Verification")
    print("="*80)
    
    test_cases = [
        ("en", "This is a test. " * 20),
        ("fa", "این یک متن است۔ " * 20),
        ("de", "Dies ist ein Test. " * 20),
    ]
    
    for lang, text in test_cases:
        chunks = split_text_into_chunks(text, max_chars=240, language=lang)
        print(f"\nLanguage: {lang}")
        print(f"  Original length: {len(text)} chars")
        print(f"  Number of chunks: {len(chunks)}")
        
        all_valid = True
        for i, chunk in enumerate(chunks, 1):
            if len(chunk) > 240:
                print(f"  ❌ Chunk {i} exceeds limit: {len(chunk)} chars")
                all_valid = False
        
        if all_valid:
            print(f"  ✓ All chunks within 240 character limit")
            max_chunk = max(len(c) for c in chunks)
            min_chunk = min(len(c) for c in chunks)
            print(f"  Chunk sizes: min={min_chunk}, max={max_chunk}")


if __name__ == "__main__":
    print("\n" + "="*80)
    print("VOICE CLONE TEXT SPLITTING TEST SUITE")
    print("="*80)
    
    try:
        test_english_text()
        test_farsi_text()
        test_german_text()
        test_short_text()
        test_multiple_sentences()
        test_edge_cases()
        test_character_limits()
        
        print("\n" + "="*80)
        print("ALL TESTS COMPLETED SUCCESSFULLY ✓")
        print("="*80 + "\n")
        
    except Exception as e:
        print(f"\n❌ ERROR: {e}")
        import traceback
        traceback.print_exc()
        sys.exit(1)
