package assembler

import (
	"Assembler/code"
	"Assembler/parser"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Translate(path string) {
	parser.OpenFile(path)
	extension := filepath.Ext(path)
	if extension != ".asm" {
		panic("Invalid file extension: can only convert .asm files.")
	}
	outputPath := strings.Replace(path, extension, ".hack", 1)
	file, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Errorf("error creating output file: %s", err))
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	for parser.Advance() {
		switch parser.InsType() {
		case parser.A_INSTRUCTION:
			symbol := parser.Symbol()
			value, _ := strconv.ParseUint(symbol, 10, 15)
			binaryString := "0" + fmt.Sprintf("%015b", value)
			_, writeError := file.WriteString(binaryString + "\n")
			if writeError != nil {
				panic(writeError)
			}

		case parser.C_INSTRUCTION:
			dest := parser.Dest()
			destBinary := code.Dest(dest)
			comp := parser.Comp()
			compBinary := code.Comp(comp)
			jump := parser.Jump()
			jumpBinary := code.Jump(jump)

			binaryString := "111" + compBinary + destBinary + jumpBinary
			_, writeError := file.WriteString(binaryString + "\n")
			if writeError != nil {
				panic(writeError)
			}

		case parser.L_INSTRUCTION:
			// label instructions generate no code
			continue
		}
	}
}
