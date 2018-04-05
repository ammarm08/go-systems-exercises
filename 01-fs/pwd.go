package main

import (
  "fmt"
  "os"
  "path/filepath"
  "log"
)

func main() {
  pwd, err := os.Getwd()
  checkError(err)

  fmt.Println(pwd)

  if len(os.Args) == 1 {
    return
  }

  if os.Args[1] != "-P" {
    return
  }

  info, err := os.Lstat(pwd)
  checkError(err)

  if info.Mode() & os.ModeSymlink != 0 {
    real, err := filepath.EvalSymlinks(pwd)
    if err == nil {
      fmt.Println(real)
    }
  }
}

func checkError(err error) {
  if err != nil {
    log.Fatal(err.Error())
  }
}
