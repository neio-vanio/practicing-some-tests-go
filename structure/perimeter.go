package structure

import "math"

// Rectangle for calc
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle for calc
type Circle struct {
	Radius float64
}

// Triangle for calc
type Triangle struct {
	Base   float64
	Height float64
}

// Shape interface area
type Shape interface {
	Area() float64
}

// Area rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter of two values
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
