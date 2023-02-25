package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("should return perimeter value", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{
			shape: Rectangle{12, 6},
			want:  72,
		},
		{
			shape: Circle{10},
			want:  314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}

func TestCircle(t *testing.T) {
	t.Run("should return circle value", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %2f", got, want)
		}
	})
}
