package config

import (
	"flag"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJoinStrings(t *testing.T) {
	type testCase struct {
		title string
		in    []string
		out   string
	}
	var cases = []testCase{
		{"not empty strings", []string{"-", "a", "b", "c"}, "a-b-c"},
		{"with empty string", []string{"_", "", "b", "c"}, "b_c"},
		{"all empty strings", []string{"#", "", "", ""}, ""},
	}
	Convey("Join Strings", t, func() {
		for _, c := range cases {
			Convey(c.title, func() {
				So(joinStrings(c.in[0], c.in[1:]...), ShouldEqual, c.out)
			})
		}
	})
}

func TestNestedPrefix(t *testing.T) {
	type testCase struct {
		title string
		in    []string
		out   string
	}
	var cases = []testCase{
		{"empty base string", []string{"", "abc"}, "abc"},
		{"not empty base string", []string{"abc", "def"}, "abc def"},
	}
	Convey("Nested Prefix", t, func() {
		for _, c := range cases {
			Convey(c.title, func() {
				So(nestedPrefix(c.in[0], c.in[1]), ShouldEqual, c.out)
			})
		}
	})
}

func TestFlagName(t *testing.T) {
	type In struct {
		field  reflect.StructField
		prefix string
	}
	type testCase struct {
		title string
		in    In
		out   string
	}
	var cases = []testCase{
		{
			"default flag name",
			In{
				reflect.StructField{
					Name: "Test",
					Tag:  "",
					Type: reflect.TypeOf(""),
				},
				"Db",
			},
			"db-test",
		},
		{
			"with provided flag",
			In{
				reflect.StructField{
					Name: "Test",
					Tag:  "default:\"10\" " + keyFlagTag + ":\"test-flag\"",
					Type: reflect.TypeOf(""),
				},
				"Db",
			},
			"test-flag",
		},
	}
	Convey("Flag Name", t, func() {
		for _, c := range cases {
			Convey(c.title, func() {
				So(flagName(c.in.field, c.in.prefix), ShouldEqual, c.out)
			})
		}
	})
}

func TestEnvName(t *testing.T) {
	EnvPrefix = "TEST"

	type In struct {
		field  reflect.StructField
		prefix string
	}
	type testCase struct {
		title string
		in    In
		out   string
	}
	var cases = []testCase{
		{
			"with default tags",
			In{
				reflect.StructField{
					Name: "Test",
					Tag:  "",
					Type: reflect.TypeOf(""),
				},
				"Db",
			},
			EnvPrefix + "_DB_TEST",
		},
		{
			"with provided tags",
			In{
				reflect.StructField{
					Name: "Test",
					Tag:  "default:\"10\" " + keyEnvTag + ":\"TEST_ENV\"",
					Type: reflect.TypeOf(""),
				},
				"Db",
			},
			"TEST_ENV",
		},
	}
	Convey("Environment values", t, func() {
		for _, c := range cases {
			Convey(c.title, func() {
				So(envName(c.in.field, c.in.prefix), ShouldEqual, c.out)
			})
		}
	})
}

func TestSetValue(t *testing.T) {
	type testStruct struct {
		I   int
		S   string
		B   bool
		F32 float32
	}
	type In struct {
		field         reflect.Value
		flgKey, value string
	}
	type testCase struct {
		title string
		in    In
		out   interface{}
	}
	var reflectStruct = reflect.Indirect(reflect.ValueOf(new(testStruct)))
	var cases = []testCase{
		{
			"int value",
			In{
				reflectStruct.FieldByName("I"),
				"flag-test",
				"123",
			},
			123,
		},
		{
			"bool value",
			In{
				reflectStruct.FieldByName("B"),
				"flag-test",
				"true",
			},
			true,
		},
		{
			"string value",
			In{
				reflectStruct.FieldByName("S"),
				"flag-test",
				"test string",
			},
			"test string",
		},
		{
			"unsupported float32 value",
			In{
				reflectStruct.FieldByName("F32"),
				"flag-test",
				"3.14159",
			},
			3.14159,
		},
		{
			"wrong int value",
			In{
				reflect.Indirect(
					reflect.ValueOf(new(testStruct)),
				).FieldByName("I"),
				"flag-test",
				"wrong",
			},
			0,
		},
		{
			"wrong bool value",
			In{
				reflect.Indirect(
					reflect.ValueOf(new(testStruct)),
				).FieldByName("B"),
				"flag-test",
				"wrong",
			},
			false,
		},
	}
	Convey("Setting values", t, func() {
		for _, c := range cases {
			Convey(c.title, func() {
				flagSet := flag.NewFlagSet("config", flag.ContinueOnError)
				err := setValue(c.in.field, flagSet, c.in.flgKey, c.in.value)
				switch c.in.field.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					So(c.in.field.Int(), ShouldEqual, c.out)
				case reflect.String:
					So(c.in.field.String(), ShouldEqual, c.out)
				case reflect.Bool:
					So(c.in.field.Bool(), ShouldEqual, c.out)
				default:
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}
