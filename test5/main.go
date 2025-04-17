package main

import (
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")
	//	fileName := "C:\\Go-Weather\\App\\Icon.png"
	//	fileName := "Icon.png"
	//	image := canvas.NewImageFromFile(fileName)
	file, _ := os.Open("Icon.png")
	defer file.Close()

	image := canvas.NewImageFromReader(file, "example-image")

	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}
