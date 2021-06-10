package main

import "fmt"

func mul(a, b int) int {
  return a * b
}

func main() {
  fmt.Println("mul: ", mul(3, 5))
}
