package machine

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Engine struct {
	input        []byte
	currentInstr int
	mem          memory
	loopStack    stack
	steps        int
	limit        int
}

func NewEngine(
	// input
	input []byte,

	// steps limit to hang due to unexpected infinite loop. must be bigger than 1000
	limit int,
) Engine {
	if limit <= 1000 {
		limit = 1000
	}
	return Engine{
		input:        input,
		loopStack:    stack{},
		mem:          newMemory(),
		currentInstr: 0,
		steps:        0,
		limit:        limit,
	}
}

func (e *Engine) Eval() error {
	for {
		// ステップ数確認
		e.steps++
		if e.steps > e.limit {
			return errors.New("step hit limit")
		}

		// 終了チェック
		if e.currentInstr >= len(e.input) {
			return nil
		}

		// 命令フェッチ
		instr := e.input[e.currentInstr]

		// 実行
		switch instr {
		case '+':
			e.mem.incr()
		case '-':
			e.mem.decr()
		case '>':
			e.mem.goNext()
		case '<':
			e.mem.goPrev()
		case '.':
			fmt.Printf("%v", e.mem.currentString())
		case ',':
			var in string
			_, err := fmt.Scan(&in)
			if err != nil {
				return err
			}
			result, err := strconv.Atoi(in)
			switch {
			case err == nil:
				e.mem.setInt(result)
			case len(in) == 1:
				r, _ := utf8.DecodeRuneInString(in)
				e.mem.setRune(r)
			default:
				return errors.New("invalid input")
			}
		case '[':
			e.loopStack.push(e.currentInstr)
			// 見てるアドレスに格納された値が0だったら対応するカッコを探す
			if e.mem.currentAddr.Value == 0 {
				braces := 0
				for i := e.currentInstr + 1; i < len(e.input); i++ {
					if e.input[i] == '[' {
						braces++
						continue
					}
					if e.input[i] == ']' {
						if braces == 0 {
							e.currentInstr = i
							break
						} else {
							braces--
						}
					}
				}
			}
		case ']':
			if e.mem.currentAddr.Value != 0 {
				e.currentInstr = e.loopStack.peek()
			} else {
				e.loopStack.pop()
			}
		}

		// 命令レジスタを進める
		e.currentInstr++
	}
}
