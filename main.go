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
)

func main() {
	fmt.Printf("ffmpeg Parser - (c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)

	exefile := "/ffmpeg/bin/ffmpeg.exe"
	exefilea := "/ffmpeg/bin/ffprobe.exe"
	drive := "c"
	wdir := drive + ":/tmp/"
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
		xdata = xdata + "<H1>ffmpeg Parse for " + wdir + "</H1>"
		for _, file := range files {
			if ValidFileType(strings.ToLower(path.Ext(file.Name()))) {
				tfile := wdir + file.Name()
				tnfile := fixFileName(tfile)
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
				xdata = xdata + "  <A HREF='file:///" + tnfile + "'>  [ " + file.Name() + " ] <BR> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "1.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "2.png" + "  ALT=error> <IMG SRC=" + fileNameWithoutExtension(tnfile) + "3.png" + "  ALT=error> </A><BR> "
				//-------------------------------------------------------------------------------------------------
				bfile := "tmp.bat"
				bdata := []byte(exefilea + " -i " + tnfile + " -show_entries stream=width,height -of csv=" + fmt.Sprintf("%q", "p=0") + ">tmp.txt")
				err := os.WriteFile(bfile, bdata, 0644)
				cmd = exec.Command(bfile)
				if err = cmd.Run(); err != nil {
					fmt.Printf("Command %s \n Error: %s\n", cmd, err)
				}
				dat := []byte("")
				dat, err = os.ReadFile("tmp.txt")
				tdata := string(dat)
				tmp := strings.Split(tdata, ",")
				xdata = xdata + "Frame width " + tmp[0] + "<BR>"
				xdata = xdata + "Frame height " + tmp[1] + "<BR>"

				//-------------------------------------------------------------------------------------------------
				bdata = []byte(exefilea + " -i " + tnfile + " -show_entries format=duration -v quiet -of csv >tmp.txt")
				err = os.WriteFile(bfile, bdata, 0644)
				cmd = exec.Command(bfile)
				if err = cmd.Run(); err != nil {
					fmt.Printf("Command %s \n Error: %s\n", cmd, err)
				}
				dat = []byte("")
				dat, err = os.ReadFile("tmp.txt")
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
				bdata = []byte(exefilea + " -i " + tnfile + " -show_entries stream=r_frame_rate  -of xml" + ">tmp.xml")
				err = os.WriteFile(bfile, bdata, 0644)
				cmd = exec.Command(bfile)
				if err = cmd.Run(); err != nil {
					fmt.Printf("Command %s \n Error: %s\n", cmd, err)
				}
				dat = []byte("")
				dat, err = os.ReadFile("tmp.xml")
				tdata = string(dat)
				tagdata := ParseXMLTag(tdata, "<stream r_frame_rate=", 1)
				//				fmt.Printf("Tag data %s\n", tagdata)
				xdata = xdata + "Frames per second  " + tagdata + " <BR>"
				//-------------------------------------------------------------------------------------------------

				xdata = xdata + "<BR><BR>"

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

func ParseXMLTag(xml string, tag string, pos int) string {
	rtn := ""
	chr := ""
	ton := false
	ttag := ""
	tdata := ""
	do := false
	dopos := 0
	vpos := 1
	v1 := 0
	v2 := 0
	v3 := 0
	v4 := 0

	for x := 0; x < len(xml); x++ {
		chr = xml[x : x+1]
		if chr == "<" {
			ton = true
			ttag = ""
		}
		if chr == ">" {
			ton = false
		}
		if ton {
			ttag = ttag + chr
		}
		if ttag == tag {
			tdata = ""
			for xx := x; xx < len(xml); xx++ {
				chr = xml[xx : xx+1]
				fmt.Printf(chr)
				if chr == ">" {
					xx = len(xml)
				}
				if strings.Contains(chr, `"`) && do == false {
					do = true
					dopos = xx

				}

				//if strings.Contains(chr, `"`) && do == true {
				//	fmt.Printf("Quote True %d \n", vpos)
				//	do = false
				//	dopos = xx
				//	switch {
				//	case vpos == 1:
				//		v1, _ = strconv.Atoi(tdata)
				//	case vpos == 2:
				//		v2, _ = strconv.Atoi(tdata)
				//	case vpos == 3:
				//		v3, _ = strconv.Atoi(tdata)
				//	case vpos == 4:
				//		v4, _ = strconv.Atoi(tdata)
				//	}
				//				}

				if strings.Contains(chr, `/`) && do == true {
					do = false
					switch {
					case vpos == 1:
						v1, _ = strconv.Atoi(tdata)
					case vpos == 2:
						v2, _ = strconv.Atoi(tdata)
					case vpos == 3:
						v3, _ = strconv.Atoi(tdata)
					case vpos == 4:
						v4, _ = strconv.Atoi(tdata)
					}

					tdata = ""
					vpos++
					dopos = xx
				}
				if do && xx > dopos {
					tdata = tdata + chr
				}
			}
			fmt.Printf("\n %d %d %d %d \n", v1, v2, v3, v4)
			rtn = tdata

		}
	}

	//numerator, _ := strconv.Atoi(tmpa[0])
	//denominator, _ := strconv.Atoi(tmpa[1])
	//				fmt.Println(numerator)
	//				fmt.Println(denominator)
	//				if denominator > 0 {
	//					fps := numerator / denominator
	//					fmt.Printf("The frame rate is: %.2f fps\n", fps)
	//		}

	return rtn
}
