package intcode

import (
  "fmt"
)

const (
  ADD = 1
  MUL = 2
  INPUT = 3
  OUTPUT = 4
  JUMP = 5
  JUMP_F = 6
  LESS = 7
  EQUALS = 8
  HALT = 99
)

var ParamsForOpcode = map[int]int {
  ADD: 3,
  MUL: 3,
  INPUT: 1,
  OUTPUT: 1,
  JUMP: 2,
  JUMP_F: 2,
  LESS: 3,
  EQUALS: 3,
  HALT: 0,
}

type Program struct {
  Memory []int
  Counter int

  In func()int
  Out func(int)
}

func (p *Program) Run() error {
  for {
    halt, err := p.NextInstruction();
    if err != nil {
      return fmt.Errorf("run: %w", err)
    }
    if halt {
      return nil
    }
  }
}

func (p *Program) NextInstruction() (bool, error) {
  op := p.Memory[p.Counter]
  code := op % 100
  addressingModes := (op - code) / 100
  params := [3]int{}
  for i := 0; i<ParamsForOpcode[code]; i++ {
    mode := addressingModes % 10
    addressingModes /= 10
    val := p.Memory[p.Counter + i + 1]
    if mode == 1 {
      params[i] = val
    } else if mode == 0 {
      params[i] = p.Memory[val]
    }
  }

  switch code {
  case ADD:
    p.Add(params[0], params[1], p.Memory[p.Counter + 3])
  case MUL:
    p.Multiply(params[0], params[1], p.Memory[p.Counter + 3])
  case HALT:
    return true, nil
  case INPUT:
    p.Input(p.Memory[p.Counter + 1])
  case OUTPUT:
    p.Output(params[0])
  case JUMP:
    p.JumpIfTrue(params[0], params[1])
  case JUMP_F:
    p.JumpIfFalse(params[0], params[1])
  case LESS:
    p.LessThan(params[0], params[1], p.Memory[p.Counter + 3])
  case EQUALS:
    p.Equals(params[0], params[1], p.Memory[p.Counter + 3])
  default:
    return false, fmt.Errorf("unknown opcode: %d", op)
  }

  return false, nil
}

func (p *Program) Add(a int, b int, reg int) {
  p.Memory[reg] = a+b
  p.Counter += ParamsForOpcode[ADD] + 1
}

func (p *Program) Multiply(a int, b int, reg int) {
  p.Memory[reg] = a*b
  p.Counter += ParamsForOpcode[MUL] + 1
}

func (p *Program) Input(reg int) {
  p.Memory[reg] = p.In()
  p.Counter += ParamsForOpcode[INPUT] + 1
}

func (p *Program) Output(a int) {
  p.Out(a)
  p.Counter += ParamsForOpcode[OUTPUT] + 1
}

func (p *Program) JumpIfTrue(a int, b int) {
  if a != 0 {
    p.Counter = b
  } else {
    p.Counter += ParamsForOpcode[JUMP] + 1
  }
}

func (p *Program) JumpIfFalse(a int, b int) {
  if a == 0 {
    p.Counter = b
  } else {
    p.Counter += ParamsForOpcode[JUMP_F] + 1
  }
}

func (p *Program) LessThan(a int, b int, reg int) {
  if a < b {
    p.Memory[reg] = 1
  } else {
    p.Memory[reg] = 0
  }
  p.Counter += ParamsForOpcode[LESS] + 1
}

func (p *Program) Equals(a int, b int, reg int) {
  if a == b {
    p.Memory[reg] = 1
  } else {
    p.Memory[reg] = 0
  }
  p.Counter += ParamsForOpcode[EQUALS] + 1
}
