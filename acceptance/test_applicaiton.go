package acceptance

import (
	"bytes"
	c "canvas/cmd/canvas"
	"canvas/pkg/assertions"
	"fmt"
	"io/ioutil"
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

func (a *TestCanvas) PrepareQuitCommand() {
	_, _ = fmt.Fprintln(a.reader, "quit")
}

func (a *TestCanvas) IsLastCanvasLike(expected string) {
	actual, _ := ioutil.ReadAll(a.printer)

	assertions.AssertEquals(a.t, expected, string(actual))
}

func (a *TestCanvas) Run() {
	c.NewApplication(a.reader, a.printer).Start()
}
