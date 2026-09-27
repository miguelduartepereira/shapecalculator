package shapecalculator

import (
	"math"
	"testing"
)

func TestRectangle(t *testing.T) {
	rectangle, err := NewRectangle(4, 3)
	if err != nil {
		t.Fatalf("NewRectangle() error = %v", err)
	}

	if got := rectangle.Area(); got != 12 {
		t.Errorf("Area() = %v, want 12", got)
	}
	if got := rectangle.Perimeter(); got != 14 {
		t.Errorf("Perimeter() = %v, want 14", got)
	}
}

func TestNewRectangleRejectsInvalidDimensions(t *testing.T) {
	for _, dimensions := range [][2]float64{{0, 3}, {4, 0}, {-1, 3}, {4, -1}} {
		if _, err := NewRectangle(dimensions[0], dimensions[1]); err != ErrInvalidDimensions {
			t.Errorf("NewRectangle(%v, %v) error = %v, want %v", dimensions[0], dimensions[1], err, ErrInvalidDimensions)
		}
	}
}

func TestCircle(t *testing.T) {
	circle, err := NewCircle(2)
	if err != nil {
		t.Fatalf("NewCircle() error = %v", err)
	}

	assertClose(t, "Area()", circle.Area(), 12.56636)
	assertClose(t, "Perimeter()", circle.Perimeter(), 4*math.Pi)
}

func TestNewCircleRejectsInvalidRadius(t *testing.T) {
	for _, radius := range []float64{0, -1} {
		if _, err := NewCircle(radius); err != ErrInvalidDimensions {
			t.Errorf("NewCircle(%v) error = %v, want %v", radius, err, ErrInvalidDimensions)
		}
	}
}

func TestTriangle(t *testing.T) {
	triangle, err := NewTriangle(3, 4, 5)
	if err != nil {
		t.Fatalf("NewTriangle() error = %v", err)
	}

	if got := triangle.Area(); got != 6 {
		t.Errorf("Area() = %v, want 6", got)
	}
	if got := triangle.Perimeter(); got != 12 {
		t.Errorf("Perimeter() = %v, want 12", got)
	}
}

func TestDescribe(t *testing.T) {
	circle, err := NewCircle(2)
	if err != nil {
		t.Fatalf("NewCircle() error = %v", err)
	}
	rectangle, err := NewRectangle(4, 3)
	if err != nil {
		t.Fatalf("NewRectangle() error = %v", err)
	}
	triangle, err := NewTriangle(3, 4, 5)
	if err != nil {
		t.Fatalf("NewTriangle() error = %v", err)
	}

	tests := []struct {
		name  string
		shape Shape
		want  string
	}{
		{"circle", circle, "Circle(radius=2)"},
		{"rectangle", rectangle, "Rectangle(width=4, height=3)"},
		{"triangle", triangle, "Triangle(a=3, b=4, c=5)"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Describe(test.shape); got != test.want {
				t.Errorf("Describe() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestNewTriangleRejectsInvalidSides(t *testing.T) {
	for _, sides := range [][3]float64{{0, 4, 5}, {-1, 4, 5}} {
		if _, err := NewTriangle(sides[0], sides[1], sides[2]); err != ErrInvalidDimensions {
			t.Errorf("NewTriangle(%v, %v, %v) error = %v, want %v", sides[0], sides[1], sides[2], err, ErrInvalidDimensions)
		}
	}

	if _, err := NewTriangle(1, 2, 3); err != ErrImpossibleTriangle {
		t.Errorf("NewTriangle(1, 2, 3) error = %v, want %v", err, ErrImpossibleTriangle)
	}
}

func assertClose(t *testing.T, name string, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("%s = %v, want %v", name, got, want)
	}
}
