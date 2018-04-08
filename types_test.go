package config

import (
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	_ flag.Value  = &arrayInt{}
	_ flag.Getter = &arrayInt{}
	_ flag.Value  = &arrayInt64{}
	_ flag.Getter = &arrayInt64{}
	_ flag.Value  = &arrayUint{}
	_ flag.Getter = &arrayUint{}
	_ flag.Value  = &arrayUint64{}
	_ flag.Getter = &arrayUint64{}
	_ flag.Value  = &arrayFloat64{}
	_ flag.Getter = &arrayFloat64{}
	_ flag.Value  = &arrayString{}
	_ flag.Getter = &arrayString{}
)

func Test_arrayInt(t *testing.T) {
	Convey("test arrayInt type", t, func() {
		ptr := new([]int)
		val := []int{-1, 2, -3, 4, -5, 6, -7, 8, -9, 0}
		arrInt := newArrayInt(val, ptr)

		Convey("get actual []int value from arrayInt", func() {
			So(arrInt.Get(), ShouldResemble, val)
		})

		Convey("parse new arrayInt value from a string", func() {
			out := []int{0, -9, 8, -7, 6, -5, 4, -3, 2, -1}
			arrInt.Set("0,-9,8,-7,6,-5,4,-3,2,-1")
			So([]int(*arrInt), ShouldResemble, out)
		})

		Convey("try to parse invalid arrayInt from a string", func() {
			So(arrInt.Set("false,42"), ShouldResemble, errCantUse("false,42", []int{}))
		})

		Convey("convert arrayInt to a string", func() {
			So(arrInt.String(), ShouldEqual, "-1,2,-3,4,-5,6,-7,8,-9,0")
		})
	})
}

func Test_arrayUint(t *testing.T) {
	Convey("test arrayUint type", t, func() {
		ptr := new([]uint)
		val := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		arrUint := newArrayUint(val, ptr)

		Convey("get actual []uint value from arrayUint", func() {
			So(arrUint.Get(), ShouldResemble, val)
		})

		Convey("parse new arrayUint value from a string", func() {
			out := []uint{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
			arrUint.Set("0,9,8,7,6,5,4,3,2,1")
			So([]uint(*arrUint), ShouldResemble, out)
		})

		Convey("try to parse invalid arrayUint from a string", func() {
			So(arrUint.Set("false,42"), ShouldResemble, errCantUse("false,42", []uint{}))
		})

		Convey("convert arrayUint to a string", func() {
			So(arrUint.String(), ShouldEqual, "1,2,3,4,5,6,7,8,9,0")
		})
	})
}

func Test_arrayInt64(t *testing.T) {
	Convey("test arrayInt64 type", t, func() {
		ptr := new([]int64)
		val := []int64{-1, 2, -3, 4, -5, 6, -7, 8, -9, 0}
		arrInt64 := newArrayInt64(val, ptr)

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

func Test_arrayUint64(t *testing.T) {
	Convey("test arrayUint64 type", t, func() {
		ptr := new([]uint64)
		val := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		arrUint64 := newArrayUint64(val, ptr)

		Convey("get actual []uint64 value from arrayUint64", func() {
			So(arrUint64.Get(), ShouldResemble, val)
		})

		Convey("parse new arrayUint64 value from a string", func() {
			out := []uint64{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
			arrUint64.Set("0,9,8,7,6,5,4,3,2,1")
			So([]uint64(*arrUint64), ShouldResemble, out)
		})

		Convey("try to parse invalid arrayUint64 from a string", func() {
			So(arrUint64.Set("false,42"), ShouldResemble, errCantUse("false,42", []uint64{}))
		})

		Convey("convert arrayUint64 to a string", func() {
			So(arrUint64.String(), ShouldEqual, "1,2,3,4,5,6,7,8,9,0")
		})
	})
}

func Test_arrayFloat64(t *testing.T) {
	Convey("test arrayFloat64 type", t, func() {
		ptr := new([]float64)
		val := []float64{-1.1, 2.2, -3.14, 4.56, -5.432, 6, -7, 8, -9, 0.01}
		arrFloat64 := newArrayFloat64(val, ptr)

		Convey("get actual []float64 value from arrayFloat64", func() {
			So(arrFloat64.Get(), ShouldResemble, val)
		})

		Convey("parse new arrayFloat64 value from a string", func() {
			out := []float64{0, -9, 8, -7, 6, -5, 4, -3, 2, -1}
			arrFloat64.Set("0,-9,8,-7,6,-5,4,-3,2,-1")
			So([]float64(*arrFloat64), ShouldResemble, out)
		})

		Convey("try to parse invalid arrayFloat64 from a string", func() {
			So(arrFloat64.Set("false,42"), ShouldResemble, errCantUse("false,42", []float64{}))
		})

		Convey("convert arrayFloat64 to a string", func() {
			So(arrFloat64.String(), ShouldEqual, "-1.1,2.2,-3.14,4.56,-5.432,6,-7,8,-9,0.01")
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
