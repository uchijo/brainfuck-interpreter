package machine

import "fmt"

func NewEnv() Env {
	return Env{currentAddr: &Node{}}
}

type Env struct {
	currentAddr *Node
}

func (e *Env) GoNext() {
	e.currentAddr = e.currentAddr.next
}

func (e *Env) GoPrev() {
	e.currentAddr = e.currentAddr.prev
}

func (e Env) Incr() {
	e.currentAddr.Value++
}

func (e Env) Decr() {
	e.currentAddr.Value--
}

func (e Env) CurrentString() string {
	return string(e.currentAddr.Value)
}

func Hoge() {
	fmt.Println("hoge")
}
