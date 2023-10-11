package demos

import (
	"fmt"
)

// var mutex sync.Mutex
// var working_g sync.WaitGroup
var counter = 0

func incr() {
	for i := 0; i < 1000; i++ {

		counter++

	}
	Wg.Done()

}
func decr() {
	for i := 0; i < 1000; i++ {

		counter--
	}
	Wg.Done()

}

func RunIncrDec() {
	Wg.Add(2)
	go incr()
	go decr()
	Wg.Wait()
	fmt.Println("done !", counter)

}
