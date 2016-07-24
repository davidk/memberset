/*

  == gopher-membership ==

  A very simple library for membership testing.

  Get started by vendoring or using go get github.com/davidk/memberset

  m := memberset.NewMemberSet()
  m.Add("123")

  # Is string "123" a member?
  if ok := m.Get("123"); ok {
    // Yes
  }

*/

package memberset

import (
  "sync"
)

type MemberSet struct {
  sync.RWMutex
  m map[interface{}]struct{}
}

func New() *MemberSet {
  return &MemberSet{ m: make(map[interface{}]struct{}) }
}

// Get() shouldn't necessarily care what type is there,
// just as long as it matches what was given in Set()
func (s *MemberSet) Get (value interface{}) (bool) {
  defer s.RUnlock()
  s.RLock()

  if _, ok := s.m[value]; ok {
    return true
  }

  return false
}


func (s *MemberSet) Add (value interface{}) {
  s.Set(value)
}

func (s *MemberSet) Set (value interface{}) {
  s.Lock()
  s.m[value] = struct{}{}
  s.Unlock()
}

func (s *MemberSet) Delete (value interface{}) {
  s.Lock()
  delete(s.m, value)
  s.Unlock()
}
