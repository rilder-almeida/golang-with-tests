package maps

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assert.Error(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add pair on dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		assert.NoError(t, err)

		got, err2 := dictionary.Search("test")
		want := "this is just a test"
		assert.Equal(t, want, got)
		assert.NoError(t, err2)
	})

	t.Run("add a existed key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "this is just a test")

		want := ErrWordExists.Error()

		assert.EqualError(t, err, want)
	})

	t.Run("add a empty value", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "")

		want := ErrEmpytValue.Error()

		assert.EqualError(t, err, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update a pair", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("test", "this not is just a test")
		assert.NoError(t, err)

		want := "this not is just a test"
		got, err2 := dictionary.Search("test")

		assert.Equal(t, want, got)
		assert.NoError(t, err2)
	})

	t.Run("update a non-existed key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "this is just a test")

		want := ErrNotFound.Error()

		assert.EqualError(t, err, want)
	})

	t.Run("update a empty value", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("test", "")

		want := ErrEmpytValue.Error()

		assert.EqualError(t, err, want)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete a pair", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Delete("test")
		assert.NoError(t, err)

		_, err2 := dictionary.Search("test")
		assert.Equal(t, ErrNotFound, err2)
	})

	t.Run("delete a non-existed key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")

		want := ErrNotFound.Error()

		assert.EqualError(t, err, want)
	})
}
