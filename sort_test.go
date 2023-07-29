package gqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSort(t *testing.T) {
	assert.Equal(t, "Hello World", ParseSort())
}
