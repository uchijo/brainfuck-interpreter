package machine

import "fmt"

func NewEnv() Memory {
	return Memory{currentAddr: &node{}}
}

type Memory struct {
	currentAddr *node
}

func (e *Memory) GoNext() {
	e.currentAddr = e.currentAddr.next
}

func (e *Memory) GoPrev() {
	e.currentAddr = e.currentAddr.prev
}

func (e Memory) Incr() {
	e.currentAddr.Value++
}

func (e Memory) Decr() {
	e.currentAddr.Value--
}

func (e Memory) CurrentString() string {
	return string(e.currentAddr.Value)
}

func Hoge() {
	fmt.Println("hoge")
}
