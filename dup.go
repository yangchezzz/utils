package utils

// RemoveDupEl remove duplicate element, but will disrupt order
func RemoveDupEl[T int8 | int16 | int | int32 | int64 | float32 | float64 | string | bool | byte](els []T) []T {
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

func RemoveDupElKeepOrder[T int8 | int16 | int | int32 | int64 | float32 | float64 | string | bool | byte](els []T) []T {
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
