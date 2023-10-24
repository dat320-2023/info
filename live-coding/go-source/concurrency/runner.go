package main

import (
	"dat320-conc/demos"
	"fmt"
)

func RunCounterWithCAS() {
	fmt.Println("*********** Lock Free Reliable Counter ************")
	sc := demos.NewSafeCounter()
	sc.Print = true
	sc.RunDemoCounterWithCAS()

}
func RunCounterWithTAS() {
	fmt.Println("!!!!!!!!!!!!! Not a reliable counter !!!!!!!!!!!!!")
	sc := demos.NewSafeCounter()
	sc.Print = true
	sc.RunDemoCounterWithTAS()

}

func main() {
	//demos.RunCondVar()
	//demos.RunIncrDec()
	//demos.RunRC()

	//sc := demos.NewSafeCounter()
	//sc.WG.Add(2)
	//go sc.Decrement(200)
	//go sc.Incr(100)
	//sc.WG.Wait()
	//fmt.Println("Final count=", sc.Val)
	demos.RunAgentDemo()
	//RunCounterWithTAS()
	//fmt.Println("----------------------------")
	//RunCounterWithCAS()

}
