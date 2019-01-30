package core

import "fmt"

func (cpu *Cpu) addWithCarry(val1, val2 byte) byte {
	result := val1 + val2

	if result < val1 {
		cpu.SetCarry()
	} else {
		cpu.ClearCarry()
	}

	cpu.setOverflowStatus(val1, val2, result)

	return result
}

func (cpu *Cpu) readOpValue(loc uint16) byte {
	opcode := cpu.mem[loc]
	mode := infoArray[opcode][AddressMode]
	fmt.Printf("OP: %x, AM: %d\n", opcode, mode)
	var v byte
	switch mode {
	case Acc:
		v = cpu.a
	case Imm:
		fmt.Println("MODE: IMM")
		v = cpu.readImm(cpu.pc + 1)
	case Zp:
		fmt.Println("MODE: ZP (Read)")
		v = cpu.readZp(cpu.pc + 1)
	case ZpX:
		fmt.Println("MODE: ZPX")
		//TODO: CHECK, Not sure whether the carry needs to be handled when adding the X index
		//to the base or not.
		v = cpu.readZpX(cpu.pc + 1)
	case Abs:
		fmt.Println("MODE: ABS")
		v = cpu.readAbs(cpu.pc + 1)
	case AbsX:
		fmt.Println("MODE: ABSX")
		//TODO: CHECK, should add read calculations from and index have a carry
		//check?
		v = cpu.readAbsX(cpu.pc + 1)
	case AbsY:
		fmt.Println("MODE: ABSY")
		v = cpu.readAbsY(cpu.pc + 1)
	case IndX:
		fmt.Println("MODE: INDX")
		v = cpu.readIndX(cpu.pc + 1)
	case IndY:
		fmt.Println("MODE: INDY")
		v = cpu.readIndY(cpu.pc + 1)
	default:
		fmt.Println("INVALID ADDRESSING MODE! (Read)")
	}
	return v
}

func (cpu *Cpu) writeOpValue(opcodeLoc uint16, value byte) {
	opcode := cpu.mem[opcodeLoc]
	mode := infoArray[opcode][AddressMode]
	switch mode {
	case Acc:
		cpu.a = value
	case Zp:
		fmt.Println("MODE: ZP (Write)")
		cpu.writeZp(cpu.pc+1, value)
	case ZpX:
		fmt.Println("MODE: ZPX (Write)")
		//TODO: CHECK, Not sure whether the carry needs to be handled when adding the X index
		//to the base or not.
		cpu.writeZpX(cpu.pc+1, value)
	case Abs:
		fmt.Println("MODE: ABS")
		cpu.writeAbs(cpu.pc+1, value)
	case AbsX:
		cpu.writeAbsX(cpu.pc+1, value)
	default:
		fmt.Println("INVALID ADDRESSING MODE! (Write)")
	}
}

func (cpu *Cpu) ADC() bool {
	v := cpu.readOpValue(cpu.pc)
	cpu.a = cpu.addWithCarry(cpu.addWithCarry(v, cpu.getCarry()), cpu.a)
	cpu.setNegativeStatus(cpu.a)
	return false
}

func (cpu *Cpu) AND() bool {
	v := cpu.readOpValue(cpu.pc)
	cpu.a &= v
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
	return false
}

func (cpu *Cpu) ASL() bool {

	v := cpu.readOpValue(cpu.pc)
	fmt.Printf("ASL Value before Op: %d\n", v)

	var isCarry bool

	if (0x80 & v) > 0 {
		isCarry = true
	} else {
		isCarry = false

	}

	v <<= 1

	cpu.writeOpValue(cpu.pc, v)

	if isCarry {
		fmt.Println("ASL: Setting Carry Flag")
		cpu.SetCarry()

	} else {
		fmt.Println("ASL: Clearing Carry Flag")
		cpu.ClearCarry()
	}

	cpu.setNegativeStatus(v)
	cpu.setZeroStatus(v)

	return false
}

//FIXME: What about the auto increment after execution?
func (cpu *Cpu) BCC() bool {
	if !cpu.isCarry() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		//relAddr := cpu.mem[cpu.pc+1]
		fmt.Printf("BCC Branching; Rel val: %d\n", relAddr)
		cpu.pc += uint16(relAddr)
		return true
	}
	fmt.Print("BCC Fall Through\n")
	return false
}

func (cpu *Cpu) BCS() bool {
	if cpu.isCarry() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BCC Branching; Rel val: %d\n", relAddr)
		cpu.pc += uint16(relAddr)
		return true
	}
	fmt.Print("BCC Fall Through\n")
	return false
}

func (cpu *Cpu) BEQ() bool {
	return true
}

func (cpu *Cpu) BIT() bool {
	return true
}

func (cpu *Cpu) BMI() bool {
	return true
}

func (cpu *Cpu) BNE() bool {
	return true
}

func (cpu *Cpu) BPL() bool {
	return true
}

func (cpu *Cpu) BRK() bool {
	return false
}

func (cpu *Cpu) BVC() bool {
	return true
}

func (cpu *Cpu) BVS() bool {
	return true
}

func (cpu *Cpu) CLC() bool {
	return false
}

func (cpu *Cpu) CLD() bool {
	return false
}

func (cpu *Cpu) CLI() bool {
	return false
}

func (cpu *Cpu) CLV() bool {
	return false
}

func (cpu *Cpu) CMP() bool {
	return false
}

func (cpu *Cpu) CPX() bool {
	return false
}

func (cpu *Cpu) CPY() bool {
	return false
}

func (cpu *Cpu) DEC() bool {
	return false
}

func (cpu *Cpu) DEX() bool {
	return false
}

func (cpu *Cpu) DEY() bool {
	return false
}

func (cpu *Cpu) EOR() bool {
	return false
}

func (cpu *Cpu) INC() bool {
	return false
}

func (cpu *Cpu) INX() bool {
	return false
}

func (cpu *Cpu) INY() bool {
	return false
}

func (cpu *Cpu) JMP() bool {
	return true
}

func (cpu *Cpu) JSR() bool {
	return true
}

func (cpu *Cpu) LDA() bool {
	return false
}

func (cpu *Cpu) LDX() bool {
	return false
}

func (cpu *Cpu) LDY() bool {
	return false
}

func (cpu *Cpu) LSR() bool {
	return false
}

func (cpu *Cpu) NOP() bool {
	return false
}

func (cpu *Cpu) ORA() bool {
	return false
}

func (cpu *Cpu) PHA() bool {
	return false
}

func (cpu *Cpu) PHP() bool {
	return false
}

func (cpu *Cpu) PLA() bool {
	return false
}

func (cpu *Cpu) PLP() bool {
	return false
}

func (cpu *Cpu) ROL() bool {
	return false
}

func (cpu *Cpu) ROR() bool {
	return false
}

func (cpu *Cpu) RTI() bool {
	return true
}

func (cpu *Cpu) RTS() bool {
	return true
}

func (cpu *Cpu) SBC() bool {
	return false
}

func (cpu *Cpu) SEC() bool {
	return false
}

func (cpu *Cpu) SED() bool {
	return false
}

func (cpu *Cpu) SEI() bool {
	return false
}

func (cpu *Cpu) STA() bool {
	return false
}

func (cpu *Cpu) STX() bool {
	return false
}

func (cpu *Cpu) STY() bool {
	return false
}

func (cpu *Cpu) TAX() bool {
	return false
}

func (cpu *Cpu) TAY() bool {
	return false
}

func (cpu *Cpu) TSX() bool {
	return false
}

func (cpu *Cpu) TXA() bool {
	return false
}

func (cpu *Cpu) TXS() bool {
	return false
}

func (cpu *Cpu) TYA() bool {
	return false
}
