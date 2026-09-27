package shapecalculator

import(
	"math"
	"errors"
	"fmt"
)

var (
	ErrImpossibleTriangle = errors.New("each side of the triangle must be shorter than the sum of the other two sides")
)

type Triangle struct {
	a   float64
	b   float64
	c   float64
}

func (t Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle(a=%g, b=%g, c=%g)", t.a, t.b, t.c)
}

func (t Triangle) Scale(factor float64) error {
	if factor <= 0 {
		return ErrInvalidFactor
	}
	t.a *= factor
	t.b *= factor
	t.c *= factor
	return nil
}

func NewTriangle(a, b, c float64) (*Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, ErrInvalidDimensions
	}
	if a >= b+c || b >= a+c || c >= a+b {
		return nil, ErrImpossibleTriangle
	}
	return &Triangle{a: a, b: b, c: c}, nil
}