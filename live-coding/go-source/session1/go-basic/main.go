package main

import (
	"fmt"
	"go-basic/package1"
	"go-basic/package2"
)

func add(a int, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello")
	var session int = 1
	fmt.Printf("Hello and welcome to session number %d in DAT320\n", session)
	//fmt.Println(os.Args[1])
	fmt.Println("-------package 1-------")
	x := 1
	y := 2
	fmt.Printf(" values as parameters: %d + %d = %d\n", x, y, package1.Add(x, y))

	fmt.Printf(" pointers as parameters: %d + %d = %d\n", x, y, *package1.Add_pointers(&x, &y))
	fmt.Printf("value of x = %d, while its reference is at  %d\n", x, &x)

	fmt.Println("----package 2-------")
	geir := package2.Human{}
	geir.IncorrectSetName("Geir")
	geir.IncorrectSetAge(30)
	fmt.Println(geir.Describe())

	geir.SetName("Geir")
	geir.SetAge(30)
	fmt.Println(geir.Describe())

	fluffy := package2.Animal{Name: "Fluffy", Age: 3}
	fmt.Println(fluffy.Describe())

	species := []package2.Species{geir, fluffy}

	for i, s := range species {
		fmt.Println(s.Describe(), " at index ", i)
	}

}
