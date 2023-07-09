package gobag

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

// SliceExclusion performs an exclusion operation on two slices, sourceSlice
// and reference. It returns two slices: elementsOnlyInSource contains elements
// that are present in sourceSlice but not in reference, and
// elementsOnlyInReference contains elements that are present in reference but
// not in sourceSlice. The elements must be of a comparable type, denoted by
// the type parameter T.
func SliceExclusion[T comparable](source, reference []T) ([]T, []T) {
	presenceMap := make(map[T]uint8)
	for _, k := range source {
		presenceMap[k] |= (1 << 0)
	}
	for _, k := range reference {
		presenceMap[k] |= (1 << 1)
	}

	var elementsOnlyInSource, elementsOnlyInReference []T
	for element, presence := range presenceMap {
		isPresentInSource := presence&(1<<0) != 0
		isPresentInReference := presence&(1<<1) != 0
		switch {
		case isPresentInSource && !isPresentInReference:
			elementsOnlyInSource = append(elementsOnlyInSource, element)
		case !isPresentInSource && isPresentInReference:
			elementsOnlyInReference = append(elementsOnlyInReference, element)
		}
	}
	return elementsOnlyInSource, elementsOnlyInReference
}

// SliceIntersection finds the intersection of two slices, source and target.
// It returns a new slice containing elements that are present in both source
// and target. The elements must be of a comparable type, denoted by the type
// parameter T.
func SliceIntersection[T comparable](source, target []T) []T {
	// Create a map to track the presence of elements in the two slices
	presenceMap := make(map[T]uint8)
	for _, element := range source {
		presenceMap[element] |= (1 << 0)
	}

	// Process sourceSlice and mark elements as present in sourceSlice
	for _, element := range target {
		presenceMap[element] |= (1 << 1)
	}

	var intersection []T
	for element, presence := range presenceMap {
		isPresentInSource := presence&(1<<0) != 0
		isPresentInTarget := presence&(1<<1) != 0
		switch {
		case isPresentInSource && isPresentInTarget:
			intersection = append(intersection, element)
		default:
			continue
		}
	}

	return intersection
}

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
