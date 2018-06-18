# memberset

[![Build Status](https://travis-ci.org/davidk/memberset.svg?branch=master)](https://travis-ci.org/davidk/memberset)
[![GoDoc](https://godoc.org/github.com/davidk/memberset?status.svg)](https://godoc.org/github.com/davidk/memberset)

A simple set used for testing memberships in Go.

# Usage

    import "github.com/davidk/memberset"

    m := memberset.New()

    // Adding a value into the set is just a m.Add() call
    // Other types are also acceptable, like ints, floats, etc (as long as it works with an interface{})
    m.Add("123")
    
    // If it helps your mnemonic memory, m.Add() is just an alias to m.Set()
    m.Set(123) // <--- This is the same as m.Add(123)

    // Eventually, we'll need to check to see if what we added is present in the set
    if ok := m.Get("123"); ok {
      // Yes, do something
    } else {
      // No, do something else
    }

    // And maybe we'll need to remove "123" from the set
    m.Delete("123")

# Installation

### Get started by vendoring or by using `go get -u github.com/davidk/memberset`

# Related

There are other projects that have more through set implementations. Check these out:

* [deckarep/golang-set](https://github.com/deckarep/golang-set)

* [faith/set](https://github.com/fatih/set)

And [a proposal](https://github.com/golang/go/issues/16466) on golang/go to add a built-in set type.

# License

[CC0 1.0 Universal](https://creativecommons.org/publicdomain/zero/1.0/)
