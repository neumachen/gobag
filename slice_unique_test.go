package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
