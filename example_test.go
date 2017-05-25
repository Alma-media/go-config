package config_test

import (
	"fmt"

	"github.com/Alma-media/go-config"
)

// TestConfig struct
type TestConfig struct {
	String string `default:"strvalue"`
	Int    int    `default:"42"`
	Bool   bool   `default:"false"`
	Nested Nested
}

// Nested config struct
type Nested struct {
	String string `default:"nestedstrvalue"`
}

func ExampleInit_Failure_NilPointer() {
	var nilPointerConf *TestConfig
	fmt.Println(config.Init(nilPointerConf, "myapp"), nilPointerConf)
	// Output: The argument to Init() func must be a non-nil pointer to a struct <nil>
}

func ExampleInit_Failure_NotAPointer() {
	var notPointerConf TestConfig
	fmt.Println(config.Init(notPointerConf, "myapp"), notPointerConf)
	// Output: The argument to Init() func must be a non-nil pointer to a struct { 0 false {}}
}

func ExampleInit_Success_CreateAPointer() {
	validPointerConfig := &TestConfig{}
	fmt.Println(config.Init(validPointerConfig, "myapp"), validPointerConfig)
	// Output: <nil> &{strvalue 42 false {nestedstrvalue}}
}

func ExampleInit_Success_PassAPointer() {
	var validConfig TestConfig
	fmt.Println(config.Init(&validConfig, "myapp"), validConfig)
	// Output: <nil> {strvalue 42 false {nestedstrvalue}}
}
