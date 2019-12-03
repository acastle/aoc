package nav

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

const (
  RIGHT = 'R'
  LEFT  = 'L'
  UP    = 'U'
  DOWN  = 'D'
)

func ParsePathsFromFile(path string) ([]Path, error) {
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  paths := []Path{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    path, err := ParsePath(scanner.Text())
    if err != nil {
      return nil, fmt.Errorf("parse path: %w", err)
    }

    paths = append(paths, path)
  }

  return paths, nil
}

func ParsePath(path string) (Path, error) {
  instructions := strings.Split(path, ",")
  paths := map[Point]int{}

  count := 0
  last := Point{0, 0}
  for _, ins := range instructions {
    dir := ins[0]
    val, err := strconv.Atoi(ins[1:])
    if err != nil {
      return nil, fmt.Errorf("parse int: %w", err)
    }

    dx, dy := 0, 0
    switch dir {
    case RIGHT:
      dx = 1
    case LEFT:
      dx = -1
    case UP:
      dy = 1
    case DOWN:
      dy = -1
    }

    for i:=0; i<val; i++ {
      count++
      next := last
      next.X += (dx * 1)
      next.Y += (dy * 1)
      if _, ok := paths[next]; !ok {
        paths[next] = count
      }

      last = next
    }
  }

  return paths, nil
}
