package main

import "testing"

func TestMul(t *testing.T) {
  m := mul(10, 10)
  if m != 100 {
    t.Errorf("result was incorrect, got: %d, want: %d.", m, 100)
  }
}

func TestMinus(t *testing.T) {
  m := minus(1, 10)
  if m != -9 {
    t.Errorf("result was incorrect, got: %d, want: %d.", m, -9)
  }
}

func TestAdd(t *testing.T) {
  s := add(1, 10)
  if s != 11 {
    t.Errorf("result was incorrect, got: %d, want: %d.", s, 11)
  }
}
