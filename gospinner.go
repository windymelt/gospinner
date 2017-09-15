package main

import (
  "fmt"
  "time"
)

func main () {
     chs := [4]rune{ '|', '/', '-', '\\' }
     i := 0
     for {
       fmt.Printf("\r%c", chs[i])
       i = (i + 1) % 4
       time.Sleep(100 * time.Millisecond)
     }
}
