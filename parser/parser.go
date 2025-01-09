package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type InstructionType int

const (
	A_INSTRUCTION = iota
	C_INSTRUCTION
	L_INSTRUCTION
)

var instructionName = map[InstructionType]string{
	A_INSTRUCTION: "A_INSTRUCTION",
	C_INSTRUCTION: "C_INSTRUCTION",
	L_INSTRUCTION: "L_INSTRUCTION",
}

var (
	currentInstruction string
	instructionType    InstructionType
	file               *os.File
	scanner            *bufio.Scanner
)

func OpenFile(path string) {
	var err error
	file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)
}

func Advance() (hasInstruction bool) {
	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// skip over lines with comments and empty lines
		for line == "" || line[0] == '/' {
			if scanner.Scan() {
				line = strings.TrimSpace(scanner.Text())
			} else {
				_ = file.Close()
				return false
			}
		}
		line = cleanLine(line)
		currentInstruction = line
		instructionType = matchInstructionType()
		return true
	} else {
		// close the file when the end is reached
		_ = file.Close()
		return false
	}
}

// Symbol returns the symbol of an A or L instruction.
func Symbol() string {
	switch instructionType {
	case L_INSTRUCTION:
		if idx := strings.IndexRune(currentInstruction, ')'); idx != -1 {
			return currentInstruction[1:idx]
		} else {
			panic(fmt.Errorf("missing ')' for label instruction: %s", currentInstruction))
		}
	case A_INSTRUCTION:
		return currentInstruction[1:]
	default:
		panic("invalid instruction type for Symbol() call")
	}
}

// Dest returns the 'dest' part of the current instruction (only when it is a C instruction).
// If 'dest' is empty, then it returns an empty string.
func Dest() string {
	if instructionType == C_INSTRUCTION {
		if idx := strings.IndexRune(currentInstruction, '='); idx != -1 {
			return currentInstruction[:idx]
		} else {
			return ""
		}
	} else {
		panic("invalid instruction type for Dest() call")
	}
}

// Comp returns the 'comp' part of the current instruction (only when it is a C instruction).
func Comp() string {
	if instructionType == C_INSTRUCTION {
		if idxEquals := strings.IndexRune(currentInstruction, '='); idxEquals != -1 {
			if idxSemicolon := strings.IndexRune(currentInstruction, ';'); idxSemicolon != -1 {
				return currentInstruction[idxEquals+1 : idxSemicolon]
			} else {
				return currentInstruction[idxEquals+1:]
			}
		} else {
			if idxSemicolon := strings.IndexRune(currentInstruction, ';'); idxSemicolon != -1 {
				return currentInstruction[:idxSemicolon]
			} else {
				return currentInstruction
			}
		}

	} else {
		panic("invalid instruction type for Dest() call")
	}
}

// Jump returns the 'jmp' part of a C instruction.
// if it is empty, it returns an empty string.
func Jump() string {
	if instructionType == C_INSTRUCTION {
		if idx := strings.IndexRune(currentInstruction, ';'); idx != -1 {
			return currentInstruction[idx+1:]
		} else {
			return ""
		}
	} else {
		panic("invalid instruction type for Jump() call")
	}
}

// cleanLine takes out the trailing comments that reside on a line containing instructions
func cleanLine(line string) string {
	if idx := strings.IndexRune(line, '/'); idx != -1 {
		return strings.TrimSpace(line[:idx])
	}
	return line
}

// matchInstructionType returns the current instruction type, based on the value of currentInstruction
func matchInstructionType() InstructionType {
	switch currentInstruction[0] {
	case '@':
		return A_INSTRUCTION
	case '(':
		return L_INSTRUCTION
	default:
		// assumes that the code is valid
		return C_INSTRUCTION
	}
}
