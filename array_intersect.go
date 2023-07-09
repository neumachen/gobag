package gobag

// SliceIntersectStrings ...
func SliceIntersectStrings(sliceA, sliceB []string) []string {
	m := make(map[string]uint8)
	for _, k := range sliceA {
		m[k] |= (1 << 0)
	}
	for _, k := range sliceB {
		m[k] |= (1 << 1)
	}

	var inAAndB []string
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		switch {
		case a && b:
			inAAndB = append(inAAndB, k)
		default:
			continue
		}
	}
	return inAAndB
}

// SliceIntersectInts ...
func SliceIntersectInts(sliceA, sliceB []int) []int {
	m := make(map[int]uint8)
	for _, k := range sliceA {
		m[k] |= (1 << 0)
	}
	for _, k := range sliceB {
		m[k] |= (1 << 1)
	}

	var inAAndB []int
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		switch {
		case a && b:
			inAAndB = append(inAAndB, k)
		default:
			continue
		}
	}
	return inAAndB
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
