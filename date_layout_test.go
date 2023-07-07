package gobag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateLayout(t *testing.T) {
	assert.Equal(t, "2006-01-02", DateLayout(YYYYMMDD))
}
