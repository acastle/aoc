package main

import (
  "reflect"
  "testing"
)

func TestSplitCharacters(t *testing.T) {
  data := []struct {
    Input int
    Expected []int
  } {
    { Input: 112233, Expected:[]int{1,1,2,2,3,3} },
    { Input: 9856245, Expected:[]int{9,8,5,6,2,4,5} },
  }

  for _, d := range data {
    result := SplitCharacters(d.Input)
    if !reflect.DeepEqual(d.Expected, result) {
      t.Fail()
    }
  }
}

func TestDoubles(t *testing.T) {
  data := []struct {
    Input []int
    Expected bool
  } {
    { Input: []int{1,1,2,2,3,3} , Expected: true },
    { Input: []int{9,8,5,6,2,4,5} , Expected: false },
  }

  for _, d := range data {
    if Doubles(d.Input) != d.Expected {
      t.Fail()
    }
  }
}

func TestDoublesButNotTriples(t *testing.T) {
  data := []struct {
    Input []int
    Expected bool
  } {
    { Input: []int{1,1,2,2,3,3} , Expected: true },
    { Input: []int{9,8,5,6,2,4,5} , Expected: false },
    { Input: []int{9,9,9,9,2,4,5} , Expected: false },
    { Input: []int{1,1,1,9,2,4,5} , Expected: false },
    { Input: []int{1,1,5,6,7,8,9} , Expected: true },
  }

  for _, d := range data {
    if ContainsStandaloneDouble(d.Input) != d.Expected {
      t.Fail()
    }
  }
}

func TestIncreasing(t *testing.T) {
  data := []struct {
    Input []int
    Expected bool
  } {
    { Input: []int{1,1,1,1,1,1} , Expected: true },
    { Input: []int{1,2,3,4,5,6} , Expected: true },
    { Input: []int{9,8,5,6,2,4,5} , Expected: false },
  }

  for _, d := range data {
    if Increasing(d.Input) != d.Expected {
      t.Fail()
    }
  }
}

func TestValidPass(t *testing.T) {
  data := []struct {
    Input int
    Expected bool
  } {
    { Input: 111111 , Expected: true },
    { Input: 223450 , Expected: false },
    { Input: 123789 , Expected: false },
  }

  for _, d := range data {
    if ValidPass(d.Input) != d.Expected {
      t.Fail()
    }
  }
}

func TestValidPass2(t *testing.T) {
  data := []struct {
    Input int
    Expected bool
  } {
    { Input: 112233 , Expected: true },
    { Input: 123444 , Expected: false },
    { Input: 111122 , Expected: true },
  }

  for _, d := range data {
    if ValidPass2(d.Input) != d.Expected {
      t.Fail()
    }
  }
}