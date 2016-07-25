/*
Package memberset is a simple set used for testing memberships in Go.

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
