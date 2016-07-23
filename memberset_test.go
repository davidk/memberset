package memberset

import (
  "testing"
)

func TestSetAddDelete(t *testing.T) {

  testSet := NewMemberSet()

  testSet.Add(1234)
  if ok := testSet.Get(1234); !ok {
    t.Errorf("Unable to find 1234 in output")
  }

  testSet.Delete(1234)
  if ok := testSet.Get(1234); ok {
    t.Errorf("1234 was not deleted.")
  }

}

func TestStringAddDelete(t *testing.T) {

  testSet := NewMemberSet()

  testSet.Add("cake")
  if ok := testSet.Get("cake"); !ok {
    t.Errorf("Unable to find cake in output")
  }

  testSet.Delete("cake")
  if ok := testSet.Get("cake"); ok {
    t.Errorf("cake was not deleted.")
  }

}

func TestIntSet(t *testing.T) {
  testSet := NewMemberSet()

  testSet.Add(1234)
  if ok := testSet.Get(1234); !ok {
    t.Errorf("Unable to find ID 1234 in output")
  }

  testSet.Set(4567)
  if ok := testSet.Get(4567); !ok {
    t.Errorf("Unable to find ID 4567 in output")
  }

}

/* === Benchmarks === */

// Aliased to Set()
func BenchmarkAdd(b *testing.B) {

  testSet := NewMemberSet()

  for i := 0; i < b.N; i++ {
    testSet.Add(i)
  }

}

// Non-aliased
func BenchmarkSet(b *testing.B) {

  testSet := NewMemberSet()

  for i := 0; i < b.N; i++ {
    testSet.Set(i)
  }
}

func BenchmarkGet(b *testing.B) {

  testSet := NewMemberSet()

  for i := 0; i < b.N; i++ {
    testSet.Get(i)
  }

}

func BenchmarkDelete(b *testing.B) {

  testSet := NewMemberSet()

  for i := 0; i < b.N; i++ {
    testSet.Delete(i)
  }

}
