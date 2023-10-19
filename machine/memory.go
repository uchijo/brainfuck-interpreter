package machine

import "fmt"

func newMemory() memory {
	return memory{currentAddr: &node{}}
}

type memory struct {
	currentAddr *node
}

func (e *memory) GoNext() {
	e.currentAddr = e.currentAddr.Next()
}

func (e *memory) GoPrev() {
	e.currentAddr = e.currentAddr.Prev()
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

func (e memory) SetInt(v int) {
	e.currentAddr.Value = int32(v)
}

func (e memory) SetRune(v rune) {
	e.currentAddr.Value = v
}

func Hoge() {
	fmt.Println("hoge")
}
