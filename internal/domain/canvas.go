package domain

import (
	"errors"
	"strconv"
	"strings"
)

var InvalidCanvasWidth = errors.New("height should be positive integer")
var InvalidCanvasHeight = errors.New("height should be positive integer")

type Canvas struct {
	Width  int
	Height int
}

func NewCanvas(width string, height string) (*Canvas, error) {
	w, err := strconv.Atoi(width)
	if err != nil || w < 0 {
		return nil, InvalidCanvasWidth
	}
	h, err := strconv.Atoi(height)
	if err != nil || h < 0 {
		return nil, InvalidCanvasHeight
	}
	return &Canvas{w, h}, nil
}

func (c Canvas) ToText() string {
	sb := strings.Builder{}
	sb.WriteString(c.makeTopRow())
	sb.WriteString(c.makeMiddleRow())
	sb.WriteString(c.makeBottomRow())
	return sb.String()
}

func (c Canvas) makeHorizontalBorder() string {
	sb := strings.Builder{}
	for i := 0; i < c.Width+2; i++ {
		sb.WriteString("-")
	}
	return sb.String()
}

func (c Canvas) makeMiddleRow() string {
	sb := strings.Builder{}
	for h := 0; h < c.Height; h++ {
		sb.WriteString("|")
		for i := 0; i < c.Width; i++ {
			sb.WriteString(" ")
		}
		sb.WriteString("|")
		sb.WriteString("\n")
	}
	return sb.String()
}

func (c Canvas) makeTopRow() string {
	return c.makeHorizontalBorder() + "\n"
}

func (c Canvas) makeBottomRow() string {
	return c.makeHorizontalBorder()
}
