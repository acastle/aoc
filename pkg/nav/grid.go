package nav

type Point struct {
	X int
	Y int
}

type Path map[Point]int

func MinStepsToIntersection(a Path, b Path) int {
  points := IntersectPaths(a, b)
  lowest := SumOfStepsToPoint(a, b, points[0])
  for _, p := range points[1:] {
    sum := SumOfStepsToPoint(a, b, p)
    lowest = Min(sum, lowest)
  }

  return lowest
}

func Min(a int, b int) int {
  if a < b {
    return a
  }

  return b
}

func SumOfStepsToPoint(a Path, b Path, p Point) int {
  return a[p] + b[p]
}

func IntersectPaths(a Path, b Path) []Point {
  intersections := []Point{}
  for p := range a {
    if _, ok := b[p]; ok && !(p.X == 0 && p.Y == 0) {
      intersections = append(intersections, p)
    }
  }

  return intersections
}

func Abs(v int) int {
  if v < 0 {
    return v * -1
  }

  return v
}

func ManhattanDistance(p Point) int{
  return Abs(p.X) + Abs(p.Y)
}
