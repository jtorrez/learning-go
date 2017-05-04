package main

import (
  "fmt"
  "time"
)

// Prints s 5 times after sleeping
func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func main() {
  go say ("world")
  say("hello")
}
