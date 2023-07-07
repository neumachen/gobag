package gobag

import "strings"

// GenRedisKey generates a key with a prefix
func GenRedisKey(keys ...string) string {
	return strings.Join(keys, ":")
}
