package integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	t.Run("add 2 numbers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assert.Equal(t, want, got)
	})
}
