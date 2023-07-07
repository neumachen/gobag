package gobag

func RemoveEmptyString(slice *[]string) {
	i := 0
	p := *slice
	for _, entry := range p {
		if !StringIsEmpty(entry) {
			p[i] = entry
			i++
		}
	}
	*slice = p[0:i]
}
