package structsmethodsinterface

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	circle := Circle{10.0}
	triangle := Triangle{10.0, 10.0}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{rectangle, 100.0},
		{circle, 314.16},
		{triangle, 50.0},
	}

	for _, tt := range areaTests {
		t.Run("Shape area", func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want

			// assert.Equal(t, want, got)
			assert.Equal(t, want, got, "%v should be equal %v", tt.want, got)
		})
	}
}
