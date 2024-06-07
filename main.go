package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	fmt.Printf("ffmpeg Parser - (c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	exefile := "/ffmpeg/bin/ffmpeg.exe"
	wdir := "/tmp/"
	if _, err := os.Stat(exefile); err == nil {
		fmt.Printf("- Parser Detected")
		files, err := ioutil.ReadDir(wdir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if path.Ext(file.Name()) == ".mp4" {
				cmd := exec.Command(exefile, "-ss", "00:00:10", "-i", wdir+file.Name(), "-vframes", "100", "-s", "640x480", wdir+fileNameWithoutExtension(file.Name())+".png")
				fmt.Println(cmd)
				if err := cmd.Run(); err != nil {
					fmt.Println("Error: ", err)
				}
			}
		}
	} else {
		fmt.Println(err)
	}

}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
