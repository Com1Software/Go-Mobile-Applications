package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Multiple Images")

	// Create URIs for multiple images
	uri1 := storage.NewURI("http://com1software.com/a1597.jpg")
	uri2 := storage.NewURI("http://com1software.com/a1597.jpg")
	uri3 := storage.NewURI("http://com1software.com/a1597.jpg")

	// Create image objects
	image1 := canvas.NewImageFromURI(uri1)
	image2 := canvas.NewImageFromURI(uri2)
	image3 := canvas.NewImageFromURI(uri3)

	// Set fill modes if needed
	image1.FillMode = canvas.ImageFillOriginal
	image2.FillMode = canvas.ImageFillOriginal
	image3.FillMode = canvas.ImageFillOriginal

	// Arrange images in a 3-column grid
	imageGrid1 := container.NewGridWithColumns(3, image1, image2, image3)
	imageGrid2 := container.NewGridWithColumns(3, image1, image2, image3)
	imageGrid3 := container.NewGridWithColumns(3, image1, image2, image3)
	imageGrid4 := container.NewGridWithColumns(3, image1, image2, image3)
	imageGrid5 := container.NewGridWithColumns(3, image1, image2, image3)
	imageGrid6 := container.NewGridWithColumns(3, image1, image2, image3)
	imageGrid7 := container.NewGridWithColumns(3, image1, image2, image3)
	imageGrid8 := container.NewGridWithColumns(3, image1, image2, image3)

	multiGrid := container.NewVBox(imageGrid1, imageGrid2, imageGrid3, imageGrid4, imageGrid5, imageGrid6, imageGrid7, imageGrid8)
	// Set the content of the window
	w.SetContent(multiGrid)

	w.ShowAndRun()
}
