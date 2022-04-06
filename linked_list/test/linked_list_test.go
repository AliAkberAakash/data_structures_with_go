package test

import (
	"linked_list/linked_list"
	"testing"
)

func TestInsertionWasSuccess(t *testing.T) {
	var list linked_list.LinkedList

	list.Add(1, 0)

	if list.Length != 1 {
		t.Error("Item was not insertd to the list")
	}
	list.Add(-124, 0)
	list.Add(111, 1)
	list.Add(0, 2)

	if list.Length != 4 {
		t.Error("Item was not insertd to the list")
	}

}
