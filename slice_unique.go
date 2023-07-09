package gobag

// FilterUniqueElements returns a new slice containing unique elements from the
// given slice. It removes any duplicate elements and preserves the order of
// the original elements. The elements must be of a comparable type, denoted by
// the type parameter T.
//
// The function iterates through the input slice and keeps track of seen
// elements using a map. Duplicate elements are skipped, and only the first
// occurrence of each unique element is included in the result. The returned
// slice contains the unique elements in the same order as they appear in the
// original slice.
func FilterUniqueElements[T comparable](slice []T) []T {
	seenElements := make(map[T]struct{}, len(slice))
	uniqueSlice := make([]T, 0, len(slice))
	for _, element := range slice {
		if _, ok := seenElements[element]; ok {
			continue
		}
		seenElements[element] = struct{}{}
		uniqueSlice = append(uniqueSlice, element)
	}
	return uniqueSlice
}
