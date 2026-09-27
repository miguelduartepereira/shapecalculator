package shapecalculator

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(width=%g, height=%g)", r.Width, r.Height)
}

func (r Rectangle) Scale(factor float64) error {
	if factor <= 0 {
		return ErrInvalidFactor
	}
	r.Width *= factor
	r.Height *= factor
	return nil
}

func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
		return nil, ErrInvalidDimensions
	}
	return &Rectangle{Width: width, Height: height}, nil
}