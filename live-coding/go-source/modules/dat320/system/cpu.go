package system

type CPU struct {
	Time_interrupt func(int)
	IO_call        func(int)
	Cycles         int
}

func (cpu *CPU) Run(pr *Proc) {
	// for the given cycles
	// pick next task and Run it
	for cpu.Cycles > 0 {
		// pick next task
		t := pr.Tasks[pr.Current_Task]
		//if _, ok := t.(task.IO_Task); ok {
		// It is of type io
		// trap into os via io call
		// os sets blocked
		// os runs task on io_device
		//} else {
		// It is of cpu
		stepped := t.Run(cpu.Cycles)
		cpu.Cycles -= stepped

		//}
		if pr.Current_Task < len(pr.Tasks)+1 {
			pr.Current_Task++
		}
	}

}
