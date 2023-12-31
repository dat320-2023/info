package demos

import (
	"fmt"
)

// var mutex sync.Mutex
// var working_g sync.WaitGroup
var counter = 0

func incr() {

	for i := 0; i < 1000; i++ {
		Mu.Lock()
		counter++ // critical section
		Mu.Unlock()

	}
	fmt.Println("incr done", counter)

	Wg.Done()

}
func decrement() {

	for i := 0; i < 1000; i++ {
		Mu.Lock()
		counter-- //critical section
		Mu.Unlock()
	}
	fmt.Println("Decrement done", counter)

	Wg.Done()

}

func RunIncrDec() {
	Wg.Add(2)
	go incr()
	go decrement()
	Wg.Wait()
	fmt.Println("done !", counter)

}
