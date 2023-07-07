package gobag

// SliceExclusionStrings ...
func SliceExclusionStrings(sliceA, sliceB []string) ([]string, []string) {
	m := make(map[string]uint8)
	for _, k := range sliceA {
		m[k] |= (1 << 0)
	}
	for _, k := range sliceB {
		m[k] |= (1 << 1)
	}

	var inAButNotB, inBButNotA []string
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		switch {
		case a && !b:
			inAButNotB = append(inAButNotB, k)
		case !a && b:
			inBButNotA = append(inBButNotA, k)
		}
	}
	return inAButNotB, inBButNotA
}

// SliceExclusionInts returns
func SliceExclusionInts(sliceA, sliceB []int) ([]int, []int) {
	m := make(map[int]uint8)
	for _, k := range sliceA {
		m[k] |= (1 << 0)
	}
	for _, k := range sliceB {
		m[k] |= (1 << 1)
	}

	var inAButNotB, inBButNotA []int
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		switch {
		case a && !b:
			inAButNotB = append(inAButNotB, k)
		case !a && b:
			inBButNotA = append(inBButNotA, k)
		}
	}
	return inAButNotB, inBButNotA
}
