package main

import (
	"fmt"
	"session2/compsys"
)

func Increment() {
	fmt.Println("will increment something")
}
func main() {
	fmt.Println("Hello, we will try to simulate a CPU")
	cpu := compsys.CPU{}
	cpu.RunCycle()
	ops := compsys.Ops{Incr: Increment}
	ops.Incr()

}
