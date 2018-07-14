# ut

[![GoDoc](https://godoc.org/github.com/iawia002/ut?status.svg)](https://godoc.org/github.com/iawia002/ut)

The missing utility library for Go.

## Installation

```
go get github.com/iawia002/ut
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/iawia002/ut"
)

func main() {
	fmt.Println(ut.Reverse("54321")) // 12345
}
```

## Examples

```go
gore> :import github.com/iawia002/ut
gore> ut.Reverse("12345")
"54321"
gore> ut.Range(0, 5)
[]int{0, 1, 2, 3, 4, 5}
gore> ut.TrimSpace("1 2\n3\t45")
"12345"
gore> ut.Domain("hello.com")
"hello"
gore> ut.ItemInSlice("a", []string{"a", "b", "c"})
true
gore> ut.Md5("123")
"202cb962ac59075b964b07152d234b70"
```