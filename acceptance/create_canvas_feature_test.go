//+build regression

package acceptance

import (
	"canvas/pkg/strings"
	"testing"
)

func TestCreateCanvas(t *testing.T) {
	testCanvas := NewTestCanvas(t)
	testCanvas.CreateCanvas(20, 4)
	testCanvas.Quit()
	testCanvas.Paint()

	expected := strings.TrimMargin(`
	----------------------
	|                    |
	|                    |
	|                    |
	|                    |
	----------------------`)

	testCanvas.IsLastCanvasLike(expected)
}
