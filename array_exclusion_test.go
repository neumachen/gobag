package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceExclusionStrings(t *testing.T) {
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
	expectedLeft := []string{"a"}
	expectedRight := []string{"d"}

	left, right := SliceExclusionStrings(a, b)
	require.True(t, testEqSliceStrs(expectedLeft, left), expectedLeft, left)
	require.True(t, testEqSliceStrs(expectedRight, right), expectedRight, right)
}

func TestSliceExclusionInts(t *testing.T) {
	a := []int{
		1,
		2,
		3,
	}
	b := []int{
		3,
		5,
		6,
	}
	expectedLeft := []int{1, 2}
	expectedRight := []int{5, 6}

	left, right := SliceExclusionInts(a, b)
	require.True(t, testEqSliceInts(expectedLeft, left), expectedLeft, left)
	require.True(t, testEqSliceInts(expectedRight, right), expectedRight, right)
}
