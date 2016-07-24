package memberset

import (
  "testing"
)

func TestSetAddDelete(t *testing.T) {

  m := New()

  m.Add(1234)
  if ok := m.Get(1234); !ok {
    t.Errorf("Unable to find 1234 in output")
  }

  m.Delete(1234)
  if ok := m.Get(1234); ok {
    t.Errorf("1234 was not deleted.")
  }

}

func TestDifferentType(t *testing.T) {
  m := New()

  m.Add(9999)
  if ok := m.Get(9999); !ok {
    t.Error("Did not store integer 9999.")
  }

  if ok := m.Get("9999"); ok {
    t.Error("Returned OK on integer value, when asking for a string value.")
  }

}

func TestStringAddDelete(t *testing.T) {

  testSet := New()

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

  testSet := New()

  testSet.Add(1234)
  if ok := testSet.Get(1234); !ok {
    t.Errorf("Unable to find ID 1234 in output")
  }

  testSet.Set(4567)
  if ok := testSet.Get(4567); !ok {
    t.Errorf("Unable to find ID 4567 in output")
  }

}

func TestInt64Set(t *testing.T) {

  testSet := New()

  testSet.Add(int64(4444))
  if ok := testSet.Get(int64(4444)); !ok {
    t.Error("Unable to add int64(4444) to set")
  }

}

func TestMultiAdd(t *testing.T) {
  testSet := New()
  testSet.Add("1234")
  testSet.Add("1234")
  testSet.Add("1234")
  testSet.Add("1234")
  if ok := testSet.Get("1234"); !ok {
    t.Error("Unable to get string1234 from set")
  }
}

/* === Benchmarks === */

// Aliased to Set()
func BenchmarkAdd(b *testing.B) {

  testSet := New()

  for i := 0; i < b.N; i++ {
    testSet.Add(i)
  }

}

// Non-aliased
func BenchmarkSet(b *testing.B) {

  testSet := New()

  for i := 0; i < b.N; i++ {
    testSet.Set(i)
  }

}

func BenchmarkGet(b *testing.B) {

  testSet := New()

  for i := 0; i < b.N; i++ {
    testSet.Get(i)
  }

}

func BenchmarkDelete(b *testing.B) {

  testSet := New()

  for i := 0; i < b.N; i++ {
    testSet.Delete(i)
  }

}

func BenchmarkAddDelete(b *testing.B) {

  testSet := New()

  for i := 0; i < b.N; i++ {
    testSet.Add(i)
    testSet.Delete(i)
  }

}

func BenchmarkAddGet(b *testing.B) {

  testSet := New()

  for i := 0; i < b.N; i++ {
    testSet.Add(i)
    testSet.Get(i)
  }

}

func BenchmarkAddGetDelete(b *testing.B) {

  testSet := New()

  for i := 0; i < b.N; i++ {
    testSet.Add(i)
    testSet.Get(i)
    testSet.Delete(i)
  }

}
