package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	result := 2 * (r.Height + r.Width)
	return result
}

func (r Rectangle) Area() float64 {
	result := r.Height * r.Width
	return result
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	result := math.Pi * c.Radius * c.Radius
	return result
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	result := (t.Base * t.Height) * 0.5
	return result
}
