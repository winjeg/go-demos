package qrcode

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
)

func GenerateQrCode() {
	err := qrcode.WriteColorFile("你是我的全部", qrcode.High, 256,
		color.RGBA{255, 100, 100, 255}, color.RGBA{100, 100, 255, 255}, "E:/desktop/a.png")
	if err != nil {
		fmt.Println(err.Error())
	}
}
