package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {
	resultado := 2 + 2

	assert.Equal(t, 4, resultado)
	assert.NotEqual(t, 5, resultado)
	assert.True(t, resultado > 0)
}
