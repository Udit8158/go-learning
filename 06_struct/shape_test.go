package shape

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10, 10})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// interface for shapes that have an area
// interface only holds the shape of something, where stcuts holds the data
// pretty similar but they are different
type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("IN %#v: got %.2f want %.2f", shape, got, want)
		}
	}

	testCases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{Width: 10, Height: 20}, 200.0},
		{"circle", Circle{Radius: 10}, 314.1592653589793},
		{"triange", Triangle{Base: 10, Height: 5}, 259.0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
	// this why we should not calcualte the area here, instead we should
	// pass the shape to the checkArea function and let it calculate the area
	// for _,tt := range testCases {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		switch tt.shape.(type) {
	// 		case Rectangle:
	// 			checkArea(t, tt.shape.(Rectangle).Area(), tt.want)
	// 		case Circle:
	// 			checkArea(t, tt.shape.(Circle).Area(), tt.want)
	// 		}
	// 	})
	// }

}
