package machine

import "fmt"

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

func NewEnv() Env {
	return Env{CurrentAddr: &Node{}}
}

type Env struct {
	CurrentAddr *Node
}

func (e *Env) GoNext() {
	e.CurrentAddr = e.CurrentAddr.next
}

func (e *Env) GoPrev() {
	e.CurrentAddr = e.CurrentAddr.prev
}

func (e Env) Incr() {
	e.CurrentAddr.Value++
}

func (e Env) Decr() {
	e.CurrentAddr.Value--
}

func (e Env) CurrentString() string {
	return string(e.CurrentAddr.Value)
}

func Hoge() {
	fmt.Println("hoge")
}
