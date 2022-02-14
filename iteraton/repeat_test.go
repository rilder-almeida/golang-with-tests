package iteraton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat char for n times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assert.Equal(t, want, got)
	})
}
