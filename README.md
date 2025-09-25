# Go Assembler

An assembler for the Hack machine language written in Go.

The HDL for the CPU can be found under ``CPU.hdl``

## Setup
1. Clone the repository and navigate to the project root.
2. Build the executable for your system
```sh
go build -o hack-assembler .
```
3. Run the assembler by specifying the path of a ``.asm`` file
```sh
./hack-assembler example.asm
```


You will now find a generated ``.hack`` file under the same path


# Output examples

| Assembly   | Machine Code         |
|------------|---------------------|
| @256       | 0000000100000000    |
| D=A        | 1110110000010000    |
| @SP        | 0000000000000000    |
| M=D        | 1110001100001000    |
| @133       | 0000000010000101    |
| 0;JMP      | 1110101010000111    |
| @R15       | 0000000000001111    |
| M=D        | 1110001100001000    |

