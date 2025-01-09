package parser

import (
	"testing"
)

func TestAdvance(t *testing.T) {
	OpenFile("testdata/test_parser.txt")
	Advance()
	if currentInstruction != "@13" {
		t.Errorf("current instruction = %s; want @13", currentInstruction)
	}
	Advance()
	if currentInstruction != "D=A+1" {
		t.Errorf("current instruction = %s; want D=A+1", currentInstruction)
	}
}

func TestCompWithDestAndJump(t *testing.T) {
	currentInstruction = "A=D-A;JNE"
	instructionType = C_INSTRUCTION
	if comp := Comp(); comp != "D-A" {
		t.Errorf("comp = %s; want D-A", comp)
	}
}

func TestCompNoDestNoJump(t *testing.T) {
	currentInstruction = "M-D"
	instructionType = C_INSTRUCTION
	if comp := Comp(); comp != "M-D" {
		t.Errorf("comp = %s; want M-D", comp)
	}

}

func TestCompWithDest(t *testing.T) {
	currentInstruction = "A=D&A"
	instructionType = C_INSTRUCTION
	if comp := Comp(); comp != "D&A" {
		t.Errorf("comp = %s; want D&A", comp)
	}
}

func TestCompWithJump(t *testing.T) {
	currentInstruction = "D-1;JMP"
	instructionType = C_INSTRUCTION
	if comp := Comp(); comp != "D-1" {
		t.Errorf("comp = %s; want D-1", comp)
	}
}

func TestDestEmpty(t *testing.T) {
	currentInstruction = "A+1;JGT"
	instructionType = C_INSTRUCTION
	if dest := Dest(); dest != "" {
		t.Errorf("dest = %s; want \"\" (empty string)", dest)
	}
}

func TestDestNotEmpty(t *testing.T) {
	currentInstruction = "D=A+1;JEQ"
	instructionType = C_INSTRUCTION
	if dest := Dest(); dest != "D" {
		t.Errorf("dest = %s; want D", dest)
	}
}
func TestSymbolAinstruction(t *testing.T) {
	currentInstruction = "@12"
	instructionType = A_INSTRUCTION
	symbol := Symbol()
	if symbol != "12" {
		t.Errorf("Symbol() = %s; want  12", symbol)
	}
}

func TestSymbolLinstruction(t *testing.T) {
	currentInstruction = "(LOOP)"
	instructionType = L_INSTRUCTION
	symbol := Symbol()
	if symbol != "LOOP" {
		t.Errorf("Symbol() = %s; want  LOOP", symbol)
	}
}
