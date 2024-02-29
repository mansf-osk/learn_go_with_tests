package structsmethodsinterfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	assertMessage(t, rectangle, got, want)
}

func TestArea(t *testing.T) {
	// Defines all tescases for areas in a struct
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 5.0}, hasArea: 50.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: math.Pi * 100},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		// Here new testcases can be added
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertMessage(t, shape, got, want)
	}

	// Runs all the testcases defined above
	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T) {
			checkArea(t, tc.shape, tc.hasArea)
		})
	}

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkArea(t, rectangle, 100.0)
	// })
	//
	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, math.Pi*100)
	// })
}

func assertMessage(t testing.TB, shape Shape, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got %g but want %g", shape, got, want)
	}
}
