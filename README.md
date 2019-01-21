# memberset

[![Go Report Card](https://goreportcard.com/badge/github.com/davidk/memberset)](https://goreportcard.com/report/github.com/davidk/memberset)
[![Build Status](https://travis-ci.org/davidk/memberset.svg?branch=master)](https://travis-ci.org/davidk/memberset)
[![GoDoc](https://godoc.org/github.com/davidk/memberset?status.svg)](https://godoc.org/github.com/davidk/memberset)

A simple, in-memory set used for storing and testing memberships in Go.

# Usage

Sample program:

```go
package main

import (
    "fmt"
    set "github.com/davidk/memberset"
)

func main() {

    m := set.New()
    
    m.Set(123)

    if ok := m.Get(123); ok {
        fmt.Println("Yes, 123 is a part of the set")
    } else {
        fmt.Println("Nope, didn't find 123 in the set")
    }

    // m.Add == m.Set -- this is mostly for mnemonic memory
    m.Add("1234")

    if ok := m.Get("1234"); ok {
        fmt.Println("Yes, string 1234 has been added to the set")
    }

    m.Delete("1234")

    if ok := m.Get("1234"); ok {
        fmt.Println("Nope, string 1234 is still set. This is a bug.")
    } else {
        fmt.Println("Looks like string 1234 was successfully deleted!")
    }

}
```

Expected output:

```bash
$ go build mtest.go
$ ./mtest
Yes, 123 is a part of the set
Yes, string 1234 has been added to the set
Looks like string 1234 was successfully deleted!
```

# Installation

### Get started by vendoring or by using `go get -u github.com/davidk/memberset`

# Related

There are other projects that have more through set implementations. Check these out:

* [deckarep/golang-set](https://github.com/deckarep/golang-set)

* [faith/set](https://github.com/fatih/set)

And [a proposal](https://github.com/golang/go/issues/16466) on golang/go to add a built-in set type. There's a link to a [playground entry](https://play.golang.org/p/VhCbdtJzwz) with a complete set implementation.

# License

[CC0 1.0 Universal](https://creativecommons.org/publicdomain/zero/1.0/)
