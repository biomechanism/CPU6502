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

func (cpu *Cpu) subWithBorrow(val1, val2 byte) byte {

	result := val1 - val2

	cpu.setBorrowStatus(val1, val2)
	cpu.setNegativeStatus(result)
	cpu.setZeroStatus(result)

	if result > val1 {
		cpu.SetCarry()
	} else {
		cpu.ClearCarry()
	}

	cpu.setOverflowStatus(val1, val2, result)
	return result
}

func (cpu *Cpu) readOpAddr(loc uint16) uint16 {
	opcode := cpu.mem[loc]
	mode := infoArray[opcode][AddressMode]
	switch mode {
	case Abs:
		return cpu.readOpAddrAbs(cpu.pc + 1)
	case Ind:
		return cpu.readOpAddrInd(cpu.pc + 1)
	default:
		fmt.Println("INVALID ADDRESSING MODE! (Read)")
		return 0
	}
}

func (cpu *Cpu) readOpValue(loc uint16) (byte, int) {
	opcode := cpu.mem[loc]
	mode := infoArray[opcode][AddressMode]
	fmt.Printf("OP: %x, AM: %d, LOC: %v\n", opcode, mode, loc)
	var v byte
	var c int
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
		v, c = cpu.readAbsX(cpu.pc + 1)
	case AbsY:
		fmt.Println("MODE: ABSY")
		v, c = cpu.readAbsY(cpu.pc + 1)
	case IndX:
		fmt.Println("MODE: INDX")
		v = cpu.readIndX(cpu.pc + 1)
	case IndY:
		fmt.Println("MODE: INDY")
		v = cpu.readIndY(cpu.pc + 1)
	default:
		fmt.Println("INVALID ADDRESSING MODE! (Read)")
	}
	return v, c
}

func (cpu *Cpu) writeOpValue(opcodeLoc uint16, value byte) int {
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
		return cpu.writeAbsX(cpu.pc+1, value)
	case AbsY:
		//TODO: Needs Tests
		return cpu.writeAbsY(cpu.pc+1, value)
	case IndX:
		//TODO: Needs Tests
		//return cpu.writeIndX(cpu.pc+1, value)
	default:
		fmt.Println("INVALID ADDRESSING MODE! (Write)")
	}

	return 0
}

// func (cpu *Cpu) pageCrossed() {

// }

func (cpu *Cpu) ADC() (bool, int) {
	fmt.Printf("IN ADC, pc = %v - OP: %v\n", cpu.pc, cpu.mem[cpu.pc])
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	fmt.Println("READ ADC CYCLES")
	v, c := cpu.readOpValue(cpu.pc)
	fmt.Printf(">>OPVAL: %v\n", v)
	cpu.a = cpu.addWithCarry(cpu.addWithCarry(v, cpu.getCarry()), cpu.a)
	cpu.setNegativeStatus(cpu.a)

	return false, cycles + c
}

func (cpu *Cpu) AND() (bool, int) {
	v, c := cpu.readOpValue(cpu.pc)
	cpu.a &= v
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	return false, cycles + c
}

func (cpu *Cpu) ASL() (bool, int) {

	v, c := cpu.readOpValue(cpu.pc)
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
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]

	return false, cycles + c
}

//Need tests for Cycles
//FIXME: What about the auto increment after execution?
func (cpu *Cpu) BCC() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if !cpu.isCarry() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BCC Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		c := boundaryCycles(uint16(relAddr), newAddr)
		cpu.pc = newAddr
		return true, cycles + c + 1
	}
	fmt.Print("BCC Fall Through\n")
	return false, cycles
}

//Need tests for Cycles
func (cpu *Cpu) BCS() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if cpu.isCarry() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BCS Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		cpu.pc = newAddr
		c := boundaryCycles(uint16(relAddr), newAddr)
		return true, cycles + c + 1
	}
	fmt.Print("BCS Fall Through\n")
	return false, cycles
}

func (cpu *Cpu) BEQ() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if cpu.isZero() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BEQ Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		c := boundaryCycles(uint16(relAddr), newAddr)
		cpu.pc = newAddr
		return true, cycles + c + 1
	}
	fmt.Print("BEQ Fall Through\n")
	return false, cycles

}

//TODO
func (cpu *Cpu) BIT() (bool, int) {
	return false, 0
}

func (cpu *Cpu) BMI() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if cpu.isNegative() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BMI Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		c := boundaryCycles(uint16(relAddr), newAddr)
		cpu.pc = newAddr
		return true, cycles + c + 1
	}
	fmt.Print("BMI Fall Through\n")
	return false, cycles
}

func (cpu *Cpu) BNE() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if !cpu.isZero() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BNE Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		c := boundaryCycles(uint16(relAddr), newAddr)
		cpu.pc = newAddr
		return true, cycles + c + 1
	}
	fmt.Print("BNE Fall Through\n")
	return false, cycles
}

func (cpu *Cpu) BPL() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if !cpu.isNegative() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BPL Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		c := boundaryCycles(uint16(relAddr), newAddr)
		cpu.pc = newAddr
		return true, cycles + c + 1
	}
	fmt.Print("BPL Fall Through\n")
	return false, cycles
}

//TODO: Need to check more details in the b status bit.
//My understanding of the BRK command is, upon BRK execution
//the program counter + 2 is pushed to the stack, one byte at a
//time, first high byte and then the low byte. After this the
//program status register is pushed to the stack. Once this is done
//the B flag on the status register is set to indicate this interrupt
//is the result of the BRK instruction and not another interrupt.
//PC should then be set to $FFFE/$FFFF, the BRK interrupt vector address.
//
//NOTE: Not clear whether the B flag should be set and pushed or pushed and
//then set. Although it sounds like it is only set on the stacks copy of
//the status register.
func (cpu *Cpu) BRK() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	pc := cpu.pc + 2
	pch := uint8((pc & 0xFF00) >> 8)
	pcl := uint8(pc & 0x00FF)
	cpu.push(byte(pch))
	cpu.push(byte(pcl))
	cpu.b = true
	cpu.pushStatusToStack()
	cpu.b = false

	pcl = cpu.mem[0xFFFE]
	pch = cpu.mem[0xFFFF]
	//fmt.Printf("PCL %v PCH %v\n", pcl, pch)
	cpu.pc = (uint16(pch) << 8) | uint16(pcl)
	//cpu.pc = 0xFFFE
	return true, cycles
}

func (cpu *Cpu) BVC() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if !cpu.isOverflow() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BVC Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		c := boundaryCycles(uint16(relAddr), newAddr)
		cpu.pc = newAddr
		return true, cycles + c + 1
	}
	fmt.Print("BPL Fall Through\n")
	return false, cycles
}

func (cpu *Cpu) BVS() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	if cpu.isOverflow() {
		relAddr := int8(cpu.mem[cpu.pc+1])
		fmt.Printf("BVS Branching; Rel val: %d\n", relAddr)
		newAddr := cpu.pc + uint16(relAddr)
		c := boundaryCycles(uint16(relAddr), newAddr)
		cpu.pc = newAddr
		return true, cycles + c + 1
	}
	fmt.Print("BPS Fall Through\n")
	return false, cycles
}

//No tests yet
func (cpu *Cpu) CLC() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.ClearCarry()
	cpu.pc++
	return true, cycles
}

//No tests yet
func (cpu *Cpu) CLD() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.ClearDecimalMode()
	cpu.pc++
	return true, cycles
}

//No tests yet
func (cpu *Cpu) CLI() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.ClearIRQDisable()
	cpu.pc++
	return true, cycles
}

//No tests yet
func (cpu *Cpu) CLV() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.clearOverflowStatus()
	cpu.pc++
	return true, cycles
}

func (cpu *Cpu) CMP() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	acc := cpu.a
	mem, c := cpu.readOpValue(cpu.pc)
	result := acc - mem
	cpu.setNegativeStatus(result)
	cpu.setBorrowStatus(acc, mem)
	cpu.setZeroStatus(result)
	return false, cycles + c
}

func (cpu *Cpu) CPX() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	x := cpu.x
	mem, _ := cpu.readOpValue(cpu.pc)
	result := x - mem
	cpu.setNegativeStatus(result)
	cpu.setBorrowStatus(x, mem)
	cpu.setZeroStatus(result)
	return false, cycles
}

func (cpu *Cpu) CPY() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	y := cpu.y
	mem, _ := cpu.readOpValue(cpu.pc)
	result := y - mem
	cpu.setNegativeStatus(result)
	cpu.setBorrowStatus(y, mem)
	cpu.setZeroStatus(result)
	return false, cycles
}

func (cpu *Cpu) DEC() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, _ := cpu.readOpValue(cpu.pc)
	val--
	cpu.writeOpValue(cpu.pc, val)
	cpu.setNegativeStatus(val)
	cpu.setZeroStatus(val)
	return false, cycles
}

func (cpu *Cpu) DEX() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.x--
	cpu.setNegativeStatus(cpu.x)
	cpu.setZeroStatus(cpu.x)
	return false, cycles
}

func (cpu *Cpu) DEY() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.y--
	cpu.setNegativeStatus(cpu.y)
	cpu.setZeroStatus(cpu.y)
	return false, cycles
}

func (cpu *Cpu) EOR() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, c := cpu.readOpValue(cpu.pc)
	cpu.a ^= val
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
	return false, cycles + c
}

func (cpu *Cpu) INC() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, _ := cpu.readOpValue(cpu.pc)
	val++
	cpu.writeOpValue(cpu.pc, val)
	cpu.setNegativeStatus(val)
	cpu.setZeroStatus(val)
	return false, cycles
}

func (cpu *Cpu) INX() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.x++
	cpu.setNegativeStatus(cpu.x)
	cpu.setZeroStatus(cpu.x)
	return false, cycles
}

func (cpu *Cpu) INY() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.y++
	cpu.setNegativeStatus(cpu.y)
	cpu.setZeroStatus(cpu.y)
	return false, cycles
}

func (cpu *Cpu) JMP() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	addr := cpu.readOpAddr(cpu.pc)
	cpu.pc = addr
	return true, cycles
}

func (cpu *Cpu) JSR() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	addr := cpu.readOpAddr(cpu.pc)
	returnAddr := cpu.pc + 2
	pcl := byte(returnAddr & 0x00FF)
	pch := byte(returnAddr >> 8)
	cpu.push(pch)
	cpu.push(pcl)
	cpu.pc = addr
	return true, cycles
}

func (cpu *Cpu) LDA() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, c := cpu.readOpValue(cpu.pc)
	cpu.a = val
	cpu.setNegativeStatus(val)
	cpu.setZeroStatus(val)
	return false, cycles + c
}

func (cpu *Cpu) LDX() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, c := cpu.readOpValue(cpu.pc)
	cpu.x = val
	cpu.setNegativeStatus(val)
	cpu.setZeroStatus(val)
	return false, cycles + c
}

func (cpu *Cpu) LDY() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, c := cpu.readOpValue(cpu.pc)
	cpu.y = val
	cpu.setNegativeStatus(val)
	cpu.setZeroStatus(val)
	return false, cycles + c
}

func (cpu *Cpu) LSR() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, _ := cpu.readOpValue(cpu.pc)
	cpu.c = val&0x01 > 0
	val >>= 1
	cpu.writeOpValue(cpu.pc, val)
	return false, cycles
}

func (cpu *Cpu) NOP() (bool, int) {
	fmt.Printf("IN NOP, pc = %v\n", cpu.pc)
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	return false, cycles
}

func (cpu *Cpu) ORA() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, c := cpu.readOpValue(cpu.pc)
	cpu.a |= val
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
	return false, cycles + c
}

func (cpu *Cpu) PHA() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.push(cpu.a)
	return false, cycles
}

func (cpu *Cpu) PHP() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.pushStatusToStack()
	return false, cycles
}

func (cpu *Cpu) PLA() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.a = cpu.pop()
	return false, cycles
}

func (cpu *Cpu) PLP() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.popStatusFromStack()
	return false, cycles
}

func (cpu *Cpu) ROL() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, _ := cpu.readOpValue(cpu.pc)
	fromCarry := cpu.b2i(cpu.c)
	toCarry := val & 0x80
	newVal := val << 1
	cpu.c = cpu.i2b(toCarry)
	newVal |= fromCarry
	cpu.setNegativeStatus(newVal)
	cpu.setZeroStatus(newVal)
	cpu.writeOpValue(cpu.pc, newVal)
	return false, cycles
}

func (cpu *Cpu) ROR() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	val, _ := cpu.readOpValue(cpu.pc)
	fromCarry := cpu.b2i(cpu.c)
	toCarry := val & 0x01
	newVal := val >> 1
	cpu.c = cpu.i2b(toCarry)
	newVal |= (fromCarry << 7)
	cpu.setNegativeStatus(newVal)
	cpu.setZeroStatus(newVal)
	cpu.writeOpValue(cpu.pc, newVal)
	return false, cycles
}

func (cpu *Cpu) RTI() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.popStatusFromStack()
	cpu.b = false

	pcl := cpu.pop()
	pch := cpu.pop()

	fmt.Printf("RTI: PCL %v, PCH %v\n", pcl, pch)

	cpu.pc = (uint16(pch) << 8) | uint16(pcl)

	return true, cycles
}

func (cpu *Cpu) RTS() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	pcl := cpu.pop()
	pch := cpu.pop()
	cpu.pc = ((uint16(pch) << 8) | uint16(pcl)) + 1
	return true, cycles
}

//Need more tests
func (cpu *Cpu) SBC() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	v, c := cpu.readOpValue(cpu.pc)
	cpu.a = cpu.subWithBorrow(cpu.a, cpu.subWithBorrow(v, cpu.getCarry()))
	return false, cycles + c
}

func (cpu *Cpu) SEC() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.SetCarry()
	return false, cycles
}

func (cpu *Cpu) SED() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.d = true
	return false, cycles
}

func (cpu *Cpu) SEI() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.i = true
	return false, cycles
}

//Needs tests
func (cpu *Cpu) STA() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.writeOpValue(cpu.pc, cpu.a)
	return false, cycles
}

//Needs tests
func (cpu *Cpu) STX() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.writeOpValue(cpu.pc, cpu.x)
	return false, cycles
}

//Needs tests
func (cpu *Cpu) STY() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.writeOpValue(cpu.pc, cpu.y)
	return false, cycles
}

//Needs tests
func (cpu *Cpu) TAX() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.x = cpu.a
	cpu.setNegativeStatus(cpu.x)
	cpu.setZeroStatus(cpu.x)
	return false, cycles
}

//Needs tests
func (cpu *Cpu) TAY() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.y = cpu.a
	cpu.setNegativeStatus(cpu.y)
	cpu.setZeroStatus(cpu.y)
	return false, cycles
}

//Needs tests
func (cpu *Cpu) TSX() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.x = cpu.s
	cpu.setNegativeStatus(cpu.x)
	cpu.setZeroStatus(cpu.x)
	return false, cycles
}

//Needs tests
func (cpu *Cpu) TXA() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.a = cpu.x
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
	return false, cycles
}

//Needs tests
func (cpu *Cpu) TXS() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.s = cpu.x
	return false, cycles
}

//Needs tests
func (cpu *Cpu) TYA() (bool, int) {
	opcode := cpu.mem[cpu.pc]
	cycles := infoArray[opcode][Cycles]
	cpu.a = cpu.y
	cpu.setNegativeStatus(cpu.a)
	cpu.setZeroStatus(cpu.a)
	return false, cycles
}
