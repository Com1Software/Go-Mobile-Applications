package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"image/color"
	"log"
	"net"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	xip := fmt.Sprintf("%s", GetOutboundIP())
	port := "8080"
	a := app.New()
	w := a.NewWindow("Listening on " + xip + ":" + port)

	button1a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	})
	button2a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
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

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

type message struct {
	Controller string `xml:"controller"`
	DateTime   string `xml:"date_time"`
	Command    string `xml:"command"`
	Value      string `xml:"value"`
}

func ReadURL(url string) string {
	msg := &message{}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)

	for {
		line, erra := reader.ReadBytes('\n')
		if erra != nil {
			log.Fatal(erra) // Ensure correct error logging
		}
		xml.Unmarshal(line, &msg)
		break // Exit after the first read
	}

	// Construct the string with extracted values
	return msg.Controller + " " + msg.DateTime + " " + msg.Command + " " + msg.Value
}

type Agent struct {
	Notifier    chan []byte
	newuser     chan chan []byte
	closinguser chan chan []byte
	user        map[chan []byte]bool
}

func SSE() (agent *Agent) {
	agent = &Agent{
		Notifier:    make(chan []byte, 1),
		newuser:     make(chan chan []byte),
		closinguser: make(chan chan []byte),
		user:        make(map[chan []byte]bool),
	}
	go agent.listen()
	return
}

func (agent *Agent) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Error ", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	mChan := make(chan []byte)
	agent.newuser <- mChan
	defer func() {
		agent.closinguser <- mChan
	}()
	notify := req.Context().Done()
	go func() {
		<-notify
		agent.closinguser <- mChan
	}()
	for {
		fmt.Fprintf(rw, "%s", <-mChan)
		flusher.Flush()
	}

}

func (agent *Agent) listen() {
	for {
		select {
		case s := <-agent.newuser:
			agent.user[s] = true
		case s := <-agent.closinguser:
			delete(agent.user, s)
		case event := <-agent.Notifier:
			for userMChan, _ := range agent.user {
				userMChan <- event
			}
		}
	}

}
