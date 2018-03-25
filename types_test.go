package config

import (
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_arrayInt(t *testing.T) {
	Convey("test arrayInt type", t, func() {
		ptr := new([]int)
		val := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		arrInt := newArrayInt(val, ptr)

		Convey("instantiate arrayInt with constructor func", func() {
			So(arrInt, ShouldHaveSameTypeAs, new(arrayInt))
			So(arrInt, ShouldImplement, (*flag.Value)(nil))
			So(arrInt, ShouldImplement, (*flag.Getter)(nil))
		})

		Convey("get actual []int value from arrayInt", func() {
			So(arrInt.Get(), ShouldResemble, val)
		})

		Convey("parse new arrayInt value from a string", func() {
			out := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
			arrInt.Set("0,9,8,7,6,5,4,3,2,1")
			So([]int(*arrInt), ShouldResemble, out)
		})

		Convey("try to parse invalid arrayInt from a string", func() {
			So(arrInt.Set("false,42"), ShouldResemble, errCantUse("false,42", []int{}))
		})

		Convey("convert arrayInt to a string", func() {
			So(arrInt.String(), ShouldEqual, "1,2,3,4,5,6,7,8,9,0")
		})
	})
}

func Test_arrayString(t *testing.T) {
	Convey("test arrayString type", t, func() {
		ptr := new([]string)
		val := []string{"foo", "bar"}
		arrStr := newArrayString(val, ptr)

		Convey("instantiate arrayString with constructor func", func() {
			So(arrStr, ShouldHaveSameTypeAs, new(arrayString))
			So(arrStr, ShouldImplement, (*flag.Value)(nil))
			So(arrStr, ShouldImplement, (*flag.Getter)(nil))
		})

		Convey("get actual []string value from arrayString", func() {
			So(arrStr.Get(), ShouldResemble, val)
		})

		Convey("parse new arrayString value from a string", func() {
			out := []string{"this", "is", "new", "value"}
			arrStr.Set("this,is,new,value")
			So([]string(*arrStr), ShouldResemble, out)
		})

		Convey("convert arrayString to string", func() {
			So(arrStr.String(), ShouldEqual, "foo,bar")
		})
	})
}
