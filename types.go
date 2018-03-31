package config

import (
	"strconv"
	"strings"
)

// arrayInt Value implements flag.Value, flag.Getter interfaces.
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
		arrStr = append(arrStr, strconv.Itoa(curr))
	}
	return strings.Join(arrStr, comma)
}

// arrayString Value implements flag.Value, flag.Getter interfaces.
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
