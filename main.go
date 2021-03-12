package main

import (
	"fmt"
	"helloWorld/List"
)

func main() {
	var list List.List = List.NewLinckedList()
	for i := 0; i <= 5; i++ {
		list.Add(i)
	}
	ans, _ := list.GetByIndex(3)
	fmt.Println(ans)

	var list2 = List.NewLinckedList()

	fmt.Println(list2.GetFirst())
	fmt.Println(list2.GetLast())

	list2.AddAll(1, 2, 3, 4, 5, 6)

	fmt.Println(list2.GetFirst())
	fmt.Println(list2.GetLast())

	ans, _ = list.GetByIndex(4)
	fmt.Println(ans)
}
