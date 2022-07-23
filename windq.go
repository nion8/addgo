package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Приложение")
	w.Resize(fyne.NewSize(400, 320))

	label := widget.NewLabel("первое число")
	entry1 := widget.NewEntry()

	labe2 := widget.NewLabel("Второе число")
	entry2 := widget.NewEntry()

	ansver := widget.NewLabel("")

	btn := widget.NewButton("Считать", func() {
		n1, err := strconv.ParseFloat(entry1.Text, 64)
		n2, er := strconv.ParseFloat(entry2.Text, 64)
	})

	w.SetContent(container.NewVBox(
		label,
		entry1,
		labe2,
		entry2,
		btn,
		ansver,
	))
	w.ShowAndRun()
}
