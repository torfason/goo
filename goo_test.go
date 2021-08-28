package goo

import (
	"errors"
	"testing"

	"github.com/torfason/goo/assert"
)

func TestMain(t *testing.T) {

	t.Run("assert.EqualInt", func(t *testing.T) {
		assert.EqualInt(t, 1, 1, "A message is optional.")
	})

	t.Run("assert.Error", func(t *testing.T) {
		assert.Error(t, errors.New("an error"), "A message is optional.")
	})

	// Test two ways of creating a slice
	// AKA: make() is never needed (except for pre-allocating memory)
	t.Run("assert.EqualSliceString", func(t *testing.T) {

		// Init empty
		x := []string{}            // With curlies
		y := make([]string, 0, 99) // With make(): cap is irrelevant for equality
		assert.EqualSliceString(t, x, y)

		// Init with elements, or append
		x = []string{"zero", "one", "two"}
		y = append(y, "zero", "one", "two")
		assert.EqualSliceString(t, x, y)

		// Append one slice to another
		z := append(x, y...)
		w := append(y, x...)
		assert.EqualSliceString(t, z, w)

	})

	// Test two ways of creating a map
	// AKA: make() is never needed (except for pre-allocating memory)
	t.Run("assert.EqualMapStringString", func(t *testing.T) {

		// Init empty
		x := map[string]string{}         // With curlies
		y := make(map[string]string, 99) // With make(): cap is irrelevant for equality
		assert.EqualMapStringString(t, x, y)

		// Init with elements and/or add after init
		x = map[string]string{"a": "Alpha", "b": "Beta"}
		y["a"] = "Alpha"
		y["b"] = "Beta"
		assert.EqualMapStringString(t, x, y)

		// Assign "" is (almost) equivalent to delete() !
		x["b"] = ""
		delete(y, "b")
		assert.EqualString(t, x["b"], y["b"])

		// But they can be distinguished with the ok param
		xVal, xOK := x["b"]
		yVal, yOK := y["b"]
		assert.EqualString(t, xVal, yVal)
		assert.True(t, xOK)
		assert.False(t, yOK)

	})
}
