package main

import (
  "fmt"
)

// Iterative Fibonacci up to 'long'
func Fib(n int) int {
  if n <= 1 {
    return n
  } else {

    // some utility variables
    tmp, cur, next := 0, 0, 1

    for i := 0; i <= n; i++ {
      tmp = cur + next

      // Look for integer overflow
      // panic() if it's found
      if next < cur {
        fmt.Println("int overflow at n=", i)
        panic("integer overflow detected in Fib()")
      }

      cur = next
      next = tmp
    }

    return cur
  }
}

// Try to Fibonacci up to 100
// You can't, but try anyway
func main() {
  for i := 1; i <= 100; i++ {
    fmt.Println(i, " | ", Fib(i))
  }
}
