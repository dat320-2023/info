package main

import (
	"dat320/system"
	"fmt"
)

func test_io(i int) {
	fmt.Println(i)
}
func test_interrupt(i int) {
	fmt.Println(i)
}
func main() {

	cpu := system.CPU{Cycles: 10, IO_call: test_io, Time_interrupt: test_interrupt}
	fmt.Println(cpu)
}
