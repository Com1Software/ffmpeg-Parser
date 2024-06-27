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
	"strconv"
	"strings"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

func main() {
	fmt.Printf("ffmpeg Parser - (c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)

	exefile := "/ffmpeg/bin/ffmpeg.exe"
	exefilea := "/ffmpeg/bin/ffprobe.exe"
	drive := "c"
	wdir := drive + ":/tmp/"
	subdir := true
	xpage := "ffmpeg-parse.htm"
	display := 0
	if _, err := os.Stat(exefile); err == nil {
		fmt.Printf("- Parser Detected")

		xdata := "<!DOCTYPE html>"
		xdata = xdata + "<html>"
		xdata = xdata + "<head>"
		switch {
		case display == 1:
			xdata = xdata + "<meta name='viewport' content='width=device-width, initial-scale=1'>"

			xdata = xdata + "<style>"
			xdata = xdata + "div.scroll-container {"
			xdata = xdata + "background-color: #333;"
			xdata = xdata + "overflow: auto;"
			xdata = xdata + "white-space: nowrap;"
			xdata = xdata + "padding: 10px;"
			xdata = xdata + "}"
			xdata = xdata + "div.scroll-container img {"
			xdata = xdata + "padding: 10px;"
			xdata = xdata + "}"
			xdata = xdata + "</style>"

		}

		xdata = xdata + "<title>ffmpeg Parse for " + wdir + "</title>"
		xdata = xdata + "</head>"
		xdata = xdata + "<body>"
		xdata = xdata + "<H1>ffmpeg Parse for " + wdir + "</H1>"
		files, err := ioutil.ReadDir(wdir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if ValidFileType(strings.ToLower(path.Ext(file.Name()))) {
				tfile := wdir + file.Name()
				tnfile := fixFileName(tfile)
				switch {
				case display == 0:
					xdata = xdata + BasicDisplay(exefile, tnfile, file.Name())
				case display == 1:
					xdata = xdata + ImageScrollDisplay(exefile, tnfile, file.Name())
				}
				xdata = xdata + FileData(exefilea, tnfile, file.Name())
				//-------------------------------------------------------------------------------------------------

			}
		}

		if subdir {
			entries, err := os.ReadDir(wdir + "./")
			if err != nil {
				log.Fatal(err)
			}

			for _, e := range entries {
				fmt.Println(e.Name())
				files, err = ioutil.ReadDir(wdir + e.Name())
				if err != nil {
					log.Fatal(err)
				}
				for _, file := range files {
					if ValidFileType(strings.ToLower(path.Ext(file.Name()))) {
						tfile := wdir + e.Name() + "/" + file.Name()
						tnfile := fixFileName(tfile)
						fmt.Println(tnfile)
						xdata = xdata + "[" + e.Name() + "]<BR>"
						switch {
						case display == 0:
							xdata = xdata + BasicDisplay(exefile, tnfile, file.Name())
						case display == 1:
							xdata = xdata + ImageScrollDisplay(exefile, tnfile, file.Name())
						}
						xdata = xdata + FileData(exefilea, tnfile, file.Name())
						//-------------------------------------------------------------------------------------------------

					}
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
	chr := ""
	ascval := 0

	for x := 0; x < len(fileName); x++ {
		chr = fileName[x : x+1]
		ascval = asciistring.StringToASCII(chr)
		switch {
		case ascval < 45:
		case ascval == 64:
		case ascval == 92:
		case ascval == 96:
		case ascval > 122:
		default:
			newName = newName + chr
		}
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

func ValidFileType(fileExt string) bool {
	rtn := false
	switch {
	case fileExt == ".mp4":
		rtn = true
	case fileExt == ".avi":
		rtn = true
	case fileExt == ".wmv":
		rtn = true
	}
	return rtn
}

func ParseFrameRate(data string) string {
	rtn := ""
	chr := ""
	do := false
	pass := 1
	v1 := ""
	v2 := ""
	add := true
	ascval := 0
	for x := 0; x < len(data); x++ {
		chr = data[x : x+1]
		add = true
		ascval = asciistring.StringToASCII(chr)
		if ascval == 13 {
			add = false
			numerator, _ := strconv.Atoi(v1)
			denominator, _ := strconv.Atoi(v2)
			if denominator > 0 {
				fps := numerator / denominator
				// fmt.Printf("The frame rate is: %.2f fps\n", fps)
				rtn = strconv.Itoa(fps)
			}
		}
		if ascval == 10 {
			add = false
			pass = 1
			do = false
		}
		if chr == "," {
			do = true
			add = false
		}
		if chr == "/" {
			pass = 2
			add = false
		}
		if do {
			if add {
				if pass == 1 {
					v1 = v1 + chr
				}
				if pass == 2 {
					v2 = v2 + chr
				}
			}
		}
	}
	return rtn
}

func ParseBitRate(data string) string {
	rtn := ""
	chr := ""
	do := false
	add := true
	pass := 1
	ascval := 0
	for x := 0; x < len(data); x++ {
		chr = data[x : x+1]
		add = true
		ascval = asciistring.StringToASCII(chr)
		if ascval == 13 {
			add = false
		}
		if ascval == 10 {
			add = false
			do = false
			pass = 2
		}
		if chr == "," {
			do = true
			add = false
		}
		if do {
			if add {
				if pass == 2 {
					rtn = rtn + chr
				}
			}
		}
	}
	return rtn
}

func FileData(exefilea string, tnfile string, fileName string) string {
	xdata := ""
	bfile := "tmp.bat"
	bdata := []byte(exefilea + " -i " + tnfile + " -show_entries stream=width,height -of csv=" + fmt.Sprintf("%q", "p=0") + ">tmp.csv")
	err := os.WriteFile(bfile, bdata, 0644)
	cmd := exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat := []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata := string(dat)
	tmp := strings.Split(tdata, ",")
	xdata = xdata + "Frame width " + tmp[0] + "<BR>"
	xdata = xdata + "Frame height " + tmp[1] + "<BR>"

	//-------------------------------------------------------------------------------------------------
	bdata = []byte(exefilea + " -i " + tnfile + " -show_entries format=duration -v quiet -of csv >tmp.csv")
	err = os.WriteFile(bfile, bdata, 0644)
	cmd = exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat = []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata = string(dat)
	tmp = strings.Split(tdata, ",")
	tmpa := strings.Split(tmp[1], ".")
	t := tmpa[0]
	i, _ := strconv.Atoi(t)
	mc := 0
	m := 0
	sc := 0
	for x := 0; x < i; x++ {
		mc++
		sc++
		if mc > 59 {
			m++
			mc = 0
			sc = 0
		}

	}
	xdata = xdata + "Length  " + strconv.Itoa(m) + ":" + strconv.Itoa(sc) + " <BR>"
	//-------------------------------------------------------------------------------------------------
	bdata = []byte(exefilea + " -i " + tnfile + " -show_entries stream=r_frame_rate  -of csv" + ">tmp.csv")

	err = os.WriteFile(bfile, bdata, 0644)
	cmd = exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat = []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata = string(dat)
	fr := ParseFrameRate(tdata)
	xdata = xdata + "Frames per second  " + fr + " <BR>"

	//-------------------------------------------------------------------------------------------------
	bdata = []byte(exefilea + " -i " + tnfile + "  -show_entries stream=bit_rate -v quiet -of csv >tmp.csv")
	err = os.WriteFile(bfile, bdata, 0644)
	cmd = exec.Command(bfile)
	if err = cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	dat = []byte("")
	dat, err = os.ReadFile("tmp.csv")
	tdata = string(dat)
	br := ParseBitRate(tdata)
	xdata = xdata + "Bit Rate " + br + " <BR>"
	xdata = xdata + "<BR><BR>"

	return xdata
}

func BasicDisplay(exefile string, tnfile string, fileName string) string {
	xdata := ""
	cmd := exec.Command(exefile, "-ss", "00:00:01", "-i", tnfile, "-vframes", "100", "-s", "128x96", fileNameWithoutExtension(tnfile)+"1.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", "00:00:10", "-i", tnfile, "-vframes", "100", "-s", "128x96", fileNameWithoutExtension(tnfile)+"2.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", "00:00:20", "-i", tnfile, "-vframes", "100", "-s", "128x96", fileNameWithoutExtension(tnfile)+"3.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	xdata = xdata + "  <A HREF='file:///" + tnfile + "'>  [ " + fileName + " ] <BR> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "1.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "2.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "3.png" + "  ALT=error> </A><BR> "
	//-------------------------------------------------------------------------------------------------
	return xdata
}

func ImageScrollDisplay(exefile string, tnfile string, fileName string) string {
	xdata := ""
	cmd := exec.Command(exefile, "-ss", "00:00:01", "-i", tnfile, "-vframes", "100", "-s", "128x96", fileNameWithoutExtension(tnfile)+"1.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", "00:00:10", "-i", tnfile, "-vframes", "100", "-s", "128x96", fileNameWithoutExtension(tnfile)+"2.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	cmd = exec.Command(exefile, "-ss", "00:00:20", "-i", tnfile, "-vframes", "100", "-s", "128x96", fileNameWithoutExtension(tnfile)+"3.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command %s \n Error: %s\n", cmd, err)
	}
	//xdata = xdata + "  <A HREF='file:///" + tnfile + "'>  [ " + fileName + " ] <BR> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "1.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "2.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "3.png" + "  ALT=error> </A><BR> "

	xdata = xdata + "<div class='scroll-container'>"
	xdata = xdata + "<img src= " + fileNameWithoutExtension(tnfile) + "1.png allt='test' width='128' height='96'>"
	xdata = xdata + "<img src= " + fileNameWithoutExtension(tnfile) + "2.png allt='test' width='128' height='96'>"
	xdata = xdata + "<img src= " + fileNameWithoutExtension(tnfile) + "3.png allt='test' width='128' height='96'>"
	xdata = xdata + "</div>"

	//-------------------------------------------------------------------------------------------------
	return xdata
}
