package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Testing")

	button2a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	})

	button1a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button 1 clicked!")
		button2a.Tapped(&fyne.PointEvent{}) // Simulate button 2 being pressed
	})

	buttonContainer1a := container.NewGridWrap(fyne.NewSize(40, 40), button1a)
	buttonContainer2a := container.NewGridWrap(fyne.NewSize(40, 40), button2a)

	layoutContainera := container.NewHBox(buttonContainer1a, buttonContainer2a)

	w.Resize(fyne.NewSize(400, 300))
	w.SetContent(container.NewVBox(
		layoutContainera,
	))
	w.ShowAndRun()
}

// CustomButton creates a button with a colored background
type CustomButton struct {
	widget.BaseWidget
	text     string
	color    color.Color
	onTapped func()
	bgRect   *canvas.Rectangle
}

func NewCustomButton(text string, bgColor color.Color, onTapped func()) *CustomButton {
	btn := &CustomButton{text: text, color: bgColor, onTapped: onTapped}
	btn.ExtendBaseWidget(btn)
	return btn
}

// CreateRenderer defines the custom rendering for the button
func (b *CustomButton) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(b.text)
	b.bgRect = canvas.NewRectangle(b.color) // Use a modifiable background

	content := container.NewMax(b.bgRect, label)
	return widget.NewSimpleRenderer(content)
}

func (b *CustomButton) Tapped(_ *fyne.PointEvent) {
	if b.onTapped != nil {
		b.onTapped()
	}
	hit := false
	if hit {
		b.bgRect.FillColor = color.RGBA{255, 0, 0, 255} // Change to Red
	} else {
		b.bgRect.FillColor = color.RGBA{255, 255, 255, 255} // Change to Green
	}
	b.bgRect.Refresh() // Update UI
}
