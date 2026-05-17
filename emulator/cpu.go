// Package emulator ...
package emulator

import (
	"errors"
	"fmt"
	"strings"
)

const (
	STACK_START_PAGE       = 0x0100
	STACK_START_VALUE byte = 0xFD
	DEFAULT_FLAG      byte = FlagUnused | FlagInterrupt
	DEFAULT_MEM_START Word = 0x8000
)

var ErrUnknownIntrustion = errors.New("unknown instruction")

type CPU struct {
	Regs       Registers
	mem        *Memory
	instrTable map[byte]instruction
}

func NewCPU() *CPU {
	cpu := &CPU{
		Regs:       Registers{},
		mem:        &Memory{},
		instrTable: map[byte]instruction{},
	}

	cpu.Reset()
	cpu.initTable()

	return cpu
}

func (cpu *CPU) Info(sep string) string {
	infos := []string{
		"%s",
		"Flags<",
		"Carry:%d", "Zero:%t", "Overflow:%t", "Negative:%t",
		"Decimal:%t", "Break:%t", "Unused:%t", "Interrupt:%t",
		">",
	}

	s := strings.Join(infos, sep)
	return fmt.Sprintf(s, cpu.Regs.String(), toInt(cpu.GetFlag(FlagCarry)), cpu.GetFlag(FlagZero), cpu.GetFlag(FlagOverflow), cpu.GetFlag(FlagNegative),
		cpu.GetFlag(FlagDecimal), cpu.GetFlag(FlagBreak), cpu.GetFlag(FlagUnused), cpu.GetFlag(FlagInterrupt),
	)
}

func (cpu *CPU) String() string {
	return "CPU<info: " + cpu.Info(" ")
}

func (cpu *CPU) LoadMemory(newMem *Memory) {
	cpu.mem = newMem
}

func (cpu *CPU) FetchWord() Word {
	low := Word(cpu.FetchByte())
	high := Word(cpu.FetchByte())

	return (high << 8) | low
}

func (cpu *CPU) FetchByte() byte {
	b := cpu.Read(cpu.Regs.PC)
	cpu.incPC()
	return b
}

func (cpu *CPU) Execute() {
	for opcode := cpu.FetchByte(); opcode != BRK; opcode = cpu.FetchByte() {
		cpu.Step(opcode)
	}
}

func (cpu *CPU) ExecuteWithLimit(limit int) (stoppedByLimit bool) {
	opcode := cpu.FetchByte()
	i := 0

	for ; i <= limit && opcode != BRK; i++ {
		cpu.Step(opcode)

		opcode = cpu.FetchByte()
	}

	if opcode == BRK {
		cpu.incPC()
	}

	return i == limit
}

func (cpu *CPU) QuickRun(program []byte) {
	cpu.LoadProgram(DEFAULT_MEM_START, program)
	cpu.Execute()
}

func (cpu *CPU) Step(opcode byte) {
	if opcode == BRK {
		return
	}

	inst, ok := cpu.instrTable[opcode]
	if !ok {
		panic(fmt.Sprintf("Unknown opcode: 0x%02X at PC=0x%04X", opcode, cpu.Regs.PC-1))
	}

	inst.exec(inst.mode)
}

func (cpu *CPU) Reset() {
	low := Word(cpu.Read(0xFFFC))
	high := Word(cpu.Read(0xFFFD))
	cpu.Regs.PC = Word(high)<<8 | Word(low)

	cpu.Regs.SP = STACK_START_VALUE
	cpu.Regs.P = DEFAULT_FLAG
}

func (cpu *CPU) LoadProgram(address Word, data []byte) {
	for i, b := range data {
		cpu.Write(address+Word(i), b)
	}

	cpu.Write(0xFFFC, byte(address&0xFF))
	cpu.Write(0xFFFD, byte(address>>8))

	cpu.Reset()
}

func (cpu *CPU) updateZN(val byte) {
	cpu.setStatusFlag(FlagZero, val == 0)
	cpu.setStatusFlag(FlagNegative, (val&0x80) != 0)
}

func (cpu *CPU) setStatusFlag(fl Flag, cond bool) {
	if cond {
		cpu.Regs.P |= fl
	} else {
		cpu.Regs.P &= ^fl
	}
}

func (cpu *CPU) GetFlag(fl Flag) bool {
	return (cpu.Regs.P & fl) != 0
}

func (cpu *CPU) incPC() {
	cpu.Regs.PC++
}

func (cpu *CPU) Read(addr Word) byte {
	return cpu.mem.Data[addr]
}

func (cpu *CPU) Write(addr Word, val byte) {
	cpu.mem.Data[addr] = val
}

func (cpu *CPU) Push(val byte) {
	addr := STACK_START_PAGE + Word(cpu.Regs.SP)
	cpu.Write(addr, val)
	cpu.Regs.SP--
}

func (cpu *CPU) Pop() byte {
	cpu.Regs.SP++
	addr := 0x0100 + Word(cpu.Regs.SP)
	val := cpu.Read(addr)
	return val
}

func toInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
