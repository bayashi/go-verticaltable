# go-verticaltable

<a href="https://github.com/bayashi/go-verticaltable/blob/main/LICENSE" title="go-verticaltable License"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png" alt="MIT License"></a>
<a href="https://github.com/bayashi/go-verticaltable/actions" title="go-verticaltable CI"><img src="https://github.com/bayashi/go-verticaltable/workflows/main/badge.svg" alt="go-verticaltable CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/go-verticaltable" title="go-verticaltable report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/go-verticaltable" alt="go-verticaltable report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/go-verticaltable" title="Go go-verticaltable package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/go-verticaltable.svg" alt="Go Reference: go-verticaltable"></a>

`go-verticaltable` provides vertical `key, value` table which is separated by horizontal header line.

## Usage

```go
package main

import (
    "os"

    "github.com/bayashi/go-verticaltable"
)

func main() {
    vt := verticaltable.NewTable(os.Stdout)

	vt.Header("foo")
	vt.Row("ID", "123")
	vt.Row("Select Type", "SIMPLE")
	vt.Row("Table", "foo")

	vt.Header("bar")
	vt.Row("ID", "456")
	vt.Row("Select Type", "UNIQUE")
	vt.Row("Table", "bar")

    tb.Render()
	// Output:
	// ********** 1. foo **********
	//          ID: 123
	// Select Type: SIMPLE
	//       Table: foo
	// ********** 2. bar **********
	//          ID: 456
	// Select Type: UNIQUE
	//       Table: bar
}
```

## Installation

```cmd
go get github.com/bayashi/go-verticaltable
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
