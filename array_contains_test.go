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

func TestCollectionContains(t *testing.T) {
	t.Run("with strings", func(t *testing.T) {
		collection := []string{"hello", "there", "manbearpig"}
		result := SliceContains[string](collection, "hello")
		require.True(t, result)

		result = SliceContains[string](collection, "momo")
		require.False(t, result)
	})

	t.Run("with ints", func(t *testing.T) {
		collection := []int{1, 2, 3, 4, 5, 6}
		result := SliceContains[int](collection, 2)
		require.True(t, result)

		result = SliceContains[int](collection, 22)
		require.False(t, result)
	})
}
