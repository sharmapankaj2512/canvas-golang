package domain

import (
	mock "canvas/internal/mocks"
	"canvas/pkg/assertions"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestDrawPointWhenItLiesInsideTheCanvas(t *testing.T) {
	x, y := 2, 1
	ctrl := gomock.NewController(t)
	c := mock.NewMockColorable(ctrl)
	c.EXPECT().Color(x, y).Return(nil)

	point := Point{x, y}

	_ = point.DrawOn(c)
}

func TestReturnErrorWhenPointLiesOutsideTheCanvas(t *testing.T) {
	x, y := -1, -1
	ctrl := gomock.NewController(t)
	c := mock.NewMockColorable(ctrl)
	c.EXPECT().Color(x, y).Return(OutsideError)

	point := Point{x, y}

	assertions.AssertEquals(t, OutsideError, point.DrawOn(c))
}
