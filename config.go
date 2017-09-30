// Package config provides flexible access to config variables by priority:
// flags - HI,
// environment variables - MID,
// default values defined with a struct field tags - LOW
package config

import (
	"errors"
	"flag"
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
		return errors.New("Unsupported type in config: " + typeName)
	}
)

// private constants
const (
	// keyDefaultTag - tag name for default value
	keyDefaultTag = "default"
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

// Init config values.
func Init(c interface{}, prefix string) error {
	rv := reflect.ValueOf(c)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errInvalidReceiver
	}
	EnvPrefix = prefix
	flagSet := flag.NewFlagSet("config", flag.ContinueOnError)
	if err := initConfig(rv, flagSet, emptyPrefix); err != nil {
		return err
	}
	flagSet.Parse(args)
	return nil
}

// initConfig recursively loads parameters to Config struct, supports nested
// anonymous structs.
func initConfig(c reflect.Value, flagSet *flag.FlagSet, prefix string) error {
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
		defValue := structField.Tag.Get(keyDefaultTag)
		if defValue != "" {
			value = defValue
		}
		envValue := os.Getenv(envName(structField, prefix))
		if envValue != "" {
			value = envValue
		}
		flgKey := flagName(structField, prefix)
		err := setValue(field, flagSet, flgKey, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// setValue casts string value and assigns it to the field of Config struct.
func setValue(field reflect.Value, flagSet *flag.FlagSet, flgKey, value string) error {
	switch field.Interface().(type) {
	case time.Duration:
		val, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		flagSet.DurationVar(field.Addr().Interface().(*time.Duration), flgKey, val, "")
	case int:
		val, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		flagSet.IntVar(field.Addr().Interface().(*int), flgKey, val, "")
	case int64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		flagSet.Int64Var(field.Addr().Interface().(*int64), flgKey, val, "")
	case uint:
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		flagSet.UintVar(field.Addr().Interface().(*uint), flgKey, uint(val), "")
	case uint64:
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		flagSet.Uint64Var(field.Addr().Interface().(*uint64), flgKey, val, "")
	case float64:
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		flagSet.Float64Var(field.Addr().Interface().(*float64), flgKey, val, "")
	case string:
		flagSet.StringVar(field.Addr().Interface().(*string), flgKey, value, "")
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
