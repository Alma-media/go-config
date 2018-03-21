package config

import (
	"strings"
)

// ArrValueDelimiter is an array value delimiter.
const ArrValueDelimiter = ","

// arrayString Value implements flag.Value, flag.Getter interfaces.
type arrayString []string

func newArrayString(val []string, p *[]string) *arrayString {
	*p = val
	return (*arrayString)(p)
}

func (s *arrayString) Set(val string) error {
	*s = arrayString(strings.Split(val, ArrValueDelimiter))
	return nil
}

func (s *arrayString) Get() interface{} {
	return []string(*s)
}

func (s *arrayString) String() string {
	return strings.Join([]string(*s), ArrValueDelimiter)
}
