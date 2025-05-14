package main

import (
	"image"
	"image/jpeg"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/disintegration/imaging"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Rotated Image")

	// Open the image file
	file, err := os.Open("test.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the image
	src, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	// Rotate the image by 90 degrees
	rotated := imaging.Rotate(src, 90, image.Transparent)

	// Convert to Fyne image
	fyneImage := canvas.NewImageFromImage(rotated)
	fyneImage.FillMode = canvas.ImageFillOriginal

	w.SetContent(fyneImage)
	w.ShowAndRun()
}
