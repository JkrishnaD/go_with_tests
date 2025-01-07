package main

import "testing"

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %2f but expected %2f", got, want)
// 		}
// 	}
// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		expected := 100.0
// 		checkArea(t, rectangle, expected)
// 	})
// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		expected := 100.0
// 		checkArea(t, circle, expected)
// 	})
// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "rectangle", shape: Rectangle{10.0, 10.0}, area: 100.0},
		{name: "circle", shape: Circle{10}, area: 00},
		{name: "trainlge", shape: Traingle{5.0, 5.0}, area: 12.5},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.area {
				t.Errorf("for %#v got %g want %g", tt.shape, got, tt.area)
			}
		})
	}
}

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Perimeter(rectangle)
// 	expected := 40.0

// 	if got != expected {
// 		t.Errorf("got %2f but expected %2f", got, expected)
// 	}
// }
