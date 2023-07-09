package gobag

// SliceExclusionStrings performs an exclusion operation on two string slices,
// source and reference. It returns two string slices: elementsOnlyInSource contains
// elements that are present in source but not in reference, and elementsOnlyInReference
// contains elements that are present in reference but not in source.
//
// Deprecated: Since Go 1.18, it is recommended to use SliceInclusion for checking element presence in slices.
func SliceExclusionStrings(source, reference []string) ([]string, []string) {
	presenceMap := make(map[string]uint8)
	for _, element := range source {
		presenceMap[element] |= (1 << 0)
	}
	for _, element := range reference {
		presenceMap[element] |= (1 << 1)
	}

	var elementsOnlyInSource, elementsOnlyInReference []string
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

// SliceExclusionInts performs an exclusion operation on two int slices, source
// and reference. It returns two int slices: elementsOnlyInSource contains
// elements that are present in source but not in reference, and
// elementsOnlyInReference contains elements that are present in reference but
// not in source.
//
// Deprecated: Since Go 1.18, it is recommended to use SliceInclusion for checking element presence in slices.
func SliceExclusionInts(source, reference []int) ([]int, []int) {
	presenceMap := make(map[int]uint8)
	for _, element := range source {
		presenceMap[element] |= (1 << 0)
	}
	for _, element := range reference {
		presenceMap[element] |= (1 << 1)
	}

	var elementsOnlyInSource, elementsOnlyInReference []int
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
