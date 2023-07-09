package gobag

// ArrayContainsStr checks if a given target string is present in the source
// string slice. It iterates through the source slice and compares each element
// with the target string. The function returns true if the target string is
// found in the source slice, and false otherwise.
//
// Deprecated: Since Go 1.18, it is recommended to use SliceContains for
// checking element presence in slices.
func ArrayContainsStr(source []string, target string) bool {
	for _, value := range source {
		if value == target {
			return true
		}
	}
	return false
}
