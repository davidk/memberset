# memberset

[![Build Status](https://travis-ci.org/davidk/memberset.svg?branch=master)](https://travis-ci.org/davidk/memberset)
[![GoDoc](https://godoc.org/github.com/davidk/memberset?status.svg)](https://godoc.org/github.com/davidk/memberset)

A simple set used for testing memberships in Go.

# Usage

### Get started by vendoring or by using `go get github.com/davidk/memberset`

    import "github.com/davidk/memberset"

    m := memberset.New()

    // Other types are also acceptable, like ints
    m.Add("123")

    // Is string "123" a member?
    if ok := m.Get("123"); ok {
      // Yes
    } else {
      // No
    }

    // To remove "123" from the set
    m.Delete("123")

# Related

There are other projects that have more through set implementations. Check these out:

* [deckarep/golang-set](https://github.com/deckarep/golang-set)

* [faith/set](https://github.com/fatih/set)

And [a proposal](https://github.com/golang/go/issues/16466) on golang/go to add a built-in set type.

# License

[CC0 1.0 Universal](https://creativecommons.org/publicdomain/zero/1.0/)
