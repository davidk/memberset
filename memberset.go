/*

  == gopher-membership ==

  A very simple library for membership testing.
  Add a key and get it back quickly.

*/

package memberset

import (
  "sync"
)

type MemberSet struct {
  sync.RWMutex
  m map[interface{}]struct{}
}

func NewMemberSet() *MemberSet {
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
