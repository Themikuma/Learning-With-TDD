package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name       string
		shape      Shape
		actualArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 3.0, Height: 4.0}, actualArea: 12.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, actualArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 5.0, Height: 6.0}, actualArea: 15},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.actualArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.actualArea)
			}
		})
	}
}
