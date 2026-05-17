package emulator

func (cpu *CPU) initTable() {
	cpu.instrTable[LDA_IMM] = newIntr("LDA_IMM", cpu.addrImmediate, cpu.lda)
	cpu.instrTable[LDA_ZP] = newIntr("LDA_ZP", cpu.addrZeroPage, cpu.lda)
	cpu.instrTable[LDA_ZP_X] = newIntr("LDA_ZP_X", cpu.addrZeroPageX, cpu.lda)
	cpu.instrTable[LDA_ABS] = newIntr("LDA_ABS_S", cpu.addrAbsolute, cpu.lda)
	cpu.instrTable[LDA_ABS_X] = newIntr("LDA_ABS_X", cpu.addrAbsoluteX, cpu.lda)
	cpu.instrTable[LDA_ABS_Y] = newIntr("LDA_ABS_Y", cpu.addrAbsoluteY, cpu.lda)
	cpu.instrTable[LDA_IND_X] = newIntr("LDA_IND_X", cpu.addrIndirectX, cpu.lda)
	cpu.instrTable[LDA_IND_Y] = newIntr("LDA_IND_Y", cpu.addrIndirectY, cpu.lda)

	cpu.instrTable[LDX_IMM] = newIntr("LDX_IMM", cpu.addrImmediate, cpu.ldx)
	cpu.instrTable[LDX_ZP] = newIntr("LDX_ZP", cpu.addrZeroPage, cpu.ldx)
	cpu.instrTable[LDX_ABS] = newIntr("LDX_ABS", cpu.addrAbsolute, cpu.ldx)
	cpu.instrTable[LDX_ZP_Y] = newIntr("LDX_ZP_Y", cpu.addrZeroPageY, cpu.ldx)
	cpu.instrTable[LDX_ABS_Y] = newIntr("LDX_ABS_Y", cpu.addrAbsoluteY, cpu.ldx)

	cpu.instrTable[LDY_IMM] = newIntr("LDY_IMM", cpu.addrImmediate, cpu.ldy)
	cpu.instrTable[LDY_ZP] = newIntr("LDY_ZP", cpu.addrZeroPage, cpu.ldy)
	cpu.instrTable[LDY_ABS] = newIntr("LDY_ABS", cpu.addrAbsolute, cpu.ldy)
	cpu.instrTable[LDY_ZP_X] = newIntr("LDY_ZP_X", cpu.addrZeroPageX, cpu.ldy)
	cpu.instrTable[LDY_ABS_X] = newIntr("LDY_ABS_X", cpu.addrAbsoluteX, cpu.ldy)

	// ADC
	cpu.instrTable[ADC_IMM] = newIntr("ADC_IMM", cpu.addrImmediate, cpu.adc)
	cpu.instrTable[ADC_ZP] = newIntr("ADC_ZP", cpu.addrZeroPage, cpu.adc)
	cpu.instrTable[ADC_ZP_X] = newIntr("ADC_ZP_X", cpu.addrZeroPageX, cpu.adc)
	cpu.instrTable[ADC_ABS_X] = newIntr("ADC_ABS_X", cpu.addrAbsoluteX, cpu.adc)
	cpu.instrTable[ADC_ABS_Y] = newIntr("ADC_ABS_Y", cpu.addrAbsoluteY, cpu.adc)
	cpu.instrTable[ADC_IND_X] = newIntr("ADC_IND_X", cpu.addrIndirectX, cpu.adc)
	cpu.instrTable[ADC_IND_Y] = newIntr("ADC_IND_Y", cpu.addrIndirectY, cpu.adc)

	// STA
	cpu.instrTable[STA_ZP] = newIntr("STA_ZP", cpu.addrZeroPage, cpu.sta)
	cpu.instrTable[STA_ZP_X] = newIntr("STA_ZP_X", cpu.addrZeroPageX, cpu.sta)
	cpu.instrTable[STA_ABS] = newIntr("STA_ABS", cpu.addrAbsolute, cpu.sta)
	cpu.instrTable[STA_ABS_X] = newIntr("STA_ABS_X", cpu.addrAbsoluteX, cpu.sta)
	cpu.instrTable[STA_IND_X] = newIntr("STA_IND_X", cpu.addrIndirectX, cpu.sta)
	cpu.instrTable[STA_ABS_Y] = newIntr("STA_ABS_Y", cpu.addrAbsoluteY, cpu.sta)
	cpu.instrTable[STA_IND_Y] = newIntr("STA_IND_Y", cpu.addrIndirectY, cpu.sta)

	// JMP
	cpu.instrTable[JMP_ABS] = newIntr("JMP_ABS", cpu.addrAbsolute, cpu.jmp)
	cpu.instrTable[JMP_IND] = newIntr("JMP_IND", cpu.addrIndirect, cpu.jmp)

	cpu.instrTable[JSR_ABS] = newIntr("JSR_ABS", cpu.addrAbsolute, cpu.jsr)

	// INC
	cpu.instrTable[INC_ZP] = newIntr("INC_ZP", cpu.addrZeroPage, cpu.inc)
	cpu.instrTable[INC_ABS] = newIntr("INC_ABS", cpu.addrAbsolute, cpu.inc)
	cpu.instrTable[INC_ZP_X] = newIntr("INC_ZP_X", cpu.addrZeroPageX, cpu.inc)
	cpu.instrTable[INC_ABS_X] = newIntr("INC_ABS_X", cpu.addrAbsoluteX, cpu.inc)

	// DEC
	cpu.instrTable[DEC_ZP] = newIntr("DEC_ZP", cpu.addrZeroPage, cpu.dec)
	cpu.instrTable[DEC_ABS] = newIntr("DEC_ABS", cpu.addrAbsolute, cpu.dec)
	cpu.instrTable[DEC_ZP_X] = newIntr("DEC_ZP_X", cpu.addrZeroPageX, cpu.dec)
	cpu.instrTable[DEC_ABS_X] = newIntr("DEC_ABS_X", cpu.addrAbsoluteX, cpu.dec)

	// INC(X/Y)
	cpu.instrTable[INX] = newIntr("INX", cpu.addrImplied, cpu.inx)
	cpu.instrTable[INY] = newIntr("INY", cpu.addrImplied, cpu.iny)

	// DEC(X/Y)
	cpu.instrTable[DEX] = newIntr("DEX", cpu.addrImplied, cpu.dex)
	cpu.instrTable[DEY] = newIntr("DEY", cpu.addrImplied, cpu.dey)

	cpu.instrTable[BNE] = newIntr("BNE", cpu.addrRelative, cpu.bne)
	cpu.instrTable[BEQ] = newIntr("BEQ", cpu.addrRelative, cpu.beq)
	cpu.instrTable[BVS] = newIntr("BVS", cpu.addrRelative, cpu.bvs)
	cpu.instrTable[BVC] = newIntr("BVC", cpu.addrRelative, cpu.bvc)
	cpu.instrTable[BCC] = newIntr("BCC", cpu.addrRelative, cpu.bcc)
	cpu.instrTable[BCS] = newIntr("BCS", cpu.addrRelative, cpu.bcs)
	cpu.instrTable[BPL] = newIntr("BPL", cpu.addrRelative, cpu.bpl)
	cpu.instrTable[BMI] = newIntr("BMI", cpu.addrRelative, cpu.bmi)

	cpu.instrTable[AND_IMM] = newIntr("AND_IMM", cpu.addrImmediate, cpu.and)
	cpu.instrTable[AND_ZP] = newIntr("AND_ZP", cpu.addrZeroPage, cpu.and)
	cpu.instrTable[AND_ZP_X] = newIntr("AND_ZP_X", cpu.addrZeroPageX, cpu.and)
	cpu.instrTable[AND_ABS] = newIntr("AND_ABS", cpu.addrAbsolute, cpu.and)
	cpu.instrTable[AND_ABS_X] = newIntr("AND_ABS_X", cpu.addrAbsoluteX, cpu.and)
	cpu.instrTable[AND_ABS_Y] = newIntr("AND_ABS_Y", cpu.addrAbsoluteY, cpu.and)
	cpu.instrTable[AND_IND_X] = newIntr("AND_IND_X", cpu.addrIndirectX, cpu.and)
	cpu.instrTable[AND_IND_Y] = newIntr("AND_IND_Y", cpu.addrIndirectY, cpu.and)

	cpu.instrTable[ORA_IMM] = newIntr("ORA_IMM", cpu.addrImmediate, cpu.ora)
	cpu.instrTable[ORA_ZP] = newIntr("ORA_ZP", cpu.addrZeroPage, cpu.ora)
	cpu.instrTable[ORA_ZP_X] = newIntr("ORA_ZP_X", cpu.addrZeroPageX, cpu.ora)
	cpu.instrTable[ORA_ABS] = newIntr("ORA_ABS", cpu.addrAbsolute, cpu.ora)
	cpu.instrTable[ORA_ABS_X] = newIntr("ORA_ABS_X", cpu.addrAbsoluteX, cpu.ora)
	cpu.instrTable[ORA_ABS_Y] = newIntr("ORA_ABS_Y", cpu.addrAbsoluteY, cpu.ora)
	cpu.instrTable[ORA_IND_X] = newIntr("ORA_IND_X", cpu.addrIndirectX, cpu.ora)
	cpu.instrTable[ORA_IND_Y] = newIntr("ORA_IND_Y", cpu.addrIndirectY, cpu.ora)

	cpu.instrTable[EOR_IMM] = newIntr("EOR_IMM", cpu.addrImmediate, cpu.eor)
	cpu.instrTable[EOR_ZP] = newIntr("EOR_ZP", cpu.addrZeroPage, cpu.eor)
	cpu.instrTable[EOR_ZP_X] = newIntr("EOR_ZP_X", cpu.addrZeroPageX, cpu.eor)
	cpu.instrTable[EOR_ABS] = newIntr("EOR_ABS", cpu.addrAbsolute, cpu.eor)
	cpu.instrTable[EOR_ABS_X] = newIntr("EOR_ABS_X", cpu.addrAbsoluteX, cpu.eor)
	cpu.instrTable[EOR_ABS_Y] = newIntr("EOR_ABS_Y", cpu.addrAbsoluteY, cpu.eor)
	cpu.instrTable[EOR_IND_X] = newIntr("EOR_IND_X", cpu.addrIndirectX, cpu.eor)
	cpu.instrTable[EOR_IND_Y] = newIntr("EOR_IND_Y", cpu.addrIndirectY, cpu.eor)

	cpu.instrTable[CMP_IMM] = newIntr("CMP_IMM", cpu.addrImmediate, cpu.cmp)
	cpu.instrTable[CMP_ZP] = newIntr("CMP_ZP", cpu.addrZeroPage, cpu.cmp)
	cpu.instrTable[CMP_ZP_X] = newIntr("CMP_ZP_X", cpu.addrZeroPageX, cpu.cmp)
	cpu.instrTable[CMP_ABS] = newIntr("CMP_ABS", cpu.addrAbsolute, cpu.cmp)
	cpu.instrTable[CMP_ABS_X] = newIntr("CMP_ABS_X", cpu.addrAbsoluteX, cpu.cmp)
	cpu.instrTable[CMP_ABS_Y] = newIntr("CMP_ABS_Y", cpu.addrAbsoluteY, cpu.cmp)
	cpu.instrTable[CMP_IND_X] = newIntr("CMP_IND_X", cpu.addrIndirectX, cpu.cmp)
	cpu.instrTable[CMP_IND_Y] = newIntr("CMP_IND_Y", cpu.addrIndirectY, cpu.cmp)

	cpu.instrTable[JMP_ABS] = newIntr("JMP_ABS", cpu.addrAbsolute, cpu.jmp)
	cpu.instrTable[RTS] = newIntr("RTS", cpu.addrImplied, cpu.rts)

	cpu.instrTable[ASL_ACC] = newIntr("ASL_ACC", cpu.addrImplied, cpu.aslAcc)
	cpu.instrTable[ASL_ZP] = newIntr("ASL_ZP", cpu.addrZeroPage, cpu.asl)
	cpu.instrTable[ASL_ZP_X] = newIntr("ASL_ZP_X", cpu.addrZeroPageX, cpu.asl)
	cpu.instrTable[ASL_ABS] = newIntr("ASL_ABS", cpu.addrAbsolute, cpu.asl)
	cpu.instrTable[ASL_ABS_X] = newIntr("ASL_ABS_X", cpu.addrAbsoluteX, cpu.asl)

	cpu.instrTable[LSR_ACC] = newIntr("LSR_ACC", cpu.addrImplied, cpu.lsrAcc)
	cpu.instrTable[LSR_ZP] = newIntr("LSR_ZP", cpu.addrZeroPage, cpu.lsr)
	cpu.instrTable[LSR_ABS] = newIntr("LSR_ABS", cpu.addrAbsolute, cpu.lsr)
	cpu.instrTable[LSR_ABS_X] = newIntr("LSR_ABS_X", cpu.addrAbsoluteX, cpu.lsr)

	cpu.instrTable[CPX_IMM] = newIntr("CPX_IMM", cpu.addrImmediate, cpu.cpx)
	cpu.instrTable[CPX_ZP] = newIntr("CPX_ZP", cpu.addrZeroPage, cpu.cpx)
	cpu.instrTable[CPY_IMM] = newIntr("CPY_IMM", cpu.addrImmediate, cpu.cpy)
	cpu.instrTable[CPY_ZP] = newIntr("CPY_ZP", cpu.addrZeroPage, cpu.cpy)

	cpu.instrTable[TXA] = newIntr("TXA", cpu.addrImplied, cpu.txa)
	cpu.instrTable[TYA] = newIntr("TYA", cpu.addrImplied, cpu.tya)
	cpu.instrTable[TAX] = newIntr("TAX", cpu.addrImplied, cpu.tax)
	cpu.instrTable[TAY] = newIntr("TAY", cpu.addrImplied, cpu.tay)

	cpu.instrTable[TXS] = newIntr("TXS", cpu.addrImplied, cpu.txs)
	cpu.instrTable[TSX] = newIntr("TSX", cpu.addrImplied, cpu.tsx)

	cpu.instrTable[CLC] = newIntr("CLC", cpu.addrImplied, cpu.clc)
	cpu.instrTable[CLV] = newIntr("CLV", cpu.addrImplied, cpu.clv)
	cpu.instrTable[CLD] = newIntr("CLD", cpu.addrImplied, cpu.cld)
	cpu.instrTable[CLI] = newIntr("CLI", cpu.addrImplied, cpu.cli)

	cpu.instrTable[SEC] = newIntr("SEC", cpu.addrImplied, cpu.sec)
	cpu.instrTable[SEI] = newIntr("SEI", cpu.addrImplied, cpu.sei)

	cpu.instrTable[PHP] = newIntr("PHP", cpu.addrImplied, cpu.php)
	cpu.instrTable[PHA] = newIntr("PHA", cpu.addrImplied, cpu.pha)
	cpu.instrTable[PLA] = newIntr("PLA", cpu.addrImplied, cpu.pla)
	cpu.instrTable[PLP] = newIntr("PLP", cpu.addrImplied, cpu.plp)

	cpu.instrTable[ROL_ACC] = newIntr("ROL_ACC", cpu.addrImplied, cpu.rolAcc)
	cpu.instrTable[ROL_ZP] = newIntr("ROL_ZP", cpu.addrZeroPage, cpu.rol)
	cpu.instrTable[ROL_ABS] = newIntr("ROL_ABS", cpu.addrAbsolute, cpu.rol)
	cpu.instrTable[ROL_ABS_X] = newIntr("ROL_ABS_X", cpu.addrAbsoluteX, cpu.rol)

	cpu.instrTable[ROR_ACC] = newIntr("ROR_ACC", cpu.addrImplied, cpu.rorAcc)
	cpu.instrTable[ROR_ZP] = newIntr("ROR_ZP", cpu.addrZeroPage, cpu.ror)
	cpu.instrTable[ROR_ABS] = newIntr("ROR_ABS", cpu.addrAbsolute, cpu.ror)
	cpu.instrTable[ROR_ABS_X] = newIntr("ROR_ABS_X", cpu.addrAbsoluteX, cpu.ror)

	cpu.instrTable[STX_ZP] = newIntr("STX_ZP", cpu.addrZeroPage, cpu.stx)
	cpu.instrTable[STX_ZP_Y] = newIntr("STX_ZP_Y", cpu.addrZeroPageY, cpu.stx)
	cpu.instrTable[STX_ABS] = newIntr("STX_ABS", cpu.addrAbsolute, cpu.stx)

	cpu.instrTable[STY_ZP] = newIntr("STY_ZP", cpu.addrZeroPage, cpu.sty)
	cpu.instrTable[STY_ZP_X] = newIntr("STY_ZP_X", cpu.addrZeroPageX, cpu.sty)
	cpu.instrTable[STY_ABS] = newIntr("STY_ABS", cpu.addrAbsolute, cpu.sty)

	cpu.instrTable[BIT_ZP] = newIntr("BIT_ZP", cpu.addrZeroPage, cpu.bit)
	cpu.instrTable[BIT_ABS] = newIntr("BIT_ABS", cpu.addrAbsolute, cpu.bit)

	cpu.instrTable[SBC_IMM] = newIntr("SBC_IMM", cpu.addrImmediate, cpu.sbc)
	cpu.instrTable[SBC_ZP] = newIntr("SBC_ZP", cpu.addrZeroPage, cpu.sbc)
	cpu.instrTable[SBC_ABS] = newIntr("SBC_ABS", cpu.addrAbsolute, cpu.sbc)
	cpu.instrTable[SBC_ZP_X] = newIntr("SBC_ZP_X", cpu.addrZeroPageX, cpu.sbc)
	cpu.instrTable[SBC_ABS_X] = newIntr("SBC_ABS_X", cpu.addrAbsoluteX, cpu.sbc)
	cpu.instrTable[SBC_ABS_Y] = newIntr("SBC_ABS_Y", cpu.addrAbsoluteY, cpu.sbc)
	cpu.instrTable[SBC_IND_X] = newIntr("SBC_IND_X", cpu.addrIndirectX, cpu.sbc)
	cpu.instrTable[SBC_IND_Y] = newIntr("SBC_IND_Y", cpu.addrIndirectY, cpu.sbc)

	cpu.instrTable[RTI] = newIntr("RTI", cpu.addrImplied, cpu.rti)

	cpu.instrTable[SED] = newIntr("SED", cpu.addrImplied, cpu.sed)

	cpu.instrTable[NOP] = newIntr("NOP", cpu.addrImplied, cpu.nop)

	// NON DOC COMMANDS
	cpu.instrTable[ISB_ZP] = newIntr("ISB_ZP", cpu.addrZeroPage, cpu.isb)
	cpu.instrTable[ISB_ZP_X] = newIntr("ISB_ZP_X", cpu.addrZeroPageX, cpu.isb)
	cpu.instrTable[ISB_ABS] = newIntr("ISB_ABS", cpu.addrAbsolute, cpu.isb)
	cpu.instrTable[ISB_ABS_X] = newIntr("ISB_ABS_X", cpu.addrAbsoluteX, cpu.isb)
	cpu.instrTable[ISB_ABS_Y] = newIntr("ISB_ABS_Y", cpu.addrAbsoluteY, cpu.isb)
	cpu.instrTable[ISB_IND_X] = newIntr("ISB_IND_X", cpu.addrIndirectX, cpu.isb)
	cpu.instrTable[ISB_IND_Y] = newIntr("ISB_IND_Y", cpu.addrIndirectY, cpu.isb)
}
