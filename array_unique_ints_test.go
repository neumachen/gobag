package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceUniqInts(t *testing.T) {
	slice := []int{
		1,
		1,
		1,
		1,
		1,
	}

	result := SliceUniqInts(slice)
	require.Equal(t, 1, len(result))
}

func TestSliceUniqInt64s(t *testing.T) {
	slice := []int64{
		1,
		1,
		1,
		1,
		1,
	}

	result := SliceUniqInt64s(slice)
	require.Equal(t, 1, len(result))
}
