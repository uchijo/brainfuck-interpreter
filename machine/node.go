package machine

type Node struct {
	next  *Node
	prev  *Node
	Value rune
}

func (n *Node) Next() *Node {
	if n.next == nil {
		n.next = &Node{
			prev: n,
		}
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == nil {
		n.prev = &Node{
			next: n,
		}
	}
	return n.prev
}

