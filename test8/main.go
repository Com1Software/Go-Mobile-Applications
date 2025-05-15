package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
	b.bgRect.FillColor = color.RGBA{0, 255, 0, 255} // Change to green
	b.bgRect.Refresh()                              // Update UI
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Custom Button")

	btn := NewCustomButton("Click Me", color.RGBA{255, 0, 0, 255}, func() {
		println("Button clicked!")
	})

	myWindow.SetContent(container.NewVBox(btn))
	myWindow.ShowAndRun()
}
