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
	switch mode {
	case Imm:
		fmt.Println("MODE: IMM")
		v = cpu.readImm(cpu.pc + 1)
	case Zp:
		fmt.Println("MODE: ZP")
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
		fmt.Println("INVALID ADDRESSING MODE!")
	}
	return v
}

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

}
