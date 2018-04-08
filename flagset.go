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

// ArrayUintVar defines an []uint flag with specified name, default value, and usage string.
// The argument p points to an []uint variable in which to store the value of the flag.
func (f *FlagSet) ArrayUintVar(p *[]uint, name string, value []uint, usage string) {
	f.Var(newArrayUint(value, p), name, usage)
}

// ArrayInt64Var defines an []int64 flag with specified name, default value, and usage string.
// The argument p points to an []int64 variable in which to store the value of the flag.
func (f *FlagSet) ArrayInt64Var(p *[]int64, name string, value []int64, usage string) {
	f.Var(newArrayInt64(value, p), name, usage)
}

// ArrayUint64Var defines an []uint64 flag with specified name, default value, and usage string.
// The argument p points to an []uint64 variable in which to store the value of the flag.
func (f *FlagSet) ArrayUint64Var(p *[]uint64, name string, value []uint64, usage string) {
	f.Var(newArrayUint64(value, p), name, usage)
}

// ArrayFloat64Var defines an []float64 flag with specified name, default value, and usage string.
// The argument p points to an []float64 variable in which to store the value of the flag.
func (f *FlagSet) ArrayFloat64Var(p *[]float64, name string, value []float64, usage string) {
	f.Var(newArrayFloat64(value, p), name, usage)
}

// ArrayStringVar defines an []string flag with specified name, default value, and usage string.
// The argument p points to an []string variable in which to store the value of the flag.
func (f *FlagSet) ArrayStringVar(p *[]string, name string, value []string, usage string) {
	f.Var(newArrayString(value, p), name, usage)
}
