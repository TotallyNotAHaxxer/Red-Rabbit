package SUPER_JPG

import (
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"os"
)

func Generate(filename string) {
	parser := filename + ".jpg"
	i := image.NewRGBA(image.Rect(0, 0, 900, 1800))
	for k := 0; k < 100*200; k++ {
		o := 4 * k
		i.Pix[0+o] = uint8(rand.Intn(256))
		i.Pix[1+o] = uint8(rand.Intn(256))
		i.Pix[2+o] = uint8(rand.Intn(256))
		i.Pix[3+o] = 255
	}
	o, x := os.Create(parser)
	if x != nil {
		fmt.Println("<RR6> File I/O: Could not create image, got error -> ")
	} else {
		jpeg.Encode(o, i, nil)
		x = o.Close()
		if x != nil {
			fmt.Println("<RR6> File I/O: Could not encode the datam, ")
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Generated image -> ", parser)
		}
	}
}
