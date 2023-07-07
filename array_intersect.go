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
