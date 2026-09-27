package shapecalculator

import (
	"fmt"
	"math"
)
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius=%g)", c.Radius)
}

func (c Circle) Scale(factor float64) error {
	if factor <= 0 {
		return ErrInvalidFactor
	}
	c.Radius *= factor
	return nil
}

func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, ErrInvalidDimensions
	}
	return &Circle{Radius: radius}, nil
}