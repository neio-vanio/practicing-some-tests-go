package structure

import (
	"testing"
)

func TestArea(t *testing.T) {

	verifyArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()

		if result != expected {
			t.Errorf("result %.2f expected %.2f", result, expected)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.00, 6.00}
		expected := 72.0
		verifyArea(t, rectangle, expected)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793
		verifyArea(t, circle, expected)
	})

	testsArea := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{12.00, 6.00}, expected: 72.00},
		{name: "Circle", shape: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.expected {
				t.Errorf("%#v result %.2f expected %.2f", tt.name, result, tt.expected)
			}
		})
	}

}
