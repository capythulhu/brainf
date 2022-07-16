package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var memory = []byte{0}
var program []byte
var mp = 0

func runOp(i int, c byte) int {
	switch c {
	case '>':
		mp++
		memory = append(memory, 0)
	case '<':
		if mp == 0 {
			panic("out of bounds")
		}
		mp--
	case '+':
		memory[mp]++
	case '-':
		memory[mp]--
	case '.':
		fmt.Printf("%c", memory[mp])
	case ',':
		var s string
		fmt.Scan(&s)
		memory[mp] = byte(s[0])
	case '[':
		k := i + 1
		for ; program[k] != ']'; k++ {
		}
		for memory[mp] != 0 {
			for j := i + 1; j < k; j++ {
				j += runOp(j, program[j])
			}
		}
		return k - i
	case ']':
		panic("unmatched ]")
	case '\n', '\r', '\t', ' ':
	case '#':
		j := i + 1
		for j < len(program) && program[j] != '\n' {
			j++
		}
		return j - i
	default:
		panic("unknown char: " + string(c))
	}
	return 0
}

func main() {
	p, err := os.ReadFile(os.Args[1])

	exten := filepath.Ext(os.Args[1])

	if exten != ".bf" {
		panic("Format not accept! " + exten)
	}

	if err != nil {
		panic(err)
	}

	program = p

	for i := 0; i < len(program); i++ {
		i += runOp(i, program[i])
	}

	fmt.Println()
}
