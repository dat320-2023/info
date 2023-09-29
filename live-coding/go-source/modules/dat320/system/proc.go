package system

const (
	RUNNING = 0
	READY   = 1
	BLOCKED = 2
)

type Proc struct {
	Tasks        []*CPU_Task
	Status       int
	Current_Task int
}
