package main

import (
	"Assembler/assembler"
	"os"
)

func main() {
	args := os.Args
	assembler.Translate(args[1])
}
