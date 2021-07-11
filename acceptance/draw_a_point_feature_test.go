//+build progress

package acceptance

import (
	"canvas/pkg/strings"
	"testing"
)

func TestDrawAPoint(t *testing.T) {
	testCanvas := NewTestCanvas(t)
	testCanvas.PrepareCreateCommand(20, 4)
	testCanvas.PrepareDrawPointCommand(0, 0)
	testCanvas.PrepareDrawPointCommand(2, 3)
	testCanvas.PrepareQuitCommand()
	testCanvas.Run()

	expected := strings.TrimMargin(`
	----------------------
	|                    |
	|                    |
	|                    |
	|                    |
	----------------------
	----------------------
	|x                   |
	|                    |
	|                    |
	|                    |
	----------------------
	----------------------
	|x                   |
	|                    |
	|                    |
	|  x                 |
	----------------------	
`)

	testCanvas.IsLastCanvasLike(expected)
}
