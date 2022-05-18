package trie

import "testing"

func TestInit(t *testing.T) {
	trie := InitTrie()

	if trie == nil {
		t.Error("Trie was not created")
	}

	if trie.Head == nil {
		t.Error("Trie was not created")
	}

	len := len(trie.Head.Children)
	if len != childrenLength {
		t.Errorf("Expected %v got %v", childrenLength, len)
	}
}

func TestInsert(t *testing.T) {
	trie := InitTrie()

	word := "golang"

	trie.Insert(word)

	expectedVal := true
	isEmpty := false

	/// Looping through all nodes in children to check if
	/// every one of them is null
	/// If the word was inserted, there would be atleast one node
	/// will be not null
	for _, val := range trie.Head.Children {
		if val != nil {
			isEmpty = true
		}
	}

	if isEmpty != expectedVal {
		t.Errorf("Expected %v got %v", expectedVal, isEmpty)
	}

}

func TestSearchFound(t *testing.T) {
	trie := InitTrie()

	word := "golang"

	trie.Insert(word)

	expectedVal := true
	found := trie.Search(word)

	if found != expectedVal {
		t.Errorf("Expected %v got %v", expectedVal, found)
	}

}

func TestSearchNot(t *testing.T) {
	trie := InitTrie()

	word := "golang"

	trie.Insert(word)

	expectedVal := false
	found := trie.Search("javaScript")

	if found != expectedVal {
		t.Errorf("Expected %v got %v", expectedVal, found)
	}

}
