//+build unit

package domain

import (
	"canvas/pkg/assertions"
	"canvas/pkg/strings"
	"testing"
)

func TestReturnsErrorWhenNonPositiveIntegerIsPassedAsWidth(t *testing.T) {
	tests := []struct {
		width string
		err   error
	}{
		{"hello", InvalidCanvasWidth},
		{"-1", InvalidCanvasWidth},
		{"2.2", InvalidCanvasWidth},
	}
	for _, tt := range tests {
		t.Run(tt.width, func(t *testing.T) {
			_, err := NewCanvas(tt.width, "0")

			assertions.AssertEquals(t, tt.err, err)
		})
	}
}

func TestReturnsNilAsErrorWhenPositiveIntegerIsPassedAsWidth(t *testing.T) {
	tests := []struct {
		width string
		err   error
	}{
		{"0", nil},
		{"12", nil},
	}
	for _, tt := range tests {
		t.Run(tt.width, func(t *testing.T) {
			_, err := NewCanvas(tt.width, "0")

			assertions.AssertEquals(t, tt.err, err)
		})
	}
}

func TestReturnsErrorWhenNonPositiveIntegerIsPassedAsHeight(t *testing.T) {
	tests := []struct {
		height string
		err    error
	}{
		{"hello", InvalidCanvasHeight},
		{"-1", InvalidCanvasHeight},
		{"2.2", InvalidCanvasHeight},
	}
	for _, tt := range tests {
		t.Run(tt.height, func(t *testing.T) {
			_, err := NewCanvas("0", tt.height)

			assertions.AssertEquals(t, tt.err, err)
		})
	}
}

func TestReturnsNilAsErrorWhenPositiveIntegerIsPassedAsHeight(t *testing.T) {
	tests := []struct {
		height string
		err    error
	}{
		{"0", nil},
		{"12", nil},
	}
	for _, tt := range tests {
		t.Run(tt.height, func(t *testing.T) {
			_, err := NewCanvas("0", tt.height)

			assertions.AssertEquals(t, tt.err, err)
		})
	}
}

func TestReturnsEmptyCanvasAsTextForWidthZeroHeightZero(t *testing.T) {
	canvas, _ := NewCanvas("0", "0")
	expected := strings.TrimMargin(`
	--
	--`)

	assertions.AssertEquals(t, expected, canvas.ToText())
}

func TestReturnsSingleSpaceCanvasAsTextForWidthOneHeightOne(t *testing.T) {
	canvas, _ := NewCanvas("1", "1")
	expected := strings.TrimMargin(`
		---
		| |
		---`)

	assertions.AssertEquals(t, expected, canvas.ToText())
}

func TestReturnsFourSpaceCanvasAsTextForWidthTwoHeightTwo(t *testing.T) {
	canvas, _ := NewCanvas("2", "2")
	expected := strings.TrimMargin(`
		----
		|  |
		|  |
		----`)

	assertions.AssertEquals(t, expected, canvas.ToText())
}

func TestReturnsSixSpaceCanvasAsTextForWidthThreeHeightTwo(t *testing.T) {
	canvas, _ := NewCanvas("3", "2")
	expected := strings.TrimMargin(`
		-----
		|   |
		|   |
		-----`)

	assertions.AssertEquals(t, expected, canvas.ToText())
}

func TestReturnsColoredCanvasAsText(t *testing.T) {
	canvas, _ := NewCanvas("3", "2")
	canvas.Color(0, 0)
	expected := strings.TrimMargin(`
		-----
		|x  |
		|   |
		-----`)

	assertions.AssertEquals(t, expected, canvas.ToText())
}

func TestReturnsErrorOnColorWhenXCoordinateIsLessThanZero(t *testing.T) {
	canvas, _ := NewCanvas("3", "2")

	assertions.AssertEquals(t, OutsideCanvasError, canvas.Color(-1, 0))
}

func TestReturnsErrorOnColorWhenXCoordinateIsGreaterThanMaxWidth(t *testing.T) {
	canvas, _ := NewCanvas("3", "2")

	assertions.AssertEquals(t, OutsideCanvasError, canvas.Color(-1, 4))
}

func TestReturnsErrorOnColorWhenYCoordinateIsLessThanZero(t *testing.T) {
	canvas, _ := NewCanvas("3", "2")

	assertions.AssertEquals(t, OutsideCanvasError, canvas.Color(0, -1))
}

func TestReturnsErrorOnColorWhenYCoordinateIsGreaterThanMaxHeight(t *testing.T) {
	canvas, _ := NewCanvas("3", "2")

	assertions.AssertEquals(t, OutsideCanvasError, canvas.Color(2, 3))
}
