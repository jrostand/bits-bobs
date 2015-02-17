package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    if i % 15 == 0 {
      fmt.Println("FizBuz")
    } else if i % 3 == 0 {
      fmt.Println("Fiz")
    } else if i % 5 == 0 {
      fmt.Println("Buz")
    } else {
      fmt.Println(i)
    }
  }
}
