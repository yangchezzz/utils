package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDupEl(t *testing.T) {
	Convey("TestRemoveDupEl", t, func() {
		Convey("int", func() {
			testCase := []TestData[int]{
				{
					Arr: []int{4, 4, 1, 3, 2, 3, 2, 99, 100, 99},
				},
			}

			for _, c := range testCase {
				t.Logf("%v", Dump(RemoveDupEl(c.Arr)))
				t.Logf("%v", Dump(RemoveDupElKeepOrder(c.Arr)))
			}
		})
		Convey("string", func() {
			testCase := []TestData[string]{
				{
					Arr: []string{"4", "4", "1", "3", "2", "3", "2", "99", "100", "99"},
				},
			}

			for _, c := range testCase {
				t.Logf("%v", Dump(RemoveDupEl(c.Arr)))
				t.Logf("%v", Dump(RemoveDupElKeepOrder(c.Arr)))
			}
		})
	})
}
