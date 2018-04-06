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

func Test_arrayInt64(t *testing.T) {
	Convey("test arrayInt64 type", t, func() {
		ptr := new([]int64)
		val := []int64{-1, 2, -3, 4, -5, 6, -7, 8, -9, 0}
		arrInt64 := newArrayInt64(val, ptr)

		Convey("instantiate arrayInt64 with constructor func", func() {
			So(arrInt64, ShouldHaveSameTypeAs, new(arrayInt64))
			So(arrInt64, ShouldImplement, (*flag.Value)(nil))
			So(arrInt64, ShouldImplement, (*flag.Getter)(nil))
		})

		Convey("get actual []int64 value from arrayInt64", func() {
			So(arrInt64.Get(), ShouldResemble, val)
		})

		Convey("parse new arrayInt64 value from a string", func() {
			out := []int64{0, -9, 8, -7, 6, -5, 4, -3, 2, -1}
			arrInt64.Set("0,-9,8,-7,6,-5,4,-3,2,-1")
			So([]int64(*arrInt64), ShouldResemble, out)
		})

		Convey("try to parse invalid arrayInt64 from a string", func() {
			So(arrInt64.Set("false,42"), ShouldResemble, errCantUse("false,42", []int64{}))
		})

		Convey("convert arrayInt64 to a string", func() {
			So(arrInt64.String(), ShouldEqual, "-1,2,-3,4,-5,6,-7,8,-9,0")
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
