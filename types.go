package config

import (
	"strings"
)

// stringArray Value
type stringArray []string

func newStringArray(val []string, p *[]string) *stringArray {
	*p = val
	return (*stringArray)(p)
}

func (s *stringArray) Set(val string) error {
	*s = stringArray(strings.Split(val, ","))
	return nil
}

func (s *stringArray) Get() interface{} {
	return []string(*s)
}

func (s *stringArray) String() string {
	return strings.Join([]string(*s), ",")
}
