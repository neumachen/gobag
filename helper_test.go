package gobag

// we don't care about order. We only want to check if all the
// elements are in the other array
func testEqSliceStrs(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	expectedA := 0
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				expectedA++
			}
		}
	}

	expectedB := 0
	for i := range b {
		for j := range a {
			if b[i] == a[j] {
				expectedB++
			}
		}
	}

	return len(a) == expectedA && len(b) == expectedB
}

// we don't care about order. We only want to check if all the
// elements are in the other array
func testEqSliceInts(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	expectedA := 0
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				expectedA++
			}
		}
	}

	expectedB := 0
	for i := range b {
		for j := range a {
			if b[i] == a[j] {
				expectedB++
			}
		}
	}

	return len(a) == expectedA && len(b) == expectedB
}
