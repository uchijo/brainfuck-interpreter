package machine

type node struct {
	next  *node
	prev  *node
	Value rune
}

func (n *node) Next() *node {
	if n.next == nil {
		n.next = &node{
			prev: n,
		}
	}
	return n.next
}

func (n *node) Prev() *node {
	if n.prev == nil {
		n.prev = &node{
			next: n,
		}
	}
	return n.prev
}
