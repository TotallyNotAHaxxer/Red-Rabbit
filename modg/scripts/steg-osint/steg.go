package steg

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"main/modg/system-runscript"
	"os"
	"strings"

	ec "main/modg/warnings"

	v "main/modg/colors"

	opc "main/modg/copt"

	constants "main/modg/imj"

	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/pflag"
)

var (
	rr6f  opc.RR6_options
	flags = pflag.FlagSet{SortFlags: false}
)

func Geo_loc(filename string) {
	f, err := os.Open(filename)
	ec.Che(err, "Could not open file", 1)

	x, err := exif.Decode(f)
	ec.Che(err, "DEBUG: ERR: FATAL: during running the decode function for EXIF", 1)

	lat, long, err := x.LatLong()
	if err != nil {
		fmt.Fprintf(os.Stderr, "LatLong: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stdout, f.Name())
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Image -> ", filename)
		fmt.Fprintf(os.Stdout, fmt.Sprintf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| lat:\t%v\n\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| long:\t%v", lat, long))
		fmt.Fprintf(os.Stdout, fmt.Sprintf("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Possible Location -> https://www.google.com/maps/@%v,%v", lat, long))
		fmt.Println("----------------------------------------------------------------------------------------------")
		fmt.Println("-> ++ > Generating MAP for geo location")
		Geo_map("GEO_LOCATION_MAP#0", lat, long)
		os.Exit(0)
	}
}

func Geo_map(mapname string, lat, lon float64) {
	ctx := sm.NewContext()
	ctx.SetSize(600, 600)
	ctx.AddObject(sm.NewMarker(s2.LatLngFromDegrees(lat, lon), color.RGBA{0xff, 0, 0, 0xff}, 16.0))
	img, err := ctx.Render()
	ec.Che(err, "ERR: Context could not render", 1)
	if err := gg.SavePNG(mapname, img); err != nil {
		fmt.Println("could not save or create map")
		panic(err)
	}
}

func Hex_dump(filename string) {
	f, err := os.Open(filename)
	ec.Che(err, "Could not open file", 1)
	defer f.Close()
	reader := bufio.NewReader(f)
	buf := make([]byte, 256)
	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Printf("%s", hex.Dump(buf))
	}
}

func File_sig(filename string) {
	for _, file := range filename {
		f, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "file: %v\n", err)
			continue
		}

		for _, val := range constants.MassSign {
			if strings.HasSuffix(filename, val.SuffixFile) || bytes.Contains(f, []byte(val.Sign)) {
				fmt.Print(v.CYN, "\n\n[INFO] DATA: ", v.BLU, system.FormatDate, "\t", file, " Was possibly detected as -> ", v.RED, val.FileFormat, "\n\n")
			}
		}
		break
	}
}
