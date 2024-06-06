package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("ffmpeg Parser - (c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	exefile := "ffmpeg"

	if CheckForFile(exefile) {
		fmt.Printf("- Parser Detected")
	}

}

func CheckForFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	file.Close()
	return true
}
