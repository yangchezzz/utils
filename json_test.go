package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDump(t *testing.T) {
	Convey("TestDump", t, func() {
		Convey("normal", func() {
			testCase := []TestData[byte]{
				{
					Arr: []byte("hello world"),
					Map: map[byte]bool{},
				},
			}
			for _, c := range testCase {
				res := Dump(c.Arr)
				t.Logf("%v", res)
				t.Logf("%v", Dump(c.Map))
			}
		})
	})
}
