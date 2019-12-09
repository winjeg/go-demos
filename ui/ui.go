package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/czxichen/otpauth"
	"time"
)

func getCode() (string, int64) {
	now := time.Now().Unix()
	now -= 30
	code, secs, err := otpauth.GenerateCode("R2SROZ7DE7O632PA", now)
	if err == nil {
		return fmt.Sprintf("%06d\n", code), secs
	}
	return "error occurred", 0
}

func FyneUi() {

	ap := app.New()
	w := ap.NewWindow("Hello")
	c, s := getCode()
	fmt.Println(c)
	fmt.Println(time.Now().Unix() - s)
	a := &widget.Box{Children: []fyne.CanvasObject{
		&widget.Label{Text: c},

		&widget.Button{Text: "Quit", OnTapped: func() {
			ap.Quit()
		}},
	}}
	w.SetContent(a)
	w.ShowAndRun()

}
