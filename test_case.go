package utils

import "time"

type TestData[T int | int64 | string | byte | time.Time] struct {
	Map map[T]bool
	Arr []T
	El  T
}
