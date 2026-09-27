package shapecalculator

import (
	"errors"
	"fmt"
)

var(
	ErrInvalidDimensions = errors.New("all dimensions must be greater than zero")
	ErrInvalidFactor = errors.New("scale factor must be greater than zero")
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Scaler interface {
	Scale(factor float64) error
}

func Describe(s Shape) string {
    switch shape := s.(type) {
    case *Circle:
        return shape.String()
    case *Rectangle:
        return shape.String()
    case *Triangle:
        return shape.String()
    default:
        return fmt.Sprintf("Unknown shape: %T", s)
    }
}