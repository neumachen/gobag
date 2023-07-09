package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayContainsStr(t *testing.T) {
	array := []string{"hello", "there", "manbearpig"}

	b := ArrayContainsStr(array, "hello")
	require.True(t, b)

	b = ArrayContainsStr(array, "momo")
	require.False(t, b)
}
