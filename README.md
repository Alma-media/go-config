# go-config

[![Build Status][circleci-badge]][circleci-link]
[![Report Card][report-badge]][report-link]
[![GoCover][cover-badge]][cover-link]

Utility package to read the configuration.

## Installation
```sh
go get -u github.com/tiny-go/config
```

## How to use?
1. Import package `github.com/tiny-go/config`
2. Create a config struct
3. Set default values (if needed) for struct fields
4. Pass the pointer to a struct to `Init()` func
5. Check error and enjoy

## Supported data types
- `bool`
- `int`, `[]int`
- `uint`, `[]uint`
- `int64`, `[]int64`
- `uint64`, `[]uint64`
- `float64`, `[]float64`
- `time.Duration`, `[]time.Duration`
- `string`, `[]string`

## Priorities
1. flags - hi
2. env vars - mid
3. defaults - low

## Examples
```go
package main

import (
	"fmt"

	"github.com/tiny-go/config"
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

[circleci-badge]: https://circleci.com/gh/tiny-go/config.svg?style=shield
[circleci-link]: https://circleci.com/gh/tiny-go/config
[report-badge]: https://goreportcard.com/badge/github.com/tiny-go/config
[report-link]: https://goreportcard.com/report/github.com/tiny-go/config
[cover-badge]: https://gocover.io/_badge/github.com/tiny-go/config
[cover-link]: https://gocover.io/github.com/tiny-go/config
