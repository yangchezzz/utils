package utils

func FindFirstElIndexInArr[T E](els []T, el T) int64 {
	if len(els) == 0 {
		return -1
	}
	
	for i, e := range els {
		if e == el {
			return int64(i)
		}
	}
	return -1
}

func FindLastElIndexInArr[T E](els []T, el T) int64 {
	if len(els) == 0 {
		return -1
	}
	
	for i := len(els) - 1; i > 0; i-- {
		e := els[i]
		if e == el {
			return int64(i)
		}
	}
	return -1
}

func CountElIndexInArr[T E](els []T, el T) int64 {
	if len(els) == 0 {
		return 0
	}
	
	res := int64(0)
	for _, e := range els {
		if e == el {
			res++
		}
	}
	return res
}