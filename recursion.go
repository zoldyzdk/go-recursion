package main

import "fmt"

func main() {
  fmt.Println(recur(5))
}

func recur(value int) int {
  if value == 0 {
    return 1
  }
  return value * recur(value-1)
}
