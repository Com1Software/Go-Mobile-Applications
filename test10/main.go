package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Testing")

	resultLabel := widget.NewLabel("Input will appear here")

	button := NewCustomButton("Open Dialog", color.RGBA{0, 0, 255, 255}, w, resultLabel)

	buttonContainer := container.NewGridWrap(fyne.NewSize(100, 40), button)

	w.Resize(fyne.NewSize(400, 300))
	w.SetContent(container.NewVBox(
		buttonContainer,
		resultLabel,
	))
	w.ShowAndRun()
}

// CustomButton creates a button with a colored background
type CustomButton struct {
	widget.BaseWidget
	text        string
	color       color.Color
	bgRect      *canvas.Rectangle
	w           fyne.Window
	resultLabel *widget.Label
}

func NewCustomButton(text string, bgColor color.Color, w fyne.Window, resultLabel *widget.Label) *CustomButton {
	btn := &CustomButton{text: text, color: bgColor, w: w, resultLabel: resultLabel}
	btn.ExtendBaseWidget(btn)
	return btn
}

// CreateRenderer defines the custom rendering for the button
func (b *CustomButton) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(b.text)
	b.bgRect = canvas.NewRectangle(b.color)

	content := container.NewMax(b.bgRect, label)
	return widget.NewSimpleRenderer(content)
}

func (b *CustomButton) Tapped(_ *fyne.PointEvent) {
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()

	dialog.ShowForm(
		"Enter Details",
		"Submit",
		"Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("First Input", entry1),
			widget.NewFormItem("Second Input", entry2),
		},
		func(response bool) {
			if response {
				b.resultLabel.SetText("First: " + entry1.Text + " | Second: " + entry2.Text)
			}
		},
		b.w,
	)
}
