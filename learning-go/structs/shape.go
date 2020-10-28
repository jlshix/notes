package structs

import "math"

// Shape interface for calculating area
type Shape interface {
	Area() float64
}

// Rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle shape
type Circle struct {
	Radius float64
}

// Area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle shape
type Triangle struct {
	Base   float64
	Height float64
}

// Area of Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
