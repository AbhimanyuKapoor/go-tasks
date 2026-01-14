package main

import (
	"fmt"
	"log"
)

func main() {
	input := "input/example.wav"
	outputDir := "output"

	paths, err := SplitAudio(input, outputDir)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Successfully created %d chunks", len(paths))
}
