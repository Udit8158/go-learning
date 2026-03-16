package shape

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// This is a method on the Rectangle struct that calculates its area.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// func AreaCircle(c Circle) float64 {
// 	return c.Radius * c.Radius * 3.141592653589793
// }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
