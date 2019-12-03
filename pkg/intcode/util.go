package intcode

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "strconv"
  "strings"
)

func ReadProgramFromFile(path string) (*Program, error) {
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  content, err := ioutil.ReadAll(file)
  if err != nil {
    return nil, fmt.Errorf("read file: %w", err)
  }

  vals := strings.Split(string(content), ",")
  data := []int{}
  for _, v := range vals {
    d, err := strconv.Atoi(v)
    if err != nil {
      return nil, fmt.Errorf("parse int: %w", err)
    }

    data = append(data, d)
  }

  return &Program{
    Memory:  data,
    Counter: 0,
  }, nil
}
