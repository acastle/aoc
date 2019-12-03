package intcode

import (
  "reflect"
  "testing"
)

func TestProgram_Run(t *testing.T) {
  data := []struct {
    Input []int
    Expected []int
  } {
    { Input: []int{1,0,0,0,99}, Expected: []int{2,0,0,0,99} },
    { Input: []int{2,3,0,3,99}, Expected: []int{2,3,0,6,99} },
    { Input: []int{2,4,4,5,99,0}, Expected: []int{2,4,4,5,99,9801} },
    { Input: []int{1,1,1,4,99,5,6,0,99}, Expected: []int{30,1,1,4,2,5,6,0,99} },
  }

  for _, d := range data {
    program := Program{ Memory: d.Input, Counter: 0 }
    err := program.Run()
    if err != nil {
      t.Fail()
    }

    if reflect.DeepEqual(program.Memory, d.Expected) {
      t.Fail()
    }
  }
}
