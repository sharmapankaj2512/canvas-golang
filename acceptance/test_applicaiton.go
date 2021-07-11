package acceptance

import (
	"bytes"
	c "canvas/cmd/canvas"
	"canvas/pkg/assertions"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

type TestCanvas struct {
	reader  *bytes.Buffer
	printer *bytes.Buffer
	t       *testing.T
}

func NewTestCanvas(t *testing.T) *TestCanvas {
	reader := bytes.NewBuffer(make([]byte, 1024))
	printer := bytes.NewBuffer(make([]byte, 1024))
	reader.Reset()
	printer.Reset()
	return &TestCanvas{
		reader,
		printer,
		t,
	}
}

func (a *TestCanvas) PrepareCreateCommand(width int, height int) {
	_, _ = fmt.Fprintln(a.reader, fmt.Sprintf("C %d %d", width, height))
}

func (a *TestCanvas) PrepareDrawPointCommand(x int, y int) {
	_, _ = fmt.Fprintln(a.reader, fmt.Sprintf("P %d %d", x, y))
}

func (a *TestCanvas) PrepareQuitCommand() {
	_, _ = fmt.Fprintln(a.reader, "quit")
}

func (a *TestCanvas) Run() {
	c.NewApplication(a.reader, a.printer).Start()
}

func (a *TestCanvas) IsLastCanvasLike(expected string) {
	actual, _ := ioutil.ReadAll(a.printer)

	assertions.AssertEquals(a.t, expected, strings.TrimSpace(string(actual)))
}
