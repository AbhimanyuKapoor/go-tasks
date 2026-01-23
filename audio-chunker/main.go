package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	input := filepath.Join("input", "normalized.wav")
	outputDir := "output"

	paths, err := SplitAudio(input, outputDir)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Successfully created %d chunks", len(paths))
}
