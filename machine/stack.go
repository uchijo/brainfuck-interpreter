package machine

import "errors"

type stack []int

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func (s *stack) pop() (int, error) {
	if s.empty() {
		return 0, errors.New("stack is empty")
	}
	retval := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return retval, nil
}

func (s stack) empty() bool {
	return len(s) <= 0
}

func (s stack) peek() int {
	return s[len(s)-1]
}
