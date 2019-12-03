package intcode

import "fmt"

type OpCode int
const (
  ADD = 1
  MUL = 2
  HALT = 99
)

type Tape []int
type Program struct {
  Memory Tape
  Counter int
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
  switch op {
  case ADD:
    p.Add(p.Memory[p.Counter+1], p.Memory[p.Counter+2], p.Memory[p.Counter+3])
  case MUL:
    p.Multiply(p.Memory[p.Counter+1], p.Memory[p.Counter+2], p.Memory[p.Counter+3])
  case HALT:
    return true, nil
  default:
    return false, fmt.Errorf("unknown opcode: %d", op)
  }

  p.Counter += 4
  return false, nil
}

func (p *Program) Add(aReg int, bReg int, outputReg int) {
  a := p.Memory[aReg]
  b := p.Memory[bReg]
  p.Memory[outputReg] = a+b
}

func (p *Program) Multiply(aReg int, bReg int, outputReg int) {
  a := p.Memory[aReg]
  b := p.Memory[bReg]
  p.Memory[outputReg] = a*b
}


