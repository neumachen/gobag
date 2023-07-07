package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseString(t *testing.T) {
	s := "1234567890"
	reversedString := ReverseString(s)
	require.Equal(t, "0987654321", reversedString)
}
