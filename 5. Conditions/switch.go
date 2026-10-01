package main

import (
	"fmt"
	"runtime"
)

func main() {
  fmt.Println("Try clicking the Run button.")
  switch os :=runtime.GOOS; os {
    case "darwin":
    fmt.Println("OS X")
    case "linux":
    fmt.Println("linux")
    default:
    fmt.Println("%s.\n", os)
  }
}