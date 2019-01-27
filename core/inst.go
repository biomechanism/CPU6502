package core

import "fmt"

func (cpu *Cpu) addWithCarry(val1, val2 byte) byte {
	result := val1 + val2

	if result < val1 {
		cpu.p |= 1
	} else {
		cpu.p &= ^byte(1)
	}

	cpu.setOverflowStatus(val1, val2, result)

	return result
}

func (cpu *Cpu) readOpValue(loc uint16) byte {
	opcode := cpu.mem[loc]
	mode := infoArray[opcode][AddressMode]
	fmt.Printf("OP: %x, AM: %d\n", opcode, mode)
	var v byte
	//var memAdder uint16
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
	default:
		fmt.Println("INVALID ADDRESSING MODE! (Write)")
	}
}

// func (cpu *Cpu) writeOpValue(mode int, value byte) {
// 	// opcode := inst
// 	// mode := infoArray[opcode][AddressMode]
// 	switch mode {
// 	case Acc:
// 		cpu.a = value
// 	default:
// 		fmt.Println("INVALID ADDRESSING MODE!")
// 	}

// }

func (cpu *Cpu) ADC() {
	v := cpu.readOpValue(cpu.pc)
	cpu.a = cpu.addWithCarry(cpu.addWithCarry(v, cpu.getCarry()), cpu.a)
	cpu.setNegativeStatus(cpu.a)
}

func (cpu *Cpu) AND() {
	v := cpu.readOpValue(cpu.pc)
	cpu.a &= v
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
}

func (cpu *Cpu) ASL() {

	v := cpu.readOpValue(cpu.pc)
	fmt.Printf("ASL Value before Op: %d\n", v)

	var isCarry bool = false

	if (0x80 & v) > 0 {
		isCarry = true

	} else {
		isCarry = false

	}

	v <<= 1

	cpu.writeOpValue(cpu.pc, v)

	if isCarry {
		fmt.Println("ASL: Setting Carry Flag")
		cpu.p |= 1
	} else {
		fmt.Println("ASL: Clearing Carry Flag")
		cpu.p &= ^byte(1)
	}

	cpu.setNegativeStatus(v)
	cpu.setZeroStatus(v)
}

func (cpu *Cpu) BCC() {

}

func (cpu *Cpu) BCS() {

}

func (cpu *Cpu) BEQ() {

}

func (cpu *Cpu) BIT() {

}

func (cpu *Cpu) BMI() {

}

func (cpu *Cpu) BNE() {

}

func (cpu *Cpu) BPL() {

}

func (cpu *Cpu) BRK() {

}

func (cpu *Cpu) BVC() {

}

func (cpu *Cpu) BVS() {

}

func (cpu *Cpu) CLC() {

}

func (cpu *Cpu) CLD() {

}

func (cpu *Cpu) CLI() {

}

func (cpu *Cpu) CLV() {

}

func (cpu *Cpu) CMP() {

}

func (cpu *Cpu) CPX() {

}

func (cpu *Cpu) CPY() {

}

func (cpu *Cpu) DEC() {

}

func (cpu *Cpu) DEX() {

}

func (cpu *Cpu) DEY() {

}

func (cpu *Cpu) EOR() {

}

func (cpu *Cpu) INC() {

}

func (cpu *Cpu) INX() {

}

func (cpu *Cpu) INY() {

}

func (cpu *Cpu) JMP() {

}

func (cpu *Cpu) JSR() {

}

func (cpu *Cpu) LDA() {

}

func (cpu *Cpu) LDX() {

}

func (cpu *Cpu) LDY() {

}

func (cpu *Cpu) LSR() {

}

func (cpu *Cpu) NOP() {

}

func (cpu *Cpu) ORA() {

}

func (cpu *Cpu) PHA() {

}

func (cpu *Cpu) PHP() {

}

func (cpu *Cpu) PLA() {

}

func (cpu *Cpu) PLP() {

}

func (cpu *Cpu) ROL() {

}

func (cpu *Cpu) ROR() {

}

func (cpu *Cpu) RTI() {

}

func (cpu *Cpu) RTS() {

}

func (cpu *Cpu) SBC() {

}

func (cpu *Cpu) SEC() {

}

func (cpu *Cpu) SED() {

}

func (cpu *Cpu) SEI() {

}

func (cpu *Cpu) STA() {

}

func (cpu *Cpu) STX() {

}

func (cpu *Cpu) STY() {

}

func (cpu *Cpu) TAX() {

}

func (cpu *Cpu) TAY() {

}

func (cpu *Cpu) TSX() {

}

func (cpu *Cpu) TXA() {

}

func (cpu *Cpu) TXS() {

}

func (cpu *Cpu) TYA() {

}
