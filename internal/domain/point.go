package domain

import "strconv"

type Point struct {
	x int
	y int
}

func (p Point) DrawOn(colorable Colorable) error {
	return colorable.Color(p.x, p.y)
}

func NewPoint(xCoordinate string, yCoordinate string) Point {
	x, _ := strconv.Atoi(xCoordinate)
	y, _ := strconv.Atoi(yCoordinate)
	return Point{x, y}
}
