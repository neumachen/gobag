package gobag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint8ToFloat64(t *testing.T) {
	expected := 1.234567890
	actual, err := Uint8ToFloat64([]byte(`1.234567890`))
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
