package cpu_test

import (
	"testing"

	"github.com/Pshimaf-Git/mos6502_emulator/emulator"
	"github.com/stretchr/testify/assert"
)

func TestCPU_Execute(t *testing.T) {
	testCases := []struct {
		desc    string
		addr    emulator.Word
		program []byte

		expectedAcc byte
		expectedX   byte
		expectedY   byte

		expectedPC *emulator.Word
		expectedSP *byte
		expectedP  *byte

		expectedMem map[emulator.Word]byte
	}{
		{
			desc: "add to numbers",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.CLC,
				emulator.LDA_IMM, 0xA,
				emulator.ADC_IMM, 0x5,
				emulator.BRK,
			},

			expectedAcc: 15,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 6),
			expectedP:  defFlagWith(),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "add to numbers with carrry = 1",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.SEC,
				emulator.LDA_IMM, 0xA,
				emulator.ADC_IMM, 0x5,
				emulator.CLC,
				emulator.BRK,
			},

			expectedAcc: 16,
			expectedP:   defFlagWith(),
			expectedPC:  ptr(emulator.DEFAULT_MEM_START + 7),
			expectedSP:  ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "add two numbers and store result in memory",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.CLC,
				emulator.LDA_IMM, 0xA,
				emulator.ADC_IMM, 0x1,
				emulator.STA_ZP, 0x1,
				emulator.BRK,
			},

			expectedAcc: 11,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 8),
			expectedP:  defFlagWith(),
			expectedSP: ptr(emulator.STACK_START_VALUE),

			expectedMem: map[emulator.Word]byte{
				0x1: 11,
			},
		},

		{
			desc: "increment loop",
			addr: emulator.DEFAULT_MEM_START,

			program: []byte{
				emulator.LDX_IMM, 0,
				emulator.INX,
				emulator.CPX_IMM, 0xA,
				emulator.BNE, 0xFB, // -5 (jump back to INX)
				emulator.BRK,
			},

			expectedX: 10,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 8),
			expectedP:  defFlagWith(emulator.FlagCarry, emulator.FlagZero),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "LDA immediate and zero flag",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0x00,
				emulator.BRK,
			},
			expectedAcc: 0,
			expectedPC:  ptr(emulator.DEFAULT_MEM_START + 3),
			expectedP:   defFlagWith(emulator.FlagZero),
			expectedSP:  ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "LDA absolute indexed X with page crossing",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDX_IMM, 0xFF,
				emulator.LDA_ABS_X, 0x80, 0x20, // address 0x2080 + 0xFF = 0x217F
				emulator.BRK,
			},
			expectedAcc: 0, // memory at 0x217F is zero (default)
			expectedX:   0xFF,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 6),
			expectedP:  defFlagWith(emulator.FlagZero),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "STA zero page",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0x42,
				emulator.STA_ZP, 0x10,
				emulator.BRK,
			},
			expectedAcc: 0x42,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(),
			expectedSP: ptr(emulator.STACK_START_VALUE),

			expectedMem: map[emulator.Word]byte{0x10: 0x42},
		},
		{
			desc: "STA absolute",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0x42,
				emulator.STA_ABS, 0x00, 0x20,
				emulator.BRK,
			},
			expectedAcc: 0x42,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 6),
			expectedP:  defFlagWith(),
			expectedSP: ptr(emulator.STACK_START_VALUE),

			expectedMem: map[emulator.Word]byte{0x2000: 0x42},
		},
		{
			desc: "AND immediate clearing bits",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0b11001100,
				emulator.AND_IMM, 0b10101010,
				emulator.BRK,
			},
			expectedAcc: 0b10001000,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(emulator.FlagNegative),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "AND immediate two same numbers",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0b01100110,
				emulator.AND_IMM, 0b01100110,
				emulator.BRK,
			},
			expectedAcc: 0b01100110,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "ORA immediate setting bits",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0b11001100,
				emulator.ORA_IMM, 0b00110011,
				emulator.BRK,
			},
			expectedAcc: 0b11111111,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(emulator.FlagNegative),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "EOR immediate toggle bits",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0b11110000,
				emulator.EOR_IMM, 0b10101010,
				emulator.BRK,
			},
			expectedAcc: 0b01011010,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "ASL accumulator",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0b10000001,
				emulator.ASL_ACC,
				emulator.BRK,
			},
			expectedAcc: 0b00000010,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 4),
			expectedP:  defFlagWith(emulator.FlagCarry),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "LSR accumulator",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0b00000011,
				emulator.LSR_ACC,
				emulator.BRK,
			},
			expectedAcc: 0b00000001,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 4),
			expectedP:  defFlagWith(emulator.FlagCarry),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "ROL with carry set",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.SEC,
				emulator.LDA_IMM, 0b10000000,
				emulator.ROL_ACC,
				emulator.BRK,
			},
			expectedAcc: 0b00000001, // 0x80 << 1 = 0x00, plus oldCarry=1 → 0x01

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(emulator.FlagCarry),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "ROR with carry clear",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.CLC,
				emulator.LDA_IMM, 0b00000001,
				emulator.ROR_ACC,
				emulator.BRK,
			},
			expectedAcc: 0b00000000,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(emulator.FlagCarry, emulator.FlagZero),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "CMP sets carry if A >= mem",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0x05,
				emulator.CMP_IMM, 0x03,
				emulator.BRK,
			},
			expectedAcc: 0x05,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(emulator.FlagCarry),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "CMP sets zero if equal",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0x05,
				emulator.CMP_IMM, 0x05,
				emulator.BRK,
			},

			expectedAcc: 5,
			expectedPC:  ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:   defFlagWith(emulator.FlagCarry, emulator.FlagZero),
			expectedSP:  ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "BNE branch taken when Z=0",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0x01,
				emulator.CMP_IMM, 0x00,
				emulator.BNE, 0x01, // skip the next byte (2 bytes forward)
				emulator.BRK, // should not execute
				emulator.LDA_IMM, 0xFF,
				emulator.BRK,
			},
			expectedAcc: 0xFF, // after branch
			expectedPC:  ptr(emulator.DEFAULT_MEM_START + 10),
			expectedP:   defFlagWith(emulator.FlagCarry, emulator.FlagNegative),
			expectedSP:  ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "PHP and PLP preserve flags",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.SEC, // set carry
				emulator.PHP, // push status
				emulator.CLC, // clear carry
				emulator.PLP, // pop back
				emulator.BRK,
			},
			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(emulator.FlagCarry),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "Stack operations PHA and PLA",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDA_IMM, 0xAB,
				emulator.PHA,
				emulator.LDA_IMM, 0xCD,
				emulator.PLA,
				emulator.BRK,
			},
			expectedAcc: 0xAB, // PLA pulls the first pushed value
			expectedPC:  ptr(emulator.DEFAULT_MEM_START + 7),
			expectedP:   defFlagWith(emulator.FlagNegative),
			expectedSP:  ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "ADC with overflow",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.CLC,
				emulator.LDA_IMM, 0x70,
				emulator.ADC_IMM, 0x70,
				emulator.BRK,
			},
			expectedAcc: 0xE0,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 6),
			expectedP:  defFlagWith(emulator.FlagNegative, emulator.FlagOverflow),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "SBC with borrow",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.SEC, // set borrow (carry=1 means no borrow)
				emulator.LDA_IMM, 0x05,
				emulator.SBC_IMM, 0x02,
				emulator.BRK,
			},
			expectedAcc: 0x03,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 6),
			expectedP:  defFlagWith(emulator.FlagCarry),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "SBC with borrow needed (carry=0)",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.CLC, // clear carry -> borrow
				emulator.LDA_IMM, 0x05,
				emulator.SBC_IMM, 0x02,
				emulator.BRK,
			},
			expectedAcc: 0x02, // 5 - 2 - 1 = 2

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 6),
			expectedP:  defFlagWith(emulator.FlagCarry),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "INX and DEX",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDX_IMM, 0x00,
				emulator.DEX,
				emulator.INX,
				emulator.INX,
				emulator.BRK,
			},
			expectedX: 0x01,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 6),
			expectedP:  defFlagWith(),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
		{
			desc: "Zero page indexed wrapping (STX ZP,Y)",
			addr: emulator.DEFAULT_MEM_START,
			program: []byte{
				emulator.LDX_IMM, 0xFF,
				emulator.STX_ZP_Y, 0x01, // address = 0x01 + 0xFF = 0x100 -> wraps to 0x00
				emulator.BRK,
			},
			expectedMem: map[emulator.Word]byte{0x01: 0xFF},

			expectedX: 0xFF,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 5),
			expectedP:  defFlagWith(emulator.FlagNegative),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},

		// === TEST NON DOC COMMANDS ===
		{
			addr: emulator.DEFAULT_MEM_START,

			program: []byte{
				emulator.SEC,
				emulator.LDA_IMM, 0x10,
				emulator.STA_ZP, 0x20,
				emulator.ISB_ZP, 0x20,
				emulator.BRK,
			},

			expectedAcc: 0xFF,

			expectedPC: ptr(emulator.DEFAULT_MEM_START + 8),
			expectedP:  defFlagWith(emulator.FlagNegative),
			expectedSP: ptr(emulator.STACK_START_VALUE),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cpu := emulator.NewCPU()
			cpu.LoadProgram(tc.addr, tc.program)

			cpu.Execute()

			assert.Equal(t, tc.expectedAcc, cpu.Regs.Accumulator, "diff accumulator")
			assert.Equal(t, tc.expectedX, cpu.Regs.X, "diff X")
			assert.Equal(t, tc.expectedY, cpu.Regs.Y, "diff Y")

			if tc.expectedP != nil {
				assert.Equalf(t, *tc.expectedP, cpu.Regs.P, "diff flag register(P) want %d have %d", *tc.expectedP, cpu.Regs.P)
			}

			if tc.expectedPC != nil {
				assert.Equal(t, *tc.expectedPC, cpu.Regs.PC, "diff PC")
			}

			if tc.expectedSP != nil {
				assert.Equal(t, *tc.expectedSP, cpu.Regs.SP, "diff SP")
			}

			if tc.expectedMem != nil {
				for addr, expectedValue := range tc.expectedMem {
					v := cpu.Read(addr)
					assert.Equalf(t, expectedValue, v, "at addr 0x%X must be %d, but have %d", addr, expectedValue, v)
				}
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}

func defFlagWith(flags ...emulator.Flag) *emulator.Flag {
	fl := emulator.DEFAULT_FLAG
	for i := range flags {
		fl |= flags[i]
	}
	return &fl
}
