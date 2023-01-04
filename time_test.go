package utils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertTimeToStr(t *testing.T) {
	Convey("TestConvertTimeToStr", t, func() {
		Convey("normal", func() {
			testCase := []TestData[time.Time]{
				{
					El: time.Now(),
				},
			}
			for _, c := range testCase {
				res := ConvertTimeToStr(c.El, "2006-01-02 15:04:05")
				t.Logf("%v", res)
			}
		})
	})
}

func TestConvertTimeToStrDefault(t *testing.T) {
	Convey("TestConvertTimeToStrDefault", t, func() {
		Convey("normal", func() {
			res := ConvertTimeToStrDefault()
			t.Logf("%v", res)
		})
	})
}
