package gobag

import "strconv"

// Uint8ToFloat64 ...
func Uint8ToFloat64(bytes []byte) (float64, error) {
	f, parseErr := strconv.ParseFloat(string(bytes), 64)
	if parseErr != nil {
		return 0, nil
	}
	return f, nil
}
