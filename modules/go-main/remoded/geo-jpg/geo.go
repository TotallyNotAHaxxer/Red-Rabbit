package SUPER_GEO

import (
	"bytes"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"strings"

	id_generation "main/modg/helpers/id-gen"

	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	"github.com/rwcarlsen/goexif/exif"
)

type FileSig struct {
	Sign       string
	SuffixFile string
	FileFormat string
}

var File_Sign = []FileSig{
	{`FFD8FF`, `*.jpg`, `JPEG image`},
	{`JFIF`, `*.jpg`, `JPEG image`},
}

func Grab_Location(filename string) {
	a := id_generation.ShortID(6)
	f, x := os.Open(filename)
	if x != nil {
		fmt.Println("\033[0;31m<RR6> Stego / Digital forensics module: Could not read filename, got error -> ", x)
	} else {
		defer f.Close()
		loc, x := exif.Decode(f)
		if x != nil {
			fmt.Println("\033[0;31m<RR6> Stego / Digital forensics module: Could not decode the image, this might be due to a few reasons -> ", x)
		} else {
			lat, lon, x := loc.LatLong()
			if x != nil {
				fmt.Println("\033[0;31m<RR6> Stego / Digital forensics module: Could not decode the image, this might be due to a few reasons such as the following")
				fmt.Println(" > There is no GPS permissions on the file")
				fmt.Println(" > This truly is not a JPEG / JPG file")
				fmt.Println(" > This is not a proper file and may be corrupted")
				fmt.Println(" > Or the premissions on the device were not set to grab or log GPS information into the ")
			} else {
				fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Pre Gen file   \033[38;5;21m: %s\n", a)
				fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File GPS Lat   \033[38;5;21m: %v\n", lat)
				fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File GPS lon   \033[38;5;21m: %v\n", lon)
				fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File GPS Map   \033[38;5;21m: https://www.google.com/maps/@%v,%v", lat, lon)
				Generate_MAP(a, lat, lon)
			}
		}
	}
}

func Test_Image(filename string) {
	f, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println("<RR6> Stego / Digital forensics module: Could not read filename, got error -> ", x)
	} else {
		finf, x := os.Stat(filename)
		if x != nil {
			log.Fatal("<RR6> Stego / Digital forensics module: Got error -> Could not stat file, [FATAL ERROR EXITING] -> ", x)
		} else {
			fmt.Println("<RR6> Stego / Digital forensics module -> Was able to sucessfully read the file and open it....")
			for _, v := range File_Sign {
				if strings.HasSuffix(filename, v.SuffixFile) || bytes.Contains(f, []byte(v.Sign)) {
					fmt.Printf("<RR6> Stego / Digital forensics module -> %s was detected as a %s moving on....\n", filename, v.FileFormat)
					fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File Sign      \033[38;5;21m: %s\n", v.Sign)
					fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File suffix    \033[38;5;21m: %s\n", v.SuffixFile)
					fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File Mode      \033[38;5;21m: %s\n", finf.Mode())
					fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File Name      \033[38;5;21m: %s\n", finf.Name())
					fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File Is a dir  \033[38;5;21m: %v\n", finf.IsDir())
					fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File Mode time \033[38;5;21m: %s\n", finf.ModTime())
					fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| File Size      \033[38;5;21m: %v\n", finf.Size())
					fmt.Printf("#######################################\n")
					Grab_Location(filename)
				} else {
					fmt.Printf("\033[0;31m<RR6> Stego / Digital forensics module: Could not detect %s as a JPEG/JPG image file format...\n", filename)
				}
			}
		}
	}
}

func Generate_MAP(name string, lat, long float64) {
	ctx := sm.NewContext()
	ctx.SetSize(600, 600)
	ctx.AddObject(sm.NewMarker(s2.LatLngFromDegrees(lat, long), color.RGBA{0xff, 0, 0, 0xff}, 16.0))
	i, x := ctx.Render()
	if x != nil {
		fmt.Println("<RR6> Stego / Digital forensics module: Could not make new image, got error -> ", x)
	} else {
		if x := gg.SavePNG(name, i); x != nil {
			fmt.Printf("\033[0;35m\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Geo Map Name   \033[38;5;21m: %v\n", name)
		}
	}
}
