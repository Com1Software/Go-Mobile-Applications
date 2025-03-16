package main

import (
	"image"
	"image/color"
	"image/draw"
	"unsafe"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/gl"
)

var glctx gl.Context
var textTexture gl.Texture

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageVisible {
					glctx, _ = e.DrawContext.(gl.Context)
					initGL()
					loadBitmapFont()
				} else {
					glctx = nil
				}
			case paint.Event:
				if glctx != nil {
					drawFrame()
					a.Publish()
				}
			}
		}
	})
}

func initGL() {
	glctx.ClearColor(0.0, 0.0, 0.0, 1.0) // Black background
}

// Creates and loads a dummy bitmap font texture
func loadBitmapFont() {
	img := createTextImage() // Generate the image with text

	// Bind and load the texture into OpenGL
	textTexture = glctx.CreateTexture()
	glctx.BindTexture(gl.TEXTURE_2D, textTexture)

	glctx.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, img.Bounds().Dx(), img.Bounds().Dy(),
		gl.RGBA, gl.UNSIGNED_BYTE, imgToBytes(img))

	// Set texture parameters
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
}

// Creates an image containing "Hello, World!"
func createTextImage() *image.RGBA {
	width, height := 256, 64
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	// Draw white text as rectangles (simplified)
	addDummyText(img, "Hello, World!", 10, 30)

	return img
}

func addDummyText(img *image.RGBA, text string, x, y int) {
	// Render each character as a white rectangle
	charWidth, charHeight := 8, 16
	for i := 0; i < len(text); i++ {
		for row := 0; row < charHeight; row++ {
			for col := 0; col < charWidth; col++ {
				img.Set(x+i*charWidth+col, y+row, color.White)
			}
		}
	}
}

// Renders the entire frame
func drawFrame() {
	glctx.Clear(gl.COLOR_BUFFER_BIT)

	// Bind the texture containing the text
	glctx.BindTexture(gl.TEXTURE_2D, textTexture)

	// Draw the texture as a quad
	drawTexturedQuad(50, 100, 256, 64) // Adjust coordinates and dimensions as needed
}

// Draws a textured quad
func drawTexturedQuad(x, y, w, h float32) {
	vertices := []float32{
		x, y, 0, 0, // Top-left
		x + w, y, 1, 0, // Top-right
		x, y + h, 0, 1, // Bottom-left
		x + w, y + h, 1, 1, // Bottom-right
	}

	vertexBytes := float32ToBytes(vertices)

	buffer := glctx.CreateBuffer()
	glctx.BindBuffer(gl.ARRAY_BUFFER, buffer)
	glctx.BufferData(gl.ARRAY_BUFFER, vertexBytes, gl.STATIC_DRAW)

	glctx.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}

// Converts an image to a byte array for OpenGL
func imgToBytes(img *image.RGBA) []byte {
	return img.Pix
}

// Converts []float32 to []byte
func float32ToBytes(data []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), len(data)*4)
}
