package machine

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Engine struct {
	input        string
	currentInstr int
	mem          memory
	loopStack    []int
	steps        int
	limit        int
}

func NewEngine(
	// input string
	input string,

	// steps limit to hang due to unexpected infinite loop. must be bigger than 1000
	limit int,
) Engine {
	if limit <= 1000 {
		limit = 1000
	}
	return Engine{
		input:        input,
		loopStack:    []int{},
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
		instr := e.input[e.currentInstr : e.currentInstr+1]

		// 実行
		switch instr {
		case "+":
			e.mem.Incr()
		case "-":
			e.mem.Decr()
		case ">":
			e.mem.GoNext()
		case "<":
			e.mem.GoPrev()
		case ".":
			fmt.Printf("%v", e.mem.CurrentString())
		case ",":
			var in string
			_, err := fmt.Scan(&in)
			if err != nil {
				return err
			}
			result, err := strconv.Atoi(in)
			switch {
			case err == nil:
				e.mem.SetInt(result)
			case len(in) == 1:
				r, _ := utf8.DecodeRuneInString(in)
				e.mem.SetRune(r)
			default:
				return errors.New("invalid input")
			}
		case "[":
			return errors.New("[ ] is not implemented")
		case "]":
			return errors.New("[ ] is not implemented")
		}

		// 命令レジスタを進める
		e.currentInstr++
	}
}
