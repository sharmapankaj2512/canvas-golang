//+build progress

package acceptance

import (
	"canvas/pkg/strings"
	"testing"
)

func TestDrawAPoint(t *testing.T) {
	testCanvas := NewTestCanvas(t)
	testCanvas.CreateCanvas(20, 4)
	testCanvas.DrawPoint(0, 0)
	testCanvas.DrawPoint(2, 3)
	testCanvas.Quit()
	testCanvas.Paint()

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
