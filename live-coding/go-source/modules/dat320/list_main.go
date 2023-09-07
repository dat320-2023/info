package main

import (
	"dat320/lists"
	"fmt"
	"reflect"
)

func main() {
	var a any = "Job1"
	var b any = "Job2"
	var c any = 3

	item1 := lists.Thing{&a, nil}
	item2 := lists.Thing{&b, nil}
	item3 := lists.Thing{&c, nil}

	//circular list
	item1.Next = &item2
	item2.Next = &item3
	item3.Next = &item1

	fmt.Println("item 1:", *item1.Item, "type: ", reflect.TypeOf(item1), "holding: ", reflect.TypeOf(*item1.Item))
	fmt.Println("item 2:", *item2.Item, "type", reflect.TypeOf(item2), "holding: ", reflect.TypeOf(*item2.Item))
	fmt.Println("item 3:", *item3.Item, "type", reflect.TypeOf(item3), "holding: ", reflect.TypeOf(*item3.Item))

	item := item1
	fmt.Print(*item.Item, "-->")
	for i := 0; i < 11; i++ {
		item = *item.Next
		if i == 10 {
			fmt.Print(*item.Item, "\n")
			break
		}
		fmt.Print(*item.Item, "-->")
	}

}
