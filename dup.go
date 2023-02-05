package utils

// RemoveDupEl remove duplicate element, but will disrupt order
func RemoveDupEl[T E](els []T) []T {
	set := make(map[T]bool)
	for _, e := range els {
		if set[e] {
			continue
		} else {
			set[e] = true
		}
	}
	return MapToArr(set)
}

func RemoveDupElKeepOrder[T E](els []T) []T {
	set := make(map[T]bool)
	res := make([]T, 0)
	for _, e := range els {
		if set[e] {
			continue
		} else {
			set[e] = true
			res = append(res, e)
		}
	}
	return res
}