package hashmap

const size = 11

type HashMap struct {
	Items [size]Bucket
}

func (hashmap *HashMap) Insert(word string) {
	var k = hash(word)
	hashmap.Items[k].insert(word)
}

func (hashmap *HashMap) Search(word string) bool {
	var k = hash(word)
	return hashmap.Items[k].search(word)
}

func (hashmap *HashMap) Delete(word string) {
	var k = hash(word)
	hashmap.Items[k].detele(word)
}

type Bucket struct {
	Head *Node
}

type Node struct {
	Data string
	Next *Node
}

func (bucket *Bucket) insert(value string) {
	node := &Node{
		Data: value,
	}
	node.Next = bucket.Head
	bucket.Head = node
}

func (bucket *Bucket) search(value string) bool {

	var head = bucket.Head

	for head != nil {
		if head.Data == value {
			return true
		}
		head = head.Next
	}

	return false
}

func (bucket *Bucket) detele(value string) {

	head := bucket.Head
	var prev *Node

	for head != nil {
		if head.Data == value {
			if prev == nil {
				bucket.Head = nil
				return
			} else {
				prev.Next = head.Next
				return
			}
		}
		prev = head
		head = head.Next
	}
}

func hash(word string) int {
	sum := 0
	for char := range word {
		sum += char
	}
	return sum % size
}
