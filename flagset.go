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

// ArrayIntVar defines an []int flag with specified name, default value, and usage string.
// The argument p points to an []int variable in which to store the value of the flag.
func (f *FlagSet) ArrayIntVar(p *[]int, name string, value []int, usage string) {
	f.Var(newArrayInt(value, p), name, usage)
}

// ArrayIntVar defines an []int flag with specified name, default value, and usage string.
// The argument p points to an []int variable in which to store the value of the flag.
func ArrayIntVar(p *[]int, name string, value []int, usage string) {
	flag.CommandLine.Var(newArrayInt(value, p), name, usage)
}

// ArrayInt64Var defines an []int64 flag with specified name, default value, and usage string.
// The argument p points to an []int64 variable in which to store the value of the flag.
func (f *FlagSet) ArrayInt64Var(p *[]int64, name string, value []int64, usage string) {
	f.Var(newArrayInt64(value, p), name, usage)
}

// ArrayInt64Var defines an []int64 flag with specified name, default value, and usage string.
// The argument p points to an []int64 variable in which to store the value of the flag.
func ArrayInt64Var(p *[]int64, name string, value []int64, usage string) {
	flag.CommandLine.Var(newArrayInt64(value, p), name, usage)
}

// ArrayStringVar defines an []string flag with specified name, default value, and usage string.
// The argument p points to an []string variable in which to store the value of the flag.
func (f *FlagSet) ArrayStringVar(p *[]string, name string, value []string, usage string) {
	f.Var(newArrayString(value, p), name, usage)
}

// ArrayStringVar defines a []string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
func ArrayStringVar(p *[]string, name string, value []string, usage string) {
	flag.CommandLine.Var(newArrayString(value, p), name, usage)
}
