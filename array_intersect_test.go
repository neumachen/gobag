package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceIntersectStrings(t *testing.T) {
	a := []string{
		"a",
		"b",
		"c",
	}
	b := []string{
		"b",
		"c",
		"d",
	}
	expected := []string{"b", "c"}

	result := SliceIntersectStrings(a, b)
	require.True(t, len(result) == 2)
	require.True(t, testEqSliceStrs(expected, result), expected, result)
}

func TestSliceIntersectInts(t *testing.T) {
	a := []int{
		1,
		2,
		3,
	}
	b := []int{
		2,
		3,
		4,
	}
	expected := []int{2, 3}

	result := SliceIntersectInts(a, b)
	require.True(t, len(result) == 2)
	require.True(t, testEqSliceInts(expected, result), expected, result)
}
