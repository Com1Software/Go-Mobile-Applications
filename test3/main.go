package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create the application
	a := app.New()

	// Create the main window
	w := a.NewWindow("Two Buttons")

	// Create the "Exit" button
	exitButton := widget.NewButton("Exit", func() {
		a.Quit()
	})

	// Create the "Say Hello" button
	helloButton := widget.NewButton("Say Hello", func() {
		dialog.ShowInformation("Hello", "Hello, World!", w)
	})

	// Set the content of the window
	w.SetContent(container.NewVBox(helloButton, exitButton))

	// Resize the window to make it larger
	w.Resize(fyne.NewSize(400, 300))

	// Show the window and run the application
	w.ShowAndRun()
}
