//+build regression

package acceptance

import (
	"canvas/pkg/strings"
	"testing"
)

func TestCreateCanvas(t *testing.T) {
	testCanvas := NewTestCanvas(t)
	testCanvas.PrepareCreateCommand(20, 4)
	testCanvas.PrepareQuitCommand()
	testCanvas.Run()

	expected := strings.TrimMargin(`
	----------------------
	|                    |
	|                    |
	|                    |
	|                    |
	----------------------`)

	testCanvas.IsLastCanvasLike(expected)
}
