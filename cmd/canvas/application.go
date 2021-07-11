package canvas

import (
	"bufio"
	d "canvas/internal/domain"
	"fmt"
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
	var canvas *d.Canvas
	var err error
	scanner := bufio.NewScanner(a.Reader)
	for {
		if scanner.Scan() {
			command := strings.ToLower(scanner.Text())
			if command == "quit" {
				return
			}

			if strings.HasPrefix(command, "c") {
				tokens := strings.Split(command, " ")
				width := tokens[1]
				height := tokens[2]

				canvas, err = d.NewCanvas(width, height)
				if err != nil {
					panic(err)
				}
			}

			if strings.HasPrefix(command, "p") {
				tokens := strings.Split(command, " ")
				x := tokens[1]
				y := tokens[2]

				point := d.NewPoint(x, y)
				point.DrawOn(canvas)
			}

			fmt.Fprintln(a.Printer, canvas.ToText())
			//a.Printer.Write([]byte())
		}
	}
}
