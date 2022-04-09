package pkg

import (
	"fmt"
	"linked_list/listerror"
)

type LinkedList struct {
	Head   *Node
	Length int
}

type Node struct {
	Data int
	Next *Node
}

func (list *LinkedList) Add(value int, position int) error {

	var err *listerror.OutOfBoundsError

	if position > list.Length || position < 0 {
		err = &listerror.OutOfBoundsError{Position: position}
		return err
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

	return nil
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

func (list *LinkedList) Delete(item int) {
	head := list.Head
	prev := head

	for head != nil {
		if head.Data == item {
			if prev == head {
				list.Head = nil

			} else {
				prev.Next = head.Next
			}
			list.Length--
			return
		}
		prev = head
		head = head.Next
	}
}
