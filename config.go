// Package config provides flexible access to config variables by priority:
// flags - HI,
// environment variables - MID,
// default values defined with a struct field tags - LOW
package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	// receiver is not a pointer to a struct or a nil pointer
	errInvalidReceiver = errors.New("The argument to Init() func must be a non-nil pointer to a struct")
	// value is not addressable and or obtained by the use of unexported struct fields
	errCantSet = errors.New("Value can not be set")
	// trying to assign the value to unsupported field type
	errUnsupportedType = func(typeName string) error {
		return fmt.Errorf("Unsupported type [%s] in config", typeName)
	}
	// value parse error
	errCantUse = func(val string, typ interface{}) error {
		return fmt.Errorf("cannot use [%s] as type [%T]", val, typ)
	}
	// missing required argument/flag
	errMissingRequired = func(name string) error {
		return fmt.Errorf("missing required [--%s] argument/flag", name)
	}
)

// private constants
const (
	// keyDefaultTag - tag name for default value
	keyDefaultTag = "default"
	// keyIsRequired shold have any non-empty value if variable is required
	keyIsRequired = "required"
	// keyEnvVar - tag name for env variable name
	keyEnvTag = "env"
	// keyFlagTag - tag name for variable flag.
	// By default (if there is no tag "flag" for struct field) will have name:
	// -structname-nestedstructname-varname
	keyFlagTag = "keyFlag"
	// constant for internal use
	emptyPrefix = ""
	// comma separator
	comma = ","
	// space symbol
	space = " "
)

// EnvPrefix is a prefix in the beginning of environment variable name (used to
// easily differentiate variables of your application).
var EnvPrefix string

// command line arguments
var args = os.Args[1:]

// required args/flags container
var seen map[string]bool

// Init config values.
func Init(c interface{}, prefix string) error {
	// check argument type (only pointer to struct is supported)
	rv := reflect.ValueOf(c)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errInvalidReceiver
	}
	// set custom "application" prefix (will be used to build ENV VAR names)
	EnvPrefix = prefix
	// init flagset
	flagSet := NewFlagSet("config", flag.ContinueOnError)
	// reset required list
	seen = make(map[string]bool)
	if err := initConfig(rv, flagSet, emptyPrefix); err != nil {
		return err
	}
	// parse flags
	if err := flagSet.Parse(args); err != nil {
		return err
	}
	// mark as seen flags that have been set
	flagSet.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	// find missing required values
	for flagName, ok := range seen {
		if !ok {
			return errMissingRequired(flagName)
		}
	}
	// success
	return nil
}

// initConfig recursively loads parameters to Config struct, supports nested
// anonymous structs.
func initConfig(c reflect.Value, flagSet *FlagSet, prefix string) error {
	c = reflect.Indirect(c)
	if c.Kind() != reflect.Struct {
		return errInvalidReceiver
	}
	for i := 0; i < c.NumField(); i++ {
		var value string
		field := c.Field(i)
		if field.Kind() == reflect.Struct {
			np := nestedPrefix(prefix, c.Type().Field(i).Name)
			err := initConfig(field.Addr(), flagSet, np)
			if err != nil {
				return err
			}
			continue
		}
		if !field.CanSet() {
			return errCantSet
		}
		structField := c.Type().Field(i)
		flgKey := flagName(structField, prefix)
		// read "is required" field tag
		if isRequired := structField.Tag.Get(keyIsRequired); isRequired != "" {
			// init map cell with flgKey (set false because it was not seen yet)
			seen[flgKey] = false
		}
		// getting value from "default" tag
		defValue := structField.Tag.Get(keyDefaultTag)
		if defValue != "" {
			value = defValue
			seen[flgKey] = true
		}
		// retrieve value from ENV variable
		envValue := os.Getenv(envName(structField, prefix))
		if envValue != "" {
			value = envValue
			seen[flgKey] = true
		}
		// set value with a flag
		err := setValue(field, flagSet, flgKey, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// setValue casts string value and assigns it to the field of Config struct.
func setValue(field reflect.Value, flagSet *FlagSet, flgKey, value string) error {
	switch t := field.Interface().(type) {
	case time.Duration:
		val, err := time.ParseDuration(value)
		if err != nil {
			return errCantUse(value, t)
		}
		flagSet.DurationVar(field.Addr().Interface().(*time.Duration), flgKey, val, "")
	case int:
		val, err := strconv.Atoi(value)
		if err != nil && value != "" {
			return errCantUse(value, t)
		}
		flagSet.IntVar(field.Addr().Interface().(*int), flgKey, val, "")
	case int64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil && value != "" {
			return errCantUse(value, t)
		}
		flagSet.Int64Var(field.Addr().Interface().(*int64), flgKey, val, "")
	case uint:
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil && value != "" {
			return errCantUse(value, t)
		}
		flagSet.UintVar(field.Addr().Interface().(*uint), flgKey, uint(val), "")
	case uint64:
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil && value != "" {
			return errCantUse(value, t)
		}
		flagSet.Uint64Var(field.Addr().Interface().(*uint64), flgKey, val, "")
	case float64:
		val, err := strconv.ParseFloat(value, 64)
		if err != nil && value != "" {
			return errCantUse(value, t)
		}
		flagSet.Float64Var(field.Addr().Interface().(*float64), flgKey, val, "")
	case string:
		flagSet.StringVar(field.Addr().Interface().(*string), flgKey, value, "")
	case []string:
		flagSet.ArrayStringVar(field.Addr().Interface().(*[]string), flgKey, strings.Split(value, ","), "")
	case bool:
		val, err := strconv.ParseBool(value)
		if err != nil {
			val = false
		}
		flagSet.BoolVar(field.Addr().Interface().(*bool), flgKey, val, "")
	default:
		return errUnsupportedType(field.Kind().String())
	}
	return nil
}

// nestedPrefix concats prefix for generating default flag names and env variable names.
func nestedPrefix(base, newPrefix string) string {
	if base == "" {
		return newPrefix
	}
	return base + " " + newPrefix
}

// envName gets environment variable name for passed field based on provided
// struct tags or default rules (ENVPREFIX_STRUCTNAME_NESTEDSTRUCTNAME_VARNAME).
func envName(field reflect.StructField, prefix string) string {
	tag := field.Tag.Get(keyEnvTag)
	if tag != "" {
		return tag
	}
	s := joinStrings("_", EnvPrefix, prefix, field.Name)
	return strings.ToUpper(s)
}

// flagName gets flag name for passed field based on provided struct tags or
// default rules (-structname-nestedstructname-varname).
func flagName(field reflect.StructField, prefix string) string {
	tag := field.Tag.Get(keyFlagTag)
	if tag != "" {
		return tag
	}
	s := joinStrings("-", prefix, field.Name)
	return strings.ToLower(s)
}

// joinStrings similar to strings.Join, but omits empty values, also replaces
// spaces with provided separator.
func joinStrings(sep string, parts ...string) string {
	var components []string
	for _, part := range parts {
		if part != "" {
			components = append(components, part)
		}
	}
	joined := strings.Join(components, sep)
	return strings.Replace(joined, " ", sep, -1)
}
