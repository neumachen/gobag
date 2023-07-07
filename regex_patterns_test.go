package gobag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTimeStampISO8601Regex(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		testData := []string{
			"2021-02-28T23:28:47Z",
			"2021-02-28T23:28:47+0000",
			"2017-06-01T18:43:26",
		}

		for _, data := range testData {
			result := TimeStampISO8601Regex.Match([]byte(data))
			require.True(t, result, data)
		}
	})

	t.Run("invalid data", func(t *testing.T) {
		testData := []string{
			"Sun, 28 Feb 2021 23:28:47 +0000",
			"2019-10-12T14:20:50.9999+0700",
			"2019-10-12T14:20:50.9999+07:00",
		}

		for _, data := range testData {
			result := TimeStampISO8601Regex.Match([]byte(data))
			require.False(t, result, data)
		}

	})
}

func TestTimeStampRFC3339Regex(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		testData := []string{
			"2021-02-28T23:28:47Z",
			"2021-02-28T23:28:47+0000",
			"2019-10-12T14:20:50+07:00",
			"2019-10-12T14:20:50.9999+0700",
			"2019-10-12T14:20:50.9999+07:00",
			"2019-10-12T14:20:50.9999Z",
		}

		for _, data := range testData {
			result := TimeStampRFC3999Regex.Match([]byte(data))
			require.True(t, result, data)
		}
	})

	t.Run("invalid data", func(t *testing.T) {
		testData := []string{
			"2017-06-01T18:43:26",
			"Sun, 28 Feb 2021 23:28:47 +0000",
		}

		for _, data := range testData {
			result := TimeStampRFC3999Regex.Match([]byte(data))
			require.False(t, result, data)
		}

	})
}
