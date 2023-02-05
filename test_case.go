package utils

import "time"

type TestData[T E | time.Time] struct {
	Map map[T]bool
	Arr []T
	El  T
}