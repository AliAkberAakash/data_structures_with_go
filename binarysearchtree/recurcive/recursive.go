package recurcive

type Node struct {
	Data      int
	LeftNode  *Node
	RightNode *Node
}

func (head *Node) Insert(value int) {

	if value > head.Data {
		if head.RightNode == nil {
			head.RightNode = &Node{
				Data: value,
			}
		} else {
			head.RightNode.Insert(value)
		}
	} else {
		if head.LeftNode == nil {
			head.LeftNode = &Node{
				Data: value,
			}
		} else {
			head.LeftNode.Insert(value)
		}
	}
}

func (head *Node) Search(value int) bool {

	if head == nil {
		return false
	}

	if value > head.Data {
		return head.RightNode.Search(value)
	} else if value < head.Data {
		return head.LeftNode.Search(value)
	}

	return true
}
