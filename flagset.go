package config

import "flag"

// FlagSet represents extended flag.FlagSet.
type FlagSet struct {
	*flag.FlagSet
}

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewFlagSet(name string, errorHandling flag.ErrorHandling) *FlagSet {
	return &FlagSet{FlagSet: flag.NewFlagSet(name, errorHandling)}
}

// ArrayStringVar defines a []string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
func (f *FlagSet) ArrayStringVar(p *[]string, name string, value []string, usage string) {
	f.Var(newStringArray(value, p), name, usage)
}
