package utils

func MapToArr[T E](elMap map[T]bool) []T {
	if len(elMap) == 0 {
		return []T{}
	}
	res := make([]T, 0, len(elMap))
	for el := range elMap {
		res = append(res, el)
	}
	return res
}

func ArrToMap[T E](elArr []T) map[T]int64 {
	if len(elArr) == 0 {
		return map[T]int64{}
	}
	res := make(map[T]int64, len(elArr))
	for _, el := range elArr {
		res[el]++
	}
	return res
}