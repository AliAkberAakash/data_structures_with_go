package singlylist

import (
	"fmt"
	"testing"
)

func TestInsertionWasSuccess(t *testing.T) {
	var list LinkedList

	err := list.Add(1, 0)

	if list.Length != 1 {
		t.Error("Item was not insertd to the list")
	}

	fmt.Print(err)
	if err != nil {
		t.Error("Failed to insert in the list")
	}

	list.Add(-124, 0)
	list.Add(111, 1)
	list.Add(0, 2)

	if list.Length != 4 {
		t.Error("Item was not insertd to the list")
	}
}

func TestListInsertionFailed(t *testing.T) {
	var list LinkedList

	err := list.Add(1, 1)

	if err == nil {
		t.Error("Failed to catch error")
	}
}

func TestFindItem(t *testing.T) {
	var list LinkedList

	list.Add(10, 0)
	list.Add(-1, 1)
	list.Add(100, 0)

	var foundPosition = list.Find(100)

	if foundPosition != 0 {
		t.Error("Item was not found in the list")
	}

	foundPosition = list.Find(0)

	if foundPosition != -1 {
		t.Error("Item should not exist in the list")
	}
}

func TestDeleteItem(t *testing.T) {
	var list LinkedList

	// Dlete last item
	list.Add(10, 0)
	list.Add(-1, 1)

	list.Delete(-1)

	if list.Find(-1) != -1 {
		t.Error("Item was not deleted")
	}

	if list.Length == 2 {
		t.Error("Item was not deleted")
	}

	// Delete first item
	list.Delete(10)

	if list.Find(10) != -1 {
		t.Error("Item was not deleted")
	}

	if list.Length == 1 {
		t.Error("Item was not deleted")
	}

	// dlete middle item
	list.Add(10, 0)
	list.Add(20, 0)
	list.Add(30, 0)

	list.Delete(20)

	if list.Find(20) != -1 {
		t.Error("Item was not deleted")
	}

	if list.Length != 2 {
		t.Error("Item was not deleted")
	}
}
