package linked_list

import "fmt"

type LinkedList struct {
	Head   *Node
	Length int
}

type Node struct {
	Data int
	Next *Node
}

func (list *LinkedList) Add(value int, position int) {

	if position > list.Length || position < 0 {
		panic("Out of bounds")
	}

	var newNode = &Node{
		Data: value,
		Next: nil,
	}
	if position == 0 {
		if list.Head == nil {
			list.Head = newNode
		} else {
			temp := list.Head
			list.Head = newNode
			newNode.Next = temp
		}
	} else {
		currNode := list.Head.Next
		prevNode := list.Head
		for position-1 > 0 {
			prevNode = currNode
			currNode = currNode.Next
			position--
		}
		temp := currNode
		prevNode.Next = newNode
		newNode.Next = temp
	}

	list.Length++
}

func (list LinkedList) Find(v int) int {
	head := list.Head
	pos := 0
	for head != nil {
		if head.Data == v {
			return pos
		}
		head = head.Next
		pos++
	}
	return -1
}

func (list LinkedList) PrintList() {
	head := list.Head
	for head != nil {
		fmt.Printf("%d", head.Data)
		if head.Next != nil {
			fmt.Print(" => ")
		}
		head = head.Next
	}
	fmt.Println("")
}
