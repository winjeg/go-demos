package qrcode

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
)

func GenerateQrCode(content string) {
	err := qrcode.WriteColorFile(content, qrcode.High, 256,
		color.RGBA{255, 255, 200, 255}, color.RGBA{100, 100, 255, 255}, "E:/desktop/a.png")
	if err != nil {
		fmt.Println(err.Error())
	}
}
