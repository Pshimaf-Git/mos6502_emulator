package emulator

import (
	"fmt"
	"strings"
)

type Registers struct {
	Accumulator byte
	X           byte
	Y           byte

	SP byte // stack pointer
	P  byte // state register

	PC Word // program counter
}

func (r Registers) String() string {
	infos := []string{
		"Registers<",
		"Accumulator:%d",
		"X:%d", "Y:%d", "SP:%d", "PC:%d", "P:%d",
		">",
	}

	sep := " "

	s := strings.Join(infos, sep)
	return fmt.Sprintf(s, r.Accumulator, r.X, r.Y, r.P, r.PC, r.P)
}

type Flag = byte

const (
	FlagCarry     Flag = 1 << 0 // C 00000001
	FlagZero           = 1 << 1 // Z 00000010
	FlagInterrupt      = 1 << 2 // I 00000100
	FlagDecimal        = 1 << 3 // D 00001000
	FlagBreak          = 1 << 4 // B 00010000
	FlagUnused         = 1 << 5 // - 00100000
	FlagOverflow       = 1 << 6 // V 01000000
	FlagNegative       = 1 << 7 // N 10000000
)
