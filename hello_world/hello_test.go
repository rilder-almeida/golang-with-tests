package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("say hallo to english people", func(t *testing.T) {
		got := Hello("Rilder", "English")
		want := "Hello, Rilder"

		assert.Equal(t, want, got)
	})
	t.Run("say hallo to spanish people", func(t *testing.T) {
		got := Hello("Rilder", "Spanish")
		want := "Hola, Rilder"

		assert.Equal(t, want, got)
	})
	t.Run("say hallo to french people", func(t *testing.T) {
		got := Hello("Rilder", "French")
		want := "Bonjour, Rilder"

		assert.Equal(t, want, got)
	})
	t.Run("say hallo to brasilian people", func(t *testing.T) {
		got := Hello("Rilder", "Portuguese")
		want := "Olá, Rilder"

		assert.Equal(t, want, got)
	})
	t.Run("say hallo to everybody", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, nobody"

		assert.Equal(t, want, got)
	})
	t.Run("no greeting for you", func(t *testing.T) {
		got := Hello("", "something")
		want := "Hello, nobody"

		assert.Equal(t, want, got)
	})
	t.Run("no greeting for you", func(t *testing.T) {
		got := Hello("something", "")
		want := "No greeting for you"

		assert.Equal(t, want, got)
	})

}
