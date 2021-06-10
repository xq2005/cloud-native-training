package main

import "testing"

func TestMul(t *testing.T) {
  m := mul(10, 10)
  if m != 100 {
    t.Errorf("Mul was incorrect, got: %d, want: %d.", m, 100)
  }
}
