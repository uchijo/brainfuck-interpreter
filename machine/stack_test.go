package machine

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Stack: test push result 1", func(t *testing.T) {
		s := stack{}
		s.push(10)
		if len(s) != 1 || s[0] != 10 {
			t.Errorf("push failed.")
		}
	})
	t.Run("Stack: test push result 2", func(t *testing.T) {
		s := stack{}
		s.push(10)
		s.push(20)
		s.push(40)
		if len(s) != 3 || s[0] != 10 || s[1] != 20 || s[2] != 40 {
			t.Errorf("push failed.")
		}
	})

	t.Run("Stack: test pop result", func(t *testing.T) {
		s := stack{}
		s.push(10)
		s.push(20)
		s.push(30)
		if result, _ := s.pop(); result != 30 {
			t.Errorf("expected to get 30, but got %v", result)
		}
		fmt.Println(s)
		if result, _ := s.pop(); result != 20 {
			t.Errorf("expected to get 20, but got %v", result)
		}
		s.pop()
		s.pop()
	})
}
