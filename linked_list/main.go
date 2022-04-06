package main

import (
	"fmt"
	"linked_list/pkg"
)

func main() {
	var list pkg.LinkedList

	list.Add(5, 0)
	list.Add(999, 0)
	list.Add(7, 0)
	//fmt.Printf("List length is %d\n", myList.Length)
	//myList.PrintList()
	list.Add(-1, 0)
	list.PrintList()
	//myList.Add(7, 0)

	fmt.Println(list.Find(100))

	//list.Add(999, -1)
}
