package canvas

import (
	"bufio"
	d "canvas/internal/domain"
	"io"
	"strings"
)

type Application struct {
	Reader  io.Reader
	Printer io.Writer
}

func NewApplication(reader io.Reader, printer io.Writer) *Application {
	return &Application{reader, printer}
}

func (a *Application) Start() {
	scanner := bufio.NewScanner(a.Reader)
	for {
		if scanner.Scan() {
			command := scanner.Text()
			if command == "quit" {
				return
			}

			tokens := strings.Split(command, " ")
			width := tokens[1]
			height := tokens[2]

			canvas, err := d.NewCanvas(width, height)
			if err != nil {
				panic(err)
			}

			a.Printer.Write([]byte(canvas.ToText()))
		}
	}
}
