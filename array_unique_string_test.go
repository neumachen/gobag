package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceUniqStrings(t *testing.T) {
	slice := []string{
		"a",
		"a",
		"a",
		"a",
		"a",
	}

	result := SliceUniqStrings(slice)
	require.Equal(t, 1, len(result))
}
