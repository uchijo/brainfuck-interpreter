package machine

func newMemory() memory {
	return memory{currentAddr: &node{}}
}

type memory struct {
	currentAddr *node
}

func (e *memory) goNext() {
	e.currentAddr = e.currentAddr.Next()
}

func (e *memory) goPrev() {
	e.currentAddr = e.currentAddr.Prev()
}

func (e memory) incr() {
	e.currentAddr.Value++
}

func (e memory) decr() {
	e.currentAddr.Value--
}

func (e memory) currentString() string {
	return string(e.currentAddr.Value)
}

func (e memory) setInt(v int) {
	e.currentAddr.Value = int32(v)
}

func (e memory) setRune(v rune) {
	e.currentAddr.Value = v
}
