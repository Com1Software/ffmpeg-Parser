package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
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
			fmt.Println(file.Name(), file.IsDir())

			cmd := exec.Command(exefile, "-i", wdir+file.Name(), "-vframes", "1", "-s", "640x480", wdir+file.Name()+".png")
			fmt.Println(cmd)
			if err := cmd.Run(); err != nil {
				fmt.Println("Error: ", err)
			}
		}
	} else {
		fmt.Println(err)
	}

}
