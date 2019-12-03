package fuel

import "testing"

func TestModule_Fuel(t *testing.T) {
  data := []struct {
    Mass float64
    Expect float64
  } {
    {12, 2 },
    {14, 2 },
    {1969, 654 },
    {100756, 33583 },
  }

  for _, d := range data {
    mod := Module{ Mass: d.Mass }
    if mod.Fuel() != d.Expect {
      t.Fail()
    }
  }
}

func TestModule_TotalFuel(t *testing.T) {
  data := []struct {
    Mass float64
    Expect float64
  } {
    {14, 2 },
    {1969, 966 },
    {100756, 50346 },
  }

  for _, d := range data {
    mod := Module{ Mass: d.Mass }
    total := mod.TotalFuel()
    if total != d.Expect {
      t.Fail()
    }
  }
}


