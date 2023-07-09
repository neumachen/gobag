package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceContains(t *testing.T) {
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

func TestSliceExclusion(t *testing.T) {
	t.Run("with ints", func(t *testing.T) {
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

		left, right := SliceExclusion[int](a, b)
		require.True(t, testEqSliceInts(expectedLeft, left), expectedLeft, left)
		require.True(t, testEqSliceInts(expectedRight, right), expectedRight, right)
	})

	t.Run("with strings", func(t *testing.T) {
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
	})
}

func TestSliceIntersection(t *testing.T) {
	t.Run("using ints", func(t *testing.T) {
		source := []int{
			1,
			2,
			3,
		}
		target := []int{
			2,
			3,
			4,
		}
		expected := []int{2, 3}

		result := SliceIntersection[int](source, target)
		require.True(t, len(result) == 2)
		require.True(t, testEqSliceInts(expected, result), expected, result)
	})
	t.Run("using strings", func(t *testing.T) {
		source := []string{
			"a",
			"b",
			"c",
		}
		target := []string{
			"b",
			"c",
			"d",
		}
		expected := []string{"b", "c"}

		result := SliceIntersection[string](source, target)
		require.True(t, len(result) == 2)
		require.True(t, testEqSliceStrs(expected, result), expected, result)
	})
}

func tesFilterUniqueElements[T comparable](t *testing.T, input []T, expected []T) {
	result := FilterUniqueElements[T](input)
	require.Equal(t, expected, result)
}

func TestFilterUniqueElements(t *testing.T) {
	tests := []struct {
		name      string
		assertion func(t *testing.T)
		input     interface{}
		expected  interface{}
	}{
		{
			name: "test with []int",
			assertion: func(t *testing.T) {
				testData := []struct {
					input    []int
					expected []int
				}{
					{
						input:    []int{1, 2, 3, 3, 4, 4, 5},
						expected: []int{1, 2, 3, 4, 5},
					},
					{
						input:    []int{1, 1, 1, 1, 1, 1, 1},
						expected: []int{1},
					},
				}

				for i := range testData {
					tesFilterUniqueElements[int](t, testData[i].input, testData[i].expected)
				}
			},
		},
		{
			name: "test with []string",
			assertion: func(t *testing.T) {
				testData := []struct {
					input    []string
					expected []string
				}{
					{
						input:    []string{"apple", "banana", "banana", "cherry", "apple"},
						expected: []string{"apple", "banana", "cherry"},
					},
					{
						input:    []string{"apple", "apple", "apple", "apple", "apple"},
						expected: []string{"apple"},
					},
				}

				for i := range testData {
					tesFilterUniqueElements[string](t, testData[i].input, testData[i].expected)
				}
			},
		},
		{
			name: "test with []float64",
			assertion: func(t *testing.T) {
				testData := []struct {
					input    []float64
					expected []float64
				}{
					{
						input:    []float64{1.2, 3.4, 5.6, 3.4, 7.8, 1.2},
						expected: []float64{1.2, 3.4, 5.6, 7.8},
					},
					{
						input:    []float64{1.2, 1.2, 1.2, 1.2, 1.2, 1.2},
						expected: []float64{1.2},
					},
				}

				for i := range testData {
					tesFilterUniqueElements[float64](t, testData[i].input, testData[i].expected)
				}
			},
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			test.assertion(t)
		})
	}
}
