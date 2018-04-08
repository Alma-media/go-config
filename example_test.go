package config_test

import (
	"fmt"

	"github.com/Alma-media/go-config"
)

// TestConfig struct
type TestConfig struct {
	String   string  `default:"strvalue"`
	Int      int     `default:"42"`
	Bool     bool    `default:"false"`
	ArrInt64 []int64 `default:"-1,2,-3,4,-5"`
	Nested   Nested
}

// Nested config struct
type Nested struct {
	String string `default:"nestedstrvalue"`
}

func ExampleInit() {
	var conf TestConfig
	fmt.Println(config.Init(&conf, "myapp"), conf)
	// Output: <nil> {strvalue 42 false [-1 2 -3 4 -5] {nestedstrvalue}}
}
