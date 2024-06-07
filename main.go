package main

import (
	"fmt"
	"os"
	"runtime"
	"io/ioutil"
        "log"
)

func main() {
	fmt.Printf("ffmpeg Parser - (c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	exefile := "ffmpeg"

	if CheckForFile(exefile) {
		fmt.Printf("- Parser Detected")
	        files, err := ioutil.ReadDir("/tmp/")
                if err != nil {
                  log.Fatal(err)
                }
                for _, file := range files {
                     fmt.Println(file.Name(), file.IsDir())
                }
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
