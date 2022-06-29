package SUPER_qr

import (
	"encoding/base64"
	"fmt"
	c "main/modg/colors"

	"github.com/skip2/go-qrcode"
)

var (
	X             error  // error handler
	Output        string // image size
	Standard_size = 256  // standard size for the image
	Data          []byte // Data byte conversion
	Data_         string // Data that will be hidden
)

func Generate(filename, data string) {
	Data_ = data
	Output = filename
	Data, X = qrcode.Encode(Data_, qrcode.Medium, Standard_size)
	if X != nil {
		fmt.Println(c.REDHB, "<RR6> QR Loader: Could not generate QR encoder, got error -> ", X, c.RET_RED)
	} else {
		encodedPngData := base64.StdEncoding.EncodeToString(Data)
		HTML_ := "<img src=\"data:image/png;base64," + encodedPngData + "\"/>"
		fmt.Println("<RR6> QR Loader: HTML TAG -> ", HTML_)
		X = qrcode.WriteFile(Output, qrcode.Medium, Standard_size, Data_)
		if X != nil {
			fmt.Println(c.REDHB, "<RR6> QR Loader: Could not generate QR code, got error -> ", X, c.RET_RED)
		} else {
			fmt.Println("<RR6> QR Loader: Generation was a sucess..")
		}
	}
}
