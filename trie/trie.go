package trie

const childrenLength = 26

type Node struct {
	Children   [childrenLength]*Node
	IsComplete bool
}

type Trie struct {
	Head *Node
}

func InitTrie() *Trie {
	return &Trie{
		Head: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	wordLength := len(word)

	currentNode := t.Head

	for i := 0; i < wordLength; i++ {
		var charIndex int = int(word[i]) - int('a')
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}
		currentNode = currentNode.Children[charIndex]
	}

	t.Head.IsComplete = true
}

func (t *Trie) Search(word string) bool {
	wordLength := len(word)

	currentNode := t.Head

	for i := 0; i < wordLength; i++ {
		var charIndex int = int(word[i]) - int('a')
		if currentNode.Children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.Children[charIndex]
	}

	if t.Head.IsComplete == true {
		return true
	}

	return false
}
