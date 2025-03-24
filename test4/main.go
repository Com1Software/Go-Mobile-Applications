package main

import (
	"os"

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
	w := a.NewWindow("Two Buttons with Memo")

	// Create the text memo field with multiline and larger height
	memo := widget.NewEntry()
	memo.SetPlaceHolder("Enter your memo here...")
	memo.MultiLine = true               // Enable multiline for larger text fields
	memo.Resize(fyne.NewSize(400, 100)) // Adjust the height (4x the default)

	// Create the "Say Hello" button
	helloButton := widget.NewButton("Say Hello", func() {
		// Display the value from the memo field in the dialog box
		dialog.ShowInformation("Hello", "Hello, "+memo.Text, w)
	})

	// Create the "Exit" button
	exitButton := widget.NewButton("Exit", func() {
		os.Exit(0)
	})

	// Set the content of the window
	w.SetContent(container.NewVBox(
		memo,        // Add the memo field
		helloButton, // Add the "Say Hello" button
		exitButton,  // Add the "Exit" button
	))

	// Resize the window to make it larger
	w.Resize(fyne.NewSize(400, 300))

	// Show the window and run the application
	w.ShowAndRun()
}
