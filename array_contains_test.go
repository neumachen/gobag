package gobag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayContainsStr(t *testing.T) {
	array := []string{"hello", "there", "manbearpig"}

	b := ArrayContainsStr(array, "hello")
	assert.True(t, b)

	b = ArrayContainsStr(array, "momo")
	assert.False(t, b)
}
