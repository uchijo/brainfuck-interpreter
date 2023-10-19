package machine

import "fmt"

func newMemory() memory {
	return memory{currentAddr: &node{}}
}

type memory struct {
	currentAddr *node
}

func (e *memory) GoNext() {
	e.currentAddr = e.currentAddr.next
}

func (e *memory) GoPrev() {
	e.currentAddr = e.currentAddr.prev
}

func (e memory) Incr() {
	e.currentAddr.Value++
}

func (e memory) Decr() {
	e.currentAddr.Value--
}

func (e memory) CurrentString() string {
	return string(e.currentAddr.Value)
}

func Hoge() {
	fmt.Println("hoge")
}
