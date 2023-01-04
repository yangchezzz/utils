package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMapToArr(t *testing.T) {
	Convey("TestMapToArr", t, func() {
		Convey("normal", func() {
			testCase := []TestData[string]{
				{
					Map: map[string]bool{
						"1": true,
						"2": true,
						"3": true,
						"4": true,
						"5": true,
					},
				},
			}

			for _, c := range testCase {
				res := MapToArr(c.Map)
				t.Logf("%v", Dump(res))
			}
		})
	})
}

func TestArrToMap(t *testing.T) {
	Convey("TestArrToMap", t, func() {
		Convey("normal", func() {
			testCase := []TestData[string]{
				{
					Arr: []string{"1", "2", "1", "3", "4"},
				},
			}

			for _, c := range testCase {
				res := ArrToMap(c.Arr)
				t.Logf("%v", Dump(res))
			}
		})
	})
}
