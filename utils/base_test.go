package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDieIf(t *testing.T) {
	err := errors.New("Hello")
	assert.PanicsWithError(t, "Hello", func() { DieIf(err) })

	assert.NotPanics(t, func() { DieIf(nil) })
}
