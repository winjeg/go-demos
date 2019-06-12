package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func FyneUi() {
	ap := app.New()
	w := ap.NewWindow("Hello")
	a := &widget.Box{Children: []fyne.CanvasObject{
		&widget.Label{Text: "Hello Fyne!"},
		&widget.Button{Text: "Quit", OnTapped: func() {
			ap.Quit()
		}},
	}}
	gr := widget.NewGroup("Hello,world", a)
	w.SetContent(gr)
	w.ShowAndRun()
}
