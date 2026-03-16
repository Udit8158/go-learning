package shape

import "math"

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func Area(rec Rectangle) float64 {
	return rec.Width * rec.Height
}

func AreaCircle(circle Circle) float64 {
	return circle.Radius * circle.Radius * math.Pi
}
