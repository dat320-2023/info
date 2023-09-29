package compsys

import (
	"fmt"
)

type Ops struct {
	Incr  func()
	Mul   func()
	Store func()
}

type IR struct {
	OpCode int
}

type CPU struct {
	MyIR IR
}

func (cpu *CPU) RunCycle() {
	cpu.Fetch()
	cpu.Decode()
	cpu.Execute()
}
func (cpu *CPU) Fetch() {

	fmt.Println("Fetching")
	cpu.MyIR.OpCode = 1
}
func (cpu *CPU) Decode() {
	fmt.Println("Decoding")
}
func (cpu *CPU) Execute() {
	fmt.Println("Executing")
}
