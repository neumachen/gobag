package typeconv

func AsInt64Ptr(i int64) *int64 {
	return &i
}

func Float64ToInt(f float64) int {
	return int(f)
}
