package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	input := filepath.Join("input", "example.mp3")
	outputDir := "output"

	outputPath, meta, err := NormalizeAudio(input, outputDir)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Successfully normalized file:", outputPath)
	fmt.Printf("Duration: %.2f seconds\n", meta.Duration)
	fmt.Println("Sample rate:", meta.SampleRate)
}
