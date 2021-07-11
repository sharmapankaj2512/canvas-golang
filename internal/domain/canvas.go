package domain

import (
	"errors"
	"strconv"
	"strings"
)

var InvalidCanvasWidth = errors.New("height should be positive integer")
var InvalidCanvasHeight = errors.New("height should be positive integer")
var OutsideCanvasError = errors.New("point should be inside")

type Coordinates struct {
	x int
	y int
}

type Canvas struct {
	Width   int
	Height  int
	colored map[Coordinates]bool
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
	return &Canvas{w, h, make(map[Coordinates]bool)}, nil
}

func (c Canvas) Color(x int, y int) error {
	if x < 0 || x >= c.Width {
		return OutsideCanvasError
	}
	if y < 0 || y >= c.Height {
		return OutsideCanvasError
	}
	c.colored[Coordinates{x, y}] = true
	return nil
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
	for y := 0; y < c.Height; y++ {
		sb.WriteString("|")
		for x := 0; x < c.Width; x++ {
			_, ok := c.colored[Coordinates{x, y}]
			if ok {
				sb.WriteString("x")
			} else {
				sb.WriteString(" ")
			}
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
