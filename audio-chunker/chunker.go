package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func SplitAudio(input string, chunkDir string) ([]string, error) {
	var chunkPaths []string

	// Create output dir if doesn't exist
	err := os.MkdirAll(chunkDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("Failed to create directory: %v", err)
	}

	// Total duration of audio file
	duration, err := getDuration(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to get duration: %v", err)
	}

	// 50% overlap: 2.5s increment of start time
	for start := 0.0; start < duration; start += 2.5 {
		outputFile := filepath.Join(chunkDir, fmt.Sprintf("chunk_%06.1f.wav", start))

		// FFmpeg command
		cmd := exec.Command("ffmpeg", "-y", "-ss", strconv.FormatFloat(start, 'f', 2, 64),
			"-i", input,
			"-t", "5",
			"-c", "copy",
			outputFile)

		err := cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("ffmpeg error at start time %f: %v", start, err)
		}

		chunkPaths = append(chunkPaths, outputFile)
	}

	return chunkPaths, nil
}

// Helper function to get duration
func getDuration(input string) (float64, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1", input)

	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	durationStr := strings.TrimSpace(string(out))
	return strconv.ParseFloat(durationStr, 64)
}
