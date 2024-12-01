package utils

import (
  "fmt"
  "os"
  "path"
)

func GetInput(day string) string {
  wd, err := os.Getwd()
  if err != nil {
    fmt.Printf("WORKDIR_Error: %v\n", err)
  }

  p := path.Join(wd, day, "input.txt") 
  b, err := os.ReadFile(p)
  if err != nil {
    fmt.Printf("WORKDIR_Error: %v\n", err)
  }

  return string(b)
}

