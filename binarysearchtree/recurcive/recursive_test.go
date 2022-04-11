package recurcive

import (
	"testing"
)

func TestInsert(t *testing.T) {
	var tree = &Node{
		Data: 0,
	}

	// Test left Node
	tree.Insert(10)
	if tree.RightNode.Data != 10 {
		t.Error("Data was not  inserted")
	}

	// Test right Node
	tree.Insert(-1)
	if tree.LeftNode.Data != -1 {
		t.Error("Data was not  inserted")
	}

}

func TestSearch(t *testing.T) {
	var tree = &Node{
		Data: 66,
	}

	tree.Insert(10)
	tree.Insert(100)
	tree.Insert(85)
	tree.Insert(52)
	tree.Insert(-11)
	tree.Insert(23)

	// Test if exists
	exists := tree.Search(-11)
	if exists == false {
		t.Error("Failed to find node")
	}

	// Test doesn't exist
	exists = tree.Search(50)
	if exists == true {
		t.Error("Found non existant node")
	}

}
