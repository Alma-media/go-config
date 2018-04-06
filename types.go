package config

import (
	"fmt"
	"strconv"
	"strings"
)

// arrayInt implements flag.Value, flag.Getter interfaces.
type arrayInt []int

func newArrayInt(val []int, p *[]int) *arrayInt {
	*p = val
	return (*arrayInt)(p)
}

func (i *arrayInt) Set(val string) error {
	*i = []int{}
	arrStr := strings.Split(val, comma)
	for _, currStr := range arrStr {
		currInt, err := strconv.Atoi(currStr)
		if err != nil {
			return errCantUse(val, []int{})
		}
		*i = append(*i, currInt)
	}
	return nil
}

func (i *arrayInt) Get() interface{} {
	return []int(*i)
}

func (i *arrayInt) String() string {
	var arrStr []string
	for _, curr := range *i {
		arrStr = append(arrStr, fmt.Sprint(curr))
	}
	return strings.Join(arrStr, comma)
}

// arrayUint Value implements flag.Value, flag.Getter interfaces.
type arrayUint []uint

func newArrayUint(val []uint, p *[]uint) *arrayUint {
	*p = val
	return (*arrayUint)(p)
}

func (u *arrayUint) Set(val string) error {
	*u = []uint{}
	arrStr := strings.Split(val, comma)
	for _, currStr := range arrStr {
		currUint64, err := strconv.ParseUint(currStr, 10, 32)
		if err != nil {
			return errCantUse(val, []uint{})
		}
		*u = append(*u, uint(currUint64))
	}
	return nil
}

func (u *arrayUint) Get() interface{} {
	return []uint(*u)
}

func (u *arrayUint) String() string {
	var arrStr []string
	for _, curr := range *u {
		arrStr = append(arrStr, fmt.Sprint(curr))
	}
	return strings.Join(arrStr, comma)
}

// arrayInt64 implements flag.Value, flag.Getter interfaces.
type arrayInt64 []int64

func newArrayInt64(val []int64, p *[]int64) *arrayInt64 {
	*p = val
	return (*arrayInt64)(p)
}

func (i64 *arrayInt64) Set(val string) error {
	*i64 = []int64{}
	arrStr := strings.Split(val, comma)
	for _, currStr := range arrStr {
		currInt, err := strconv.ParseInt(currStr, 10, 64)
		if err != nil {
			return errCantUse(val, []int64{})
		}
		*i64 = append(*i64, currInt)
	}
	return nil
}

func (i64 *arrayInt64) Get() interface{} {
	return []int64(*i64)
}

func (i64 *arrayInt64) String() string {
	var arrStr []string
	for _, curr := range *i64 {
		arrStr = append(arrStr, fmt.Sprint(curr))
	}
	return strings.Join(arrStr, comma)
}

// arrayString implements flag.Value, flag.Getter interfaces.
type arrayString []string

func newArrayString(val []string, p *[]string) *arrayString {
	*p = val
	return (*arrayString)(p)
}

func (s *arrayString) Set(val string) error {
	*s = arrayString(strings.Split(val, comma))
	return nil
}

func (s *arrayString) Get() interface{} {
	return []string(*s)
}

func (s *arrayString) String() string {
	return strings.Join([]string(*s), comma)
}
