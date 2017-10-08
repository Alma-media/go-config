# go-config
Utility package to read the configuration.

## Installation
```sh
go get -u github.com/Alma-media/go-config
```

## How to use?
1. Create a config struct
2. Set default values (if needed) for struct fields
3. Import package `github.com/Alma-media/go-config`
4. Pass the struct (or better a pointer to it if you are going to pass the config across the application) to `Init()` func
5. Check error and enjoy

## Priorities
1. flags - hi
2. env vars - mid
3. defaults - low

## Examples
```go
package main

import (
	"fmt"

	"github.com/Alma-media/go-config"
)

type Config struct {
	Boolean bool `default:"true"`
	Nested  struct {
		Integer int     `default:"42"`
		Float64 float64 `default:"3.14"`
		String  string  `default:"text"`
	}
}

func main() {
	// create an instance
	conf := &Config{}
	// pass to Init() func and check the error
	if err := config.Init(conf, "MYAPP"); err != nil {
		fmt.Println(err)
	}
	// use your config
	fmt.Println(conf)
}
```
