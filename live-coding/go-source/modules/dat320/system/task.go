package system

import "math"

const (
	IO      = 0
	COMPUTE = 1
)

type Task interface {
	Run(cycles int) int
}
type IO_Task struct {
	Cycles    int
	Remaining int
}
type CPU_Task struct {
	Cycles    int
	Remaining int
}

func (task *IO_Task) Run(cycles int) int {
	steps := int(math.Min(float64(task.Remaining), float64(cycles)))
	task.Cycles += steps
	task.Remaining -= steps
	return steps
}
func (task *CPU_Task) Run(cycles int) int {
	steps := int(math.Min(float64(task.Remaining), float64(cycles)))
	task.Cycles += steps
	task.Remaining -= steps
	return steps
}
