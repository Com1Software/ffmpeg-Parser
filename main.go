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
	xpage := "ffmpeg-parse.htm"
	if _, err := os.Stat(exefile); err == nil {
		fmt.Printf("- Parser Detected")
		files, err := ioutil.ReadDir(wdir)
		if err != nil {
			log.Fatal(err)
		}
		xdata := "<!DOCTYPE html>"
		xdata = xdata + "<html>"
		xdata = xdata + "<head>"
		xdata = xdata + "<title>ffmpeg Parse for " + wdir + "</title>"
		xdata = xdata + "</head>"
		xdata = xdata + "<body>"
		xdata = xdata + "<H1>ffmpeg Parse for" + wdir + "</H1>"
		for _, file := range files {
			if path.Ext(file.Name()) == ".mp4" {
				tfile := wdir + file.Name()
				tnfile := fixFileName(tfile)
				cmd := exec.Command(exefile, "-ss", "00:00:10", "-i", tnfile, "-vframes", "100", "-s", "64x48", fileNameWithoutExtension(tnfile)+".png")
				fmt.Println(cmd)

				xdata = xdata + "  <A HREF='file:///C:/" + wdir + tnfile + "'> <IMG SRC=" + fileNameWithoutExtension(tnfile) + ".png" + "  ALT=error>  [ " + file.Name() + " ] </A> <BR>"
				if err := cmd.Run(); err != nil {
					fmt.Println("Error: ", err)
				}
			}
		}

		xdata = xdata + " </body>"
		xdata = xdata + " </html>"
		err = os.WriteFile(xpage, []byte(xdata), 0644)
		if err != nil {
			fmt.Println(err)
		}
		Openbrowser(xpage)
	} else {
		fmt.Println(err)
	}

}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func fixFileName(fileName string) string {
	newName := ""
	tmp := strings.Split(fileName, " ")
	for x := 0; x < len(tmp); x++ {
		newName = newName + tmp[x]
	}
	err := os.Rename(fileName, newName)
	if err != nil {
		fmt.Println("Error renaming file:", err)
	} else {
		fmt.Println("File renamed successfully")
	}
	return newName
}

func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)

	return exec.Command(cmd, args...).Start()
}
