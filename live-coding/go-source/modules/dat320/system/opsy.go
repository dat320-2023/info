package system

import "fmt"

type OS struct {
	HT map[int]func()
}

func NewOS() *OS {
	return &OS{map[int]func(){0: Timer_interrupt_handler, 1: Sys_call_handler}}
}
func Timer_interrupt_handler() {
	fmt.Println("I am a time interrupt handler")
}
func Sys_call_handler() {
	fmt.Println("I am a system call handler")
}
