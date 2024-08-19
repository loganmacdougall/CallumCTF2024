package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var pointer = 0
var stack = []int32{}
var currInsDebug string
var runmode string

func main() {

	if len(os.Args) < 2 {
		fmt.Println("USAGE: starct [.strack file] [run|step|info]")
		os.Exit(-1)
	}

	runmode = "run"
	if len(os.Args) >= 3 {
		runmode = os.Args[2]
		if runmode != "run" && runmode != "step" && runmode != "info" {
			fmt.Println("Runmode can only be run, step, or info")
			os.Exit(-1)
		}
	}

	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var instructions = tokenizeInstructions(string(code))
	if runmode == "step" || runmode == "info" {
		fmt.Println(instructions)
	}
	var length = len(instructions)

	for {
		if pointer < 0 || pointer >= length {
			break
		}

		var currIns = instructions[pointer]
		currInsDebug = currIns
		if runmode == "step" || runmode == "info" {
			step_print()
		}

		if runmode == "step" {
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		} else if runmode == "info" {
			fmt.Print("\n")
		}

		if currIns[0] == '"' {
			var str = strings.Trim(currIns, "\"")
			pushStringToStack(str)
		} else if currIns[0] == '-' || (currIns[0] >= '0' && currIns[0] <= '9') {
			num, err := strconv.Atoi(currIns)
			if err != nil {
				panic(err)
			}

			if num > 2147483647 || num < -2147483648 {
				panic(fmt.Sprintf("Pusing number %d which doesn't fall within the range of an int32\n", num))
			}

			stack = append(stack, int32(num))
		} else {
			handleFunction(instructions)
		}

		pointer++
	}

	if runmode == "step" || runmode == "info" {
		dump()
	}
}

func tokenizeInstructions(instructionStr string) []string {
	instructionStr = strings.TrimRight(instructionStr, " \n")
	instructions := []string{}
	length := len(instructionStr)

	ins := []byte{}
	inString := false

	for i := 0; i < length; i++ {
		c := instructionStr[i]

		switch c {
		case '\\':
			switch instructionStr[i+1] {
			case 'n':
				ins = append(ins, '\n')
			case 't':
				ins = append(ins, '\t')
			case '"':
				ins = append(ins, '"')
			case '\\':
				ins = append(ins, '\\')
			default:
				panic(fmt.Sprintf("Trying to escape: %c\n", instructionStr[i+1]))
			}
			i++
			continue

		case '#':
			if inString {
				ins = append(ins, c)
			} else {
				for i < len(instructionStr) && instructionStr[i] != '\n' {
					i++
				}
			}
		case '"':
			inString = !inString
			ins = append(ins, c)
		case ' ', '\n', '\t':
			if inString {
				ins = append(ins, c)
			} else if len(ins) > 0 {
				instructions = append(instructions, string(ins))
				ins = []byte{}
			}
		case '\r':
		default:
			ins = append(ins, c)
		}

	}

	if len(ins) > 0 {
		instructions = append(instructions, string(ins))
	}
	return instructions
}

func pushStringToStack(str string) {
	length := len(str)
	padding := (4 - ((length + 1) % 4)) % 4
	finalLength := length + padding + 1
	byteArr := make([]byte, finalLength)

	byteArr[0] = 0
	for i := 0; i < length; i++ {
		byteArr[i+1] = str[length-1-i]
	}
	for i := length + 1; i < finalLength; i++ {
		byteArr[i] = 0
	}

	for i := 0; i < finalLength; i += 4 {
		var num = int(byteArr[i])
		num = num<<8 + int(byteArr[i+1])
		num = num<<8 + int(byteArr[i+2])
		num = num<<8 + int(byteArr[i+3])

		stack = append(stack, int32(num))
	}
}

func printFromStack() {
	if runmode == "step" || runmode == "info" {
		fmt.Print("Printed: ```\n")
	}

	finish := false
	for !finish {
		cell := consume(1)[0]
		msg := make([]byte, 5)
		for i := 0; i < 4; i++ {
			b := byte(cell)
			if b != 0 {
				msg = append(msg, b)
			} else if i == 3 {
				finish = true
			}
			cell = cell >> 8
		}
		fmt.Print(string(bytes.Trim(msg, "\x00")))
	}

	if runmode == "step" || runmode == "info" {
		fmt.Print("\n```\n\n")
	}
}

func consume(num int) []int32 {
	if len(stack) < num {
		fmt.Printf("\nTried consuming %d cells when there's only %d cells in the stack\n", num, len(stack))
		dump()
		os.Exit(22)
	}
	args := stack[len(stack)-num:]
	stack = stack[:len(stack)-num]

	slices.Reverse(args)

	return args
}

func handleFunction(instructions []string) {
	currIns := instructions[pointer]

	switch currIns {
	case "DUP":
		args := consume(1)
		stack = append(stack, stack[len(stack)-int(args[0]):]...)

	case "POP":
		args := consume(1)
		stack = stack[:len(stack)-int(args[0])]

	case "PRINT":
		printFromStack()

	case "ADD":
		args := consume(2)
		stack = append(stack, args[0]+args[1])

	case "SUB":
		args := consume(2)
		stack = append(stack, args[0]-args[1])

	case "MULT":
		args := consume(2)
		stack = append(stack, args[0]*args[1])

	case "MOD":
		args := consume(2)
		stack = append(stack, args[1]%args[0])

	case "RSFT":
		args := consume(2)
		stack = append(stack, args[1]>>args[0])

	case "LSFT":
		args := consume(2)
		stack = append(stack, args[1]<<args[0])

	case "MORE":
		args := consume(2)
		value := int32(1)
		if args[0] <= args[1] {
			value = 0
		}
		stack = append(stack, value)

	case "LESS":
		args := consume(2)
		value := int32(1)
		if args[0] >= args[1] {
			value = 0
		}
		stack = append(stack, value)

	case "EQ":
		args := consume(2)
		value := int32(1)
		if args[0] != args[1] {
			value = 0
		}
		stack = append(stack, value)

	case "NOT":
		args := consume(1)
		value := int32(0)
		if args[0] == 0 {
			value = 1
		}
		stack = append(stack, value)

	case "AND":
		args := consume(2)
		stack = append(stack, args[0]&args[1])

	case "OR":
		args := consume(2)
		stack = append(stack, args[0]|args[1])

	case "XOR":
		args := consume(2)
		stack = append(stack, args[0]^args[1])

	case "INV":
		args := consume(1)
		stack = append(stack, args[0]^int32(-1))

	case "INTSTRING":
		args := consume(1)
		pushStringToStack(strconv.Itoa(int(int32(args[0]))))

	case "CJUMP":
		args := consume(2)
		if args[1] != 0 {
			pointer += int(args[0]) - 1
		}

	default:
		fmt.Printf("\nInvalid function with name: %s\n", currIns)
		dump()
		os.Exit(5)
	}
}

func intToDebugString(num int32) string {
	dStr := ""
	lastByte := byte(0)
	sameByteCount := 0

	for i := 0; i <= 4; i++ {
		c := byte(0)
		if i == 4 {
			c = lastByte + 1
		} else {
			c = byte(num >> (24 - i*8))
		}

		if i == 0 {
			lastByte = c
			continue
		}

		if lastByte != c {
			sc := ""
			if lastByte >= 32 && lastByte <= 126 {
				sc = string(lastByte)
			} else if lastByte == '\n' {
				sc = "\\n"
			} else if lastByte == '\t' {
				sc = "\\t"
			} else if lastByte == '\r' {
				sc = "\\r"
			} else {
				sc = fmt.Sprintf("\\%d", int(lastByte))
			}

			if sameByteCount >= 1 {
				dStr += fmt.Sprintf("*%s", sc)
			} else {
				dStr += sc
			}

			if i < 4 {
				dStr += " "
			}
			sameByteCount = 0
		} else {
			sameByteCount += 1
		}

		lastByte = c
	}

	return dStr
}

func dump() {
	fmt.Printf("\nIns: %s (#%d)\n", currInsDebug, pointer)
	fmt.Println("Stack: ", stack)
}

func step_print() {
	stackStr := ""
	for i, num := range stack {
		if i == 0 {
			stackStr += "["
		} else {
			stackStr += "|"
		}
		stackStr += fmt.Sprintf(" %d (%s) ", num, intToDebugString(num))
	}

	if len(stackStr) > 1 {
		stackStr += "]"
	} else {
		stackStr += "|"
	}

	fmt.Print(stackStr, " ", currInsDebug, " (", pointer, ")")
}
