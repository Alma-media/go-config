// Package config provides flexible access to config variables by priority:
// flags - HI,
// environment variables - MID,
// default values defined with a struct field tags - LOW
package config

import (
	"flag"
	"os"
	"reflect"
	"strconv"
	"strings"
	"errors"
)


var (
	// receiver is not a pointer to a struct or a nil pointer
	errInvalidReceiver = errors.New("The argument to Init() func must be a non-nil pointer to a struct")
	// value is not addressable and or obtained by the use of unexported struct fields
	errCantSet = errors.New("Value can not be set")
	// trying to assign the value to unsupported field type
	errUnsupportedType = func(typeName string) error {
		return errors.New("Unsupported type: " + typeName)
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

var EnvPrefix string

// command line arguments
var args = os.Args[1:]

// New is config constructor. Creates an instance of config and fills it with
// provided values by priority.
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

// initConfig recursively loads parameters to Config struct, supports nested structs.
func initConfig(c reflect.Value, flagSet *flag.FlagSet, prefix string) error {
        c = reflect.Indirect(c)
        if c.Kind() != reflect.Struct {
                return errInvalidReceiver
        }
	for i := 0; i < c.NumField(); i++ {
		var value string
		field := c.Field(i)
		if field.Kind() == reflect.Struct {
			np := nestedPrefix(prefix, field.Type().Name())
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
// Currently supports: bool, int and string values.
func setValue(field reflect.Value, flagSet *flag.FlagSet, flgKey, value string) error {
	switch field.Kind() {
	case reflect.Int:
		val, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		ptr := field.Addr().Interface().(*int)
		flagSet.IntVar(ptr, flgKey, val, "")
	case reflect.String:
		ptr := field.Addr().Interface().(*string)
		flagSet.StringVar(ptr, flgKey, value, "")
	case reflect.Bool:
		val, err := strconv.ParseBool(value)
		if err != nil {
			val = false
		}
		ptr := field.Addr().Interface().(*bool)
		flagSet.BoolVar(ptr, flgKey, val, "")
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
// struct tags or default rules (ENVPREFIX_STRUCTNAME_NESTEDSTRUCTNAME_VARNAME)
func envName(field reflect.StructField, prefix string) string {
	tag := field.Tag.Get(keyEnvTag)
	if tag != "" {
		return tag
	}
	s := joinStrings("_", EnvPrefix, prefix, field.Name)
	return strings.ToUpper(s)
}

// flagName gets flag name for passed field based on provided struct tags or
// default rules (-structname-nestedstructname-varname)
func flagName(field reflect.StructField, prefix string) string {
	tag := field.Tag.Get(keyFlagTag)
	if tag != "" {
		return tag
	}
	s := joinStrings("-", prefix, field.Name)
	return strings.ToLower(s)
}

// joinStrings similar to strings.Join, but omits empty values, also replaces
// spaces with provided separator
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
