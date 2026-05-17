# MOS6502 Emulator

A **MOS Technology 6502 CPU emulator** written in Go. This package accurately simulates the behavior of the classic 6502 microprocessor, including its instruction set, addressing modes, flags, and stack operations. It also supports several **unofficial (illegal) opcodes** commonly found in real-world software and demos.

## Features

- **Full 6502 instruction set** (official opcodes)
- **All addressing modes**:
  - Implied, Immediate, Relative
  - Zero Page, Zero Page X, Zero Page Y
  - Absolute, Absolute X, Absolute Y
  - Indirect, Indirect X, Indirect Y
- **Flag support**: Carry (C), Zero (Z), Interrupt Disable (I), Decimal (D), Break (B), Overflow (V), Negative (N)
- **Stack operations** 
- **Interrupt handling** (IRQ/BRK) and reset vector (`$FFFC`-`$FFFD`)
- **Unofficial opcodes** (e.g., `ISB` – increment then subtract)
- **Memory** – 64KB addressable space
- **Test suite** – extensive unit tests covering most instructions and edge cases
- **Easy integration** – use as a Go library in your own emulator projects (NES, Apple II, etc.)

## Getting Started

### Prerequisites

- Go 1.24 or later (see `go.mod`)

### Installation

```bash
go get github.com/Pshimaf-Git/mos6502_emulator
```


### Example 

```go 
package main

import (
    "fmt"
    "github.com/Pshimaf-Git/mos6502_emulator/emulator"
)

func main() {
  cpu := emulator.NewCPU()
   // Simple program: LDA #$0A, ADC #$05, BRK
  program := []byte{
        emulator.LDA_IMM, 0xA,
        emulator.ADC_IMM, 0x5,
        emulator.BRK,
  }

  cpu.QuickStart(program)
  fmt.Printf("Result: %d\n", cpu.Regs.Accumulator) // Output: 15
}
```


### Project structure

```
```
.
├── addr.go           – addressing mode functions
├── arithmetic.go     – ADC, SBC, INC, DEC, etc.
├── branch.go         – jumps, branches, subroutine calls
├── cmp.go            – compare instructions
├── cpu.go            – core CPU logic (fetch, execute, reset)
├── data_transfer.go  – LDA, STA, LDX, etc.
├── flags_op.go       – flag control (SEC, CLC, ...)
├── instr_table.go    – opcode → instruction mapping
├── intrs.go          – instruction and opcode constants
├── logical.go        – AND, ORA, EOR, BIT
├── memory.go         – 64KB memory model
├── nop.go            – NOP implementation
├── registers.go      – CPU registers and flag definitions
├── rotates.go        – ASL, LSR, ROL, ROR
├── stack_ops.go      – PHA, PHP, PLA, PLP
├── word.go           – type alias for uint16
├── cpu_test.go       – exhaustive test suite
└── go.mod            – module definition
```
```
```
```

### Supported Instructions

- Load/Store: `LDA`, `LDX`, `LDY`, `STA`, `STX`, `STY`

- Arithmetic: `ADC`, `SBC`, `INC`, `DEC`, `INX`, `INY`, `DEX`, `DEY`

- Logical: `AND`, `ORA`, `EOR`, `BIT`

- Shifts/Rotates: `ASL`, `LSR`, `ROL`, `ROR`

- Branches: `BCC`, `BCS`, `BEQ`, `BNE`, `BMI`, `BPL`, `BVC`, `BVS`

- Jumps/Subroutines: `JMP`, `JSR`, `RTS`, `RTI`

- Stack: `PHA`, `PHP`, `PLA`, `PLP`

- Flags: `CLC`, `SEC`, `CLI`, `SEI`, `CLD`, `SED`, `CLV`

- Transfer: `TAX`, `TAY`, `TXA`, `TYA`, `TSX`, `TXS`

- Compare: `CMP`, `CPX`, `CPY`

- Misc: `NOP`, `BRK`

- Unofficial: `ISB`

[see MOS6502 commands system](https://emuverse.ru/wiki/MOS_Technology_6502/%D0%A1%D0%B8%D1%81%D1%82%D0%B5%D0%BC%D0%B0_%D0%BA%D0%BE%D0%BC%D0%B0%D0%BD%D0%B4)

### Contributing

Contributions are very very welcome! Feel free to open an issue or pull request for missing instructions, bug fixes, or performance improvements.

### License

This project is open source and available under the [MIT License](./LICENSE).
