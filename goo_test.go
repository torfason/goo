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
}
