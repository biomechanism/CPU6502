package core

func (cpu *Cpu) Decode() func() {
	opcode := cpu.mem[cpu.pc]
	switch opcode {
	case adcAbs, adcAbsX, adcAbsY, adcImm, adcIndX, adcIndY, adcZp, adcZpX:
		return func() {
			executor(cpu.ADC, cpu)
		}
	case andAbs, andAbsX, andAbsY, andImm, andIndX, andIndY, andZp, andZpX:
		return func() {
			executor(cpu.AND, cpu)
		}
	case aslAbs, aslAbsX, aslAcc, aslZp, aslZpX:
		return func() {
			executor(cpu.ASL, cpu)
		}
	case bcc:
		return func() {
			executor(cpu.BCC, cpu)
		}
	case bcs:
		return func() {
			executor(cpu.BCS, cpu)
		}
	case beq:
		return func() {
			executor(cpu.BEQ, cpu)
		}
	case bitAbs, bitZp:
		return func() {
			executor(cpu.BIT, cpu)
		}
	case bmi:
		return func() {
			executor(cpu.BMI, cpu)
		}
	case bne:
		return func() {
			executor(cpu.BNE, cpu)
		}
	case bpl:
		return func() {
			executor(cpu.BPL, cpu)
		}
	case brk:
		return func() {
			executor(cpu.BRK, cpu)
		}
	case bvc:
		return func() {
			executor(cpu.BVC, cpu)
		}
	case bvs:
		return func() {
			executor(cpu.BVS, cpu)
		}
	case clc:
		return func() {
			executor(cpu.CLC, cpu)
		}
	case cld:
		return func() {
			executor(cpu.CLD, cpu)
		}
	case cli:
		return func() {
			executor(cpu.CLI, cpu)
		}
	case clv:
		return func() {
			executor(cpu.CLV, cpu)
		}
	case cmpAbs, cmpAbsX, cmpAbsY, cmpImm, cmpIndX, cmpIndY, cmpZp, cmpZpX:
		return func() {
			executor(cpu.CMP, cpu)
		}
	case cpxAbs, cpxImm, cpxZp:
		return func() {
			executor(cpu.CPX, cpu)
		}
	case cpyAbs, cpyImm, cpyZp:
		return func() {
			executor(cpu.CPY, cpu)
		}
	case decAbs, decAbsX, decZp, decZpX:
		return func() {
			executor(cpu.DEC, cpu)
		}
	case dex:
		return func() {
			executor(cpu.DEX, cpu)
		}
	case dey:
		return func() {
			executor(cpu.DEY, cpu)
		}
	case eorAbs, eorAbsX, eorAbsY, eorImm, eorIndX, eorIndY, eorZp, eorZpX:
		return func() {
			executor(cpu.EOR, cpu)
		}
	case incAbs, incAbsX, incZp, incZpX:
		return func() {
			executor(cpu.INC, cpu)
		}
	case inx:
		return func() {
			executor(cpu.INX, cpu)
		}
	case iny:
		return func() {
			executor(cpu.INY, cpu)
		}
	case jmpAbs, jmpInd:
		return func() {
			executor(cpu.JMP, cpu)
		}
	case jsr:
		return func() {
			executor(cpu.JSR, cpu)
		}
	case ldaAbs, ldaAbsX, ldaAbsY, ldaImm, ldaIndX, ldaIndY, ldaZp, ldaZpX:
		return func() {
			executor(cpu.LDA, cpu)
		}
	case ldxAbs, ldxAbsY, ldxImm, ldxZp, ldxZpY:
		return func() {
			executor(cpu.LDX, cpu)
		}
	case ldyAbs, ldyAbsX, ldyImm, ldyZp, ldyZpX:
		return func() {
			executor(cpu.LDY, cpu)
		}
	case lsrAbs, lsrAbsX, lsrAcc, lsrZp, lsrZpX:
		return func() {
			executor(cpu.LSR, cpu)
		}
	case nop:
		return func() {
			executor(cpu.NOP, cpu)
		}
	case oraAbs, oraAbsX, oraAbsY, oraImm, oraIndX, oraIndY, oraZp, oraZpX:
		return func() {
			executor(cpu.ORA, cpu)
		}
	case pha:
		return func() {
			executor(cpu.PHA, cpu)
		}
	case php:
		return func() {
			executor(cpu.PHP, cpu)
		}
	case pla:
		return func() {
			executor(cpu.PLP, cpu)
		}
	case plp:
		return func() {
			executor(cpu.PLP, cpu)
		}
	case rolAbs, rolAbsX, rolAcc, rolZp, rolZpX:
		return func() {
			executor(cpu.ROL, cpu)
		}
	case rorAbs, rorAbsX, rorAcc, rorZp, rorZpX:
		return func() {
			executor(cpu.ROR, cpu)
		}
	case rti:
		return func() {
			executor(cpu.RTI, cpu)
		}
	case rts:
		return func() {
			executor(cpu.RTS, cpu)
		}
	case sbcAbs, sbcAbsX, sbcAbsY, sbcImm, sbcIndX, sbcIndY, sbcZp, sbcZpX:
		return func() {
			executor(cpu.SBC, cpu)
		}
	case sec:
		return func() {
			executor(cpu.SEC, cpu)
		}
	case sed:
		return func() {
			executor(cpu.SED, cpu)
		}
	case sei:
		return func() {
			executor(cpu.SEI, cpu)
		}
	case staAbs, staAbsX, staAbsY, staIndX, staIndY, staZp, staZpX:
		return func() {
			executor(cpu.STA, cpu)
		}
	case stxAbs, stxZp, stxZpY:
		return func() {
			executor(cpu.STX, cpu)
		}
	case styAbs, styZp, styZpX:
		return func() {
			executor(cpu.STY, cpu)
		}
	case tax:
		return func() {
			executor(cpu.TAX, cpu)
		}
	case tay:
		return func() {
			executor(cpu.TAY, cpu)
		}
	case tsx:
		return func() {
			executor(cpu.TSX, cpu)
		}
	case txa:
		return func() {
			executor(cpu.TXA, cpu)
		}
	case txs:
		return func() {
			executor(cpu.TXS, cpu)
		}
	case tya:
		return func() {
			executor(cpu.TYA, cpu)
		}
	default:
		//Replace with proper function
		return nil
	}

}

func executor(fn func() bool, cpu *Cpu) {
	opCode := cpu.mem[cpu.pc]
	opSize := infoArray[opCode][Size]
	if !fn() {
		cpu.pc += uint16(opSize)
	}
}
