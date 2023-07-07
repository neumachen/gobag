package gobag

import (
	"regexp"
	"strings"
)

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

// SnakeCase transforms the s into snake cased format
func SnakeCase(s string) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	key := ""
	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' && key[len(key)-1] != '_' {
			key += "_" + string(v)
		} else if v == ' ' {
			key += "_"
		} else if v == '/' || v == '\\' {
			key += "_or_"
		} else {
			key = key + string(v)
		}
	}
	return strings.ToLower(key)
}
