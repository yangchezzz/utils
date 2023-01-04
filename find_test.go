package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountElIndexInArr(t *testing.T) {
	Convey("TestCountElIndexInArr", t, func() {
		Convey("normal", func() {
			testCase := []TestData[int64]{
				{
					Arr: []int64{1, 2, 2, 1, 1, 1, 1},
				},
			}
			for _, c := range testCase {
				res := CountElIndexInArr(c.Arr, 1)
				t.Logf("%v", res)
				So(res, ShouldEqual, int64(5))
			}
		})
	})
}

func TestFindFirstElIndexInArr(t *testing.T) {
	Convey("TestFindFirstElIndexInArr", t, func() {
		Convey("normal", func() {
			testCase := []TestData[int64]{
				{
					Arr: []int64{1, 2, 2, 10, 1, 1, 1, 1},
				},
			}
			for _, c := range testCase {
				res := FindFirstElIndexInArr(c.Arr, 10)
				t.Logf("%v", res)
				So(res, ShouldEqual, int64(3))
			}
		})
	})
}

func TestFindLastElIndexInArr(t *testing.T) {
	Convey("TestFindLastElIndexInArr", t, func() {
		Convey("normal", func() {
			testCase := []TestData[int64]{
				{
					Arr: []int64{1, 2, 2, 10, 1, 1, 1, 1},
				},
			}
			for _, c := range testCase {
				res := FindLastElIndexInArr(c.Arr, 10)
				t.Logf("%v", res)
				So(res, ShouldEqual, int64(3))
			}
		})
	})
}
