package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// Check if text argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: xay \"text to speak\"")
		os.Exit(1)
	}

	// Get OpenAI API key from environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable is not set")
		fmt.Println("Please set your OpenAI API key using:")
		fmt.Println("export OPENAI_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	// Get text to speak by joining all arguments
	text := strings.Join(os.Args[1:], " ")

	// Create OpenAI client
	client := openai.NewClient(apiKey)

	// Create temporary file for audio
	tempDir := os.TempDir()
	audioFile := filepath.Join(tempDir, "xay_output.mp3")

	// Generate speech using OpenAI API
	resp, err := client.CreateSpeech(context.Background(), openai.CreateSpeechRequest{
		Model:          openai.TTSModel1,
		Input:          text,
		Voice:          openai.VoiceAlloy,
		ResponseFormat: "mp3",
	})
	if err != nil {
		fmt.Printf("Error creating speech: %v\n", err)
		os.Exit(1)
	}
	defer resp.Close()

	// Save audio to temporary file
	audioData, err := os.Create(audioFile)
	if err != nil {
		fmt.Printf("Error creating audio file: %v\n", err)
		os.Exit(1)
	}
	defer audioData.Close()

	_, err = audioData.ReadFrom(resp)
	if err != nil {
		fmt.Printf("Error writing audio data: %v\n", err)
		os.Exit(1)
	}

	// Play the audio file using afplay (macOS built-in audio player)
	cmd := exec.Command("afplay", audioFile)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error playing audio: %v\n", err)
		os.Exit(1)
	}

	// Clean up temporary file
	os.Remove(audioFile)
} 