package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// arrayInt implements flag.Value, flag.Getter interfaces.
type arrayInt []int

func newArrayInt(val []int, p *[]int) *arrayInt {
	*p = val
	return (*arrayInt)(p)
}

func (i *arrayInt) Set(val string) error {
	*i = []int{}
	if val != "" {
		arrStr := strings.Split(val, comma)
		for _, currStr := range arrStr {
			currInt, err := strconv.Atoi(currStr)
			if err != nil {
				return errCantUse(val, []int{})
			}
			*i = append(*i, currInt)
		}
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

// arrayUint implements flag.Value, flag.Getter interfaces.
type arrayUint []uint

func newArrayUint(val []uint, p *[]uint) *arrayUint {
	*p = val
	return (*arrayUint)(p)
}

func (u *arrayUint) Set(val string) error {
	*u = []uint{}
	if val != "" {
		arrStr := strings.Split(val, comma)
		for _, currStr := range arrStr {
			currUint64, err := strconv.ParseUint(currStr, 10, 32)
			if err != nil {
				return errCantUse(val, []uint{})
			}
			*u = append(*u, uint(currUint64))
		}
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
	if val != "" {
		arrStr := strings.Split(val, comma)
		for _, currStr := range arrStr {
			currInt, err := strconv.ParseInt(currStr, 10, 64)
			if err != nil {
				return errCantUse(val, []int64{})
			}
			*i64 = append(*i64, currInt)
		}
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

// arrayUint64 implements flag.Value, flag.Getter interfaces.
type arrayUint64 []uint64

func newArrayUint64(val []uint64, p *[]uint64) *arrayUint64 {
	*p = val
	return (*arrayUint64)(p)
}

func (u64 *arrayUint64) Set(val string) error {
	*u64 = []uint64{}
	if val != "" {
		arrStr := strings.Split(val, comma)
		for _, currStr := range arrStr {
			curr, err := strconv.ParseUint(currStr, 10, 64)
			if err != nil {
				return errCantUse(val, []uint64{})
			}
			*u64 = append(*u64, curr)
		}
	}
	return nil
}

func (u64 *arrayUint64) Get() interface{} {
	return []uint64(*u64)
}

func (u64 *arrayUint64) String() string {
	var arrStr []string
	for _, curr := range *u64 {
		arrStr = append(arrStr, fmt.Sprint(curr))
	}
	return strings.Join(arrStr, comma)
}

// arrayFloat64 implements flag.Value, flag.Getter interfaces.
type arrayFloat64 []float64

func newArrayFloat64(val []float64, p *[]float64) *arrayFloat64 {
	*p = val
	return (*arrayFloat64)(p)
}

func (f64 *arrayFloat64) Set(val string) error {
	*f64 = []float64{}
	if val != "" {
		arrStr := strings.Split(val, comma)
		for _, currStr := range arrStr {
			curr, err := strconv.ParseFloat(currStr, 64)
			if err != nil {
				return errCantUse(val, []float64{})
			}
			*f64 = append(*f64, curr)
		}
	}
	return nil
}

func (f64 *arrayFloat64) Get() interface{} {
	return []float64(*f64)
}

func (f64 *arrayFloat64) String() string {
	var arrStr []string
	for _, curr := range *f64 {
		arrStr = append(arrStr, fmt.Sprint(curr))
	}
	return strings.Join(arrStr, comma)
}

// arrayDuration implements flag.Value, flag.Getter interfaces.
type arrayDuration []time.Duration

func newArrayDuration(val []time.Duration, p *[]time.Duration) *arrayDuration {
	*p = val
	return (*arrayDuration)(p)
}

func (d *arrayDuration) Set(val string) error {
	*d = []time.Duration{}
	if val != "" {
		arrStr := strings.Split(val, comma)
		for _, currStr := range arrStr {
			curr, err := time.ParseDuration(currStr)
			if err != nil {
				return errCantUse(val, []time.Duration{})
			}
			*d = append(*d, curr)
		}
	}
	return nil
}

func (d *arrayDuration) Get() interface{} {
	return []time.Duration(*d)
}

func (d *arrayDuration) String() string {
	var arrStr []string
	for _, curr := range *d {
		arrStr = append(arrStr, curr.String())
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
	*s = []string{}
	if val != "" {
		*s = arrayString(strings.Split(val, comma))
	}
	return nil
}

func (s *arrayString) Get() interface{} {
	return []string(*s)
}

func (s *arrayString) String() string {
	return strings.Join([]string(*s), comma)
}
