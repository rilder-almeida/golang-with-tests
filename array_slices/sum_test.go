package arrayslices

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("sum all numbers of array or slice", func(t *testing.T) {
		got := SumArray([]int{1, 2, 3, 4, 5})
		want := 15

		assert.Equal(t, want, got)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum multiple arrays oy alices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := 12

		assert.Equal(t, want, got)
	})
}

func TestSumTails(t *testing.T) {
	t.Run("sum tails of arrays", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		want := 11

		assert.Equal(t, want, got)
	})
	t.Run("sum tails of empty arrays", func(t *testing.T) {
		got := SumAll([]int{}, []int{})
		want := 0

		assert.Equal(t, want, got)
	})
}
