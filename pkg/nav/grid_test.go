package nav

import (
  "testing"
)

func TestMinStepsToIntersection(t *testing.T) {
  paths := make([]Path, 4)
  paths[0], _ = ParsePath("R75,D30,R83,U83,L12,D49,R71,U7,L72")
  paths[1], _ = ParsePath("U62,R66,U55,R34,D71,R55,D58,R83")

  paths[2], _ = ParsePath("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
  paths[3], _ = ParsePath("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")


  data := []struct{
    A Path
    B Path
    Expect int
  } {
    { A: paths[0], B: paths[1], Expect: 610 },
    { A: paths[2], B: paths[3], Expect: 410 },
  }

  for _, d := range data {
    if MinStepsToIntersection(d.A, d.B) != d.Expect {
      t.Fail()
    }
  }
}