package core

func (cpu *Cpu) Decode() func() int {
	opcode := cpu.mem[cpu.pc]
	switch opcode {
	case adcAbs, adcAbsX, adcAbsY, adcImm, adcIndX, adcIndY, adcZp, adcZpX:
		return func() int {
			return executor(cpu.ADC, cpu)
		}
	case andAbs, andAbsX, andAbsY, andImm, andIndX, andIndY, andZp, andZpX:
		return func() int {
			return executor(cpu.AND, cpu)
		}
	case aslAbs, aslAbsX, aslAcc, aslZp, aslZpX:
		return func() int {
			return executor(cpu.ASL, cpu)
		}
	case bcc:
		return func() int {
			return executor(cpu.BCC, cpu)
		}
	case bcs:
		return func() int {
			return executor(cpu.BCS, cpu)
		}
	case beq:
		return func() int {
			return executor(cpu.BEQ, cpu)
		}
	case bitAbs, bitZp:
		return func() int {
			return executor(cpu.BIT, cpu)
		}
	case bmi:
		return func() int {
			return executor(cpu.BMI, cpu)
		}
	case bne:
		return func() int {
			return executor(cpu.BNE, cpu)
		}
	case bpl:
		return func() int {
			return executor(cpu.BPL, cpu)
		}
	case brk:
		return func() int {
			return executor(cpu.BRK, cpu)
		}
	case bvc:
		return func() int {
			return executor(cpu.BVC, cpu)
		}
	case bvs:
		return func() int {
			return executor(cpu.BVS, cpu)
		}
	case clc:
		return func() int {
			return executor(cpu.CLC, cpu)
		}
	case cld:
		return func() int {
			return executor(cpu.CLD, cpu)
		}
	case cli:
		return func() int {
			return executor(cpu.CLI, cpu)
		}
	case clv:
		return func() int {
			return executor(cpu.CLV, cpu)
		}
	case cmpAbs, cmpAbsX, cmpAbsY, cmpImm, cmpIndX, cmpIndY, cmpZp, cmpZpX:
		return func() int {
			return executor(cpu.CMP, cpu)
		}
	case cpxAbs, cpxImm, cpxZp:
		return func() int {
			return executor(cpu.CPX, cpu)
		}
	case cpyAbs, cpyImm, cpyZp:
		return func() int {
			return executor(cpu.CPY, cpu)
		}
	case decAbs, decAbsX, decZp, decZpX:
		return func() int {
			return executor(cpu.DEC, cpu)
		}
	case dex:
		return func() int {
			return executor(cpu.DEX, cpu)
		}
	case dey:
		return func() int {
			return executor(cpu.DEY, cpu)
		}
	case eorAbs, eorAbsX, eorAbsY, eorImm, eorIndX, eorIndY, eorZp, eorZpX:
		return func() int {
			return executor(cpu.EOR, cpu)
		}
	case incAbs, incAbsX, incZp, incZpX:
		return func() int {
			return executor(cpu.INC, cpu)
		}
	case inx:
		return func() int {
			return executor(cpu.INX, cpu)
		}
	case iny:
		return func() int {
			return executor(cpu.INY, cpu)
		}
	case jmpAbs, jmpInd:
		return func() int {
			return executor(cpu.JMP, cpu)
		}
	case jsr:
		return func() int {
			return executor(cpu.JSR, cpu)
		}
	case ldaAbs, ldaAbsX, ldaAbsY, ldaImm, ldaIndX, ldaIndY, ldaZp, ldaZpX:
		return func() int {
			return executor(cpu.LDA, cpu)
		}
	case ldxAbs, ldxAbsY, ldxImm, ldxZp, ldxZpY:
		return func() int {
			return executor(cpu.LDX, cpu)
		}
	case ldyAbs, ldyAbsX, ldyImm, ldyZp, ldyZpX:
		return func() int {
			return executor(cpu.LDY, cpu)
		}
	case lsrAbs, lsrAbsX, lsrAcc, lsrZp, lsrZpX:
		return func() int {
			return executor(cpu.LSR, cpu)
		}
	case nop:
		return func() int {
			return executor(cpu.NOP, cpu)
		}
	case oraAbs, oraAbsX, oraAbsY, oraImm, oraIndX, oraIndY, oraZp, oraZpX:
		return func() int {
			return executor(cpu.ORA, cpu)
		}
	case pha:
		return func() int {
			return executor(cpu.PHA, cpu)
		}
	case php:
		return func() int {
			return executor(cpu.PHP, cpu)
		}
	case pla:
		return func() int {
			return executor(cpu.PLA, cpu)
		}
	case plp:
		return func() int {
			return executor(cpu.PLP, cpu)
		}
	case rolAbs, rolAbsX, rolAcc, rolZp, rolZpX:
		return func() int {
			return executor(cpu.ROL, cpu)
		}
	case rorAbs, rorAbsX, rorAcc, rorZp, rorZpX:
		return func() int {
			return executor(cpu.ROR, cpu)
		}
	case rti:
		return func() int {
			return executor(cpu.RTI, cpu)
		}
	case rts:
		return func() int {
			return executor(cpu.RTS, cpu)
		}
	case sbcAbs, sbcAbsX, sbcAbsY, sbcImm, sbcIndX, sbcIndY, sbcZp, sbcZpX:
		return func() int {
			return executor(cpu.SBC, cpu)
		}
	case sec:
		return func() int {
			return executor(cpu.SEC, cpu)
		}
	case sed:
		return func() int {
			return executor(cpu.SED, cpu)
		}
	case sei:
		return func() int {
			return executor(cpu.SEI, cpu)
		}
	case staAbs, staAbsX, staAbsY, staIndX, staIndY, staZp, staZpX:
		return func() int {
			return executor(cpu.STA, cpu)
		}
	case stxAbs, stxZp, stxZpY:
		return func() int {
			return executor(cpu.STX, cpu)
		}
	case styAbs, styZp, styZpX:
		return func() int {
			return executor(cpu.STY, cpu)
		}
	case tax:
		return func() int {
			return executor(cpu.TAX, cpu)
		}
	case tay:
		return func() int {
			return executor(cpu.TAY, cpu)
		}
	case tsx:
		return func() int {
			return executor(cpu.TSX, cpu)
		}
	case txa:
		return func() int {
			return executor(cpu.TXA, cpu)
		}
	case txs:
		return func() int {
			return executor(cpu.TXS, cpu)
		}
	case tya:
		return func() int {
			return executor(cpu.TYA, cpu)
		}
	default:
		//Replace with proper function
		return nil
	}

}

func executor(fn func() (bool, int), cpu *Cpu) int {
	opCode := cpu.mem[cpu.pc]
	opSize := infoArray[opCode][Size]
	b, c := fn()
	if b == false {
		cpu.pc += uint16(opSize)
	}

	return c
}
