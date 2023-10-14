package main

import (
	"dat320-conc/demos"
)

func main() {
	demos.RunCondVar()
	//demos.RunIncrDec()
	//demos.RunRC()

	//sc := demos.NewSafeCounter()
	//sc.WG.Add(2)
	//go sc.Decrement(200)
	//go sc.Incr(100)
	//sc.WG.Wait()
	//fmt.Println("Final count=", sc.Val)

}
