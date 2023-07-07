package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendJSON(t *testing.T) {
	var emptyBytesA []byte
	emptyBytesB := []byte(nil)
	tests := []struct {
		originalBytes []byte
		bytesToAdd    []byte
		appendedBytes []byte
	}{
		{
			originalBytes: nil,
			bytesToAdd:    nil,
			appendedBytes: nil,
		},
		{
			originalBytes: []byte(`{}`),
			bytesToAdd:    []byte(`{}`),
			appendedBytes: nil,
		},
		{
			originalBytes: emptyBytesA,
			bytesToAdd:    emptyBytesA,
			appendedBytes: nil,
		},
		{
			originalBytes: emptyBytesB,
			bytesToAdd:    emptyBytesB,
			appendedBytes: nil,
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    []byte(`{}`),
			appendedBytes: []byte(`{"test":"originalBytes"}`),
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    nil,
			appendedBytes: []byte(`{"test":"originalBytes"}`),
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    emptyBytesA,
			appendedBytes: []byte(`{"test":"originalBytes"}`),
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    nil,
			appendedBytes: []byte(`{"test":"originalBytes"}`),
		},
		{
			originalBytes: []byte(`{}`),
			bytesToAdd:    []byte(`{"test":"bytesToAdd"}`),
			appendedBytes: []byte(`{"test":"bytesToAdd"}`),
		},
		{
			originalBytes: nil,
			bytesToAdd:    []byte(`{"test":"bytesToAdd"}`),
			appendedBytes: []byte(`{"test":"bytesToAdd"}`),
		},
		{
			originalBytes: emptyBytesA,
			bytesToAdd:    []byte(`{"test":"bytesToAdd"}`),
			appendedBytes: []byte(`{"test":"bytesToAdd"}`),
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    []byte(`{"test":"bytesToAdd"}`),
			appendedBytes: []byte(`{"test":"originalBytes","test":"bytesToAdd"}`),
		},
		{
			originalBytes: []byte(`{"test":"originalBytes"}`),
			bytesToAdd:    []byte(`{"test":"bytesToAdd", "nested": {"moo":"moo"}}`),
			appendedBytes: []byte(`{"test":"originalBytes","test":"bytesToAdd", "nested": {"moo":"moo"}}`),
		},
	}

	for i := range tests {
		actual := AppendJSON(tests[i].originalBytes, tests[i].bytesToAdd)
		require.Equal(t, tests[i].appendedBytes, actual)
	}

}
