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

// SliceContains checks if a given target element is present in the source
// slice. It iterates through the source slice and compares each element with
// the target element. The function returns true if the target element is found
// in the source slice, and false otherwise. The elements must be of a
// comparable type, denoted by the type parameter T.
func SliceContains[T comparable](source []T, target T) bool {
	for i := range source {
		if source[i] == target {
			return true
		}
	}
	return false
}
