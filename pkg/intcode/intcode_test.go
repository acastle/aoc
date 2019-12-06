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
    { Input: []int{1002,4,3,4,33}, Expected: []int{1002,4,3,4,99} },
    { Input: []int{1101,100,-1,4,0}, Expected: []int{1101,100,-1,4,99} },
  }

  for _, d := range data {
    program := Program{ Memory: d.Input, Counter: 0 }
    err := program.Run()
    if err != nil {
      t.Fail()
    }

    if !reflect.DeepEqual(program.Memory, d.Expected) {
      t.Fail()
    }
  }
}

func TestProgram_Input(t *testing.T) {
  data := []struct {
    Input int
    Addr int
    Expected []int
  } {
    { Input: 10, Addr: 1, Expected: []int{0,10,0,0} },
    { Input: 198451, Addr: 3, Expected: []int{0,0,0,198451} },
  }

  for _, d := range data {
    program := Program{ Memory: []int {0, 0, 0, 0}, Counter: 0, In: func()int {return d.Input}}
    program.Input(d.Addr)
    if !reflect.DeepEqual(program.Memory, d.Expected) {
      t.Fail()
    }
  }
}

func TestProgram_Output(t *testing.T) {
  data := []struct {
    Input []int
    Expected int
  } {
    { Input: []int{104,10,99}, Expected: 10 },
    { Input: []int{4,2,99}, Expected: 99 },
  }

  for _, d := range data {
    var out int
    program := Program{ Memory: d.Input, Counter: 0, Out: func(a int) { out = a }}
    err := program.Run()
    if err != nil {
      t.Fail()
    }

    if out != d.Expected {
      t.Fail()
    }
  }
}

func TestProgram_Jump(t *testing.T) {
  data := []struct {
    Program []int
    Input int
    Output int
  } {
    { Program: []int{3,9,8,9,10,9,4,9,99,-1,8}, Input: 8, Output: 1 },
    { Program: []int{3,9,8,9,10,9,4,9,99,-1,8}, Input: 8, Output: 1 },
  }

  for _, d := range data {
    var out int
    program := Program {
      Memory: d.Program,
      Counter: 0,
      Out: func(a int) { out = a },
      In: func()int{return d.Input},
    }

    err := program.Run()
    if err != nil {
      t.Fail()
    }

    if out != d.Output {
      t.Fail()
    }
  }
}

