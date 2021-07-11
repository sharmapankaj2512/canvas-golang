package domain

import "errors"

var OutsideError = errors.New("point should be inside")

type Colorable interface {
	Color(x int, y int) error
}
