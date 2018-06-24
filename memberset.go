/*
Package memberset is a simple set used for testing memberships in Go.

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
*/
package memberset

import (
  "sync"
)

// MemberSet wraps up a map[interface{}]struct{} for ease of use.
type MemberSet struct {
  sync.RWMutex
  m map[interface{}]struct{}
}

// New initializes a new MemberSet
func New() *MemberSet {
  return &MemberSet{ m: make(map[interface{}]struct{}) }
}

// Get looks for an item in the set. It doesn't really retrieve
// anything, just checks for existence.
func (s *MemberSet) Get (value interface{}) (bool) {
  defer s.RUnlock()
  s.RLock()

  if _, ok := s.m[value]; ok {
    return true
  }

  return false
}

// Add is an alias to Set()
func (s *MemberSet) Add (value interface{}) {
  s.Set(value)
}

// Set adds a value to the set
func (s *MemberSet) Set (value interface{}) {
  s.Lock()
  s.m[value] = struct{}{}
  s.Unlock()
}

// Delete removes the value from the set
func (s *MemberSet) Delete (value interface{}) {
  s.Lock()
  delete(s.m, value)
  s.Unlock()
}
