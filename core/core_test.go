package core

import (
	"fmt"
	"testing"
)

func init() {

}

func TestLoad(t *testing.T) {

}

func TestADCImmediate(t *testing.T) {

	cpu := newCpu()

	cpu.a = 10
	cpu.mem[0] = adcImm
	cpu.mem[1] = 4

	inst := cpu.Decode()
	inst()

	if cpu.a != 14 {
		t.Errorf("Expexted %d, Actual %d\n", 14, cpu.a)
	}

}

func TestADCImmediateOverflow(t *testing.T) {
	cpu := newCpu()

	cpu.a = 127
	cpu.mem[0] = adcImm
	cpu.mem[1] = 1

	inst := cpu.Decode()
	inst()

	if !cpu.isOverflow() {
		t.Errorf("Expexted %d, Actual %d\n", 64, cpu.p&(1<<6))
	}

}

func TestADCImmediateCarry(t *testing.T) {
	fmt.Println("ADC Immediate Carry")
	cpu := newCpu()
	cpu.a = 254
	cpu.mem[0] = adcImm
	cpu.mem[1] = 1
	cpu.SetCarry()

	inst := cpu.Decode()
	inst()

	if cpu.a != 0 {
		t.Errorf("Expexted %v, Actual %v\n", 0, cpu.a)
	}

	if !cpu.isCarry() {
		t.Errorf("Expexted %v, Actual %v\n", true, cpu.c)
	}
}

func TestADCImmediateZeroPage(t *testing.T) {
	fmt.Println("ADC Immediate Zero Page")

	cpu := newCpu()
	cpu.a = 100
	cpu.mem[0] = adcZp
	cpu.mem[1] = 4
	cpu.mem[4] = 8

	inst := cpu.Decode()
	inst()

	if cpu.a != 108 {
		t.Errorf("Expexted %d, Actual %d\n", 108, cpu.a)
	}
}

func TestADCImmediateZeroPageX(t *testing.T) {
	fmt.Println("ADC Immediate Zero Page X")

	cpu := newCpu()
	cpu.a = 11
	cpu.x = 3
	cpu.mem[0] = adcZpX
	cpu.mem[1] = 4
	cpu.mem[7] = 6

	inst := cpu.Decode()
	inst()

	if cpu.a != 17 {
		t.Errorf("Expexted %d, Actual %d\n", 17, cpu.a)
	}
}

func TestADCAbsolute(t *testing.T) {
	fmt.Println("ADC Absolute")
	cpu := newCpu()

	cpu.a = 10
	cpu.mem[0] = adcAbs
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03

	cpu.mem[1000] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 13 {
		t.Errorf("Expexted %d, Actual %d\n", 13, cpu.a)
	}

}

func TestADCAbsoluteX(t *testing.T) {
	fmt.Println("ADC Absolute X")
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = adcAbsX
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03
	cpu.x = 2

	cpu.mem[1002] = 5

	inst := cpu.Decode()
	inst()

	if cpu.a != 15 {
		t.Errorf("Expexted %d, Actual %d\n", 15, cpu.a)
	}

}

func TestADCAbsoluteY(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = adcAbsY
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03
	cpu.y = 2

	cpu.mem[1002] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 13 {
		t.Errorf("Expected %d, Actual %d\n", 13, cpu.a)
	}
}

func TestADCIndX(t *testing.T) {

	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = adcIndX
	cpu.mem[1] = 4 //Zero Page Index

	cpu.mem[4] = 0xe8
	cpu.mem[5] = 0x03

	cpu.mem[1000] = 3

	cpu.x = 0

	inst := cpu.Decode()
	inst()

	if cpu.a != 13 {
		t.Errorf("Expected %d, Actual %d\n", 13, cpu.a)
	}

}

func TestADCIndY(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10
	cpu.y = 1

	cpu.mem[0] = adcIndY
	cpu.mem[1] = 4 //Zero Page Index

	cpu.mem[4] = 0xe8
	cpu.mem[5] = 0x03

	cpu.mem[1001] = 9

	inst := cpu.Decode()
	inst()

	if cpu.a != 19 {
		t.Errorf("Expected %d, Actual %d\n", 19, cpu.a)
	}

}

func TestANDImmediate(t *testing.T) {
	cpu := newCpu()

	cpu.a = 2
	cpu.mem[0] = andImm
	cpu.mem[1] = 6

	inst := cpu.Decode()
	inst()

	//fmt.Printf("OVERFLOW: %d\n", cpu.p&1<<6)

	if cpu.a != 2 {
		t.Errorf("Expexted %d, Actual %d\n", 2, cpu.a)
	}

	if cpu.pc != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.pc)
	}
}

// AND does not affect overflow
// func TestANDOverflowCleared(t *testing.T) {
// 	cpu := newCpu()

// 	cpu.a = 10
// 	cpu.mem[0] = andImm
// 	cpu.mem[1] = 6

// 	cpu.p |= (1 << 6)

// 	inst := cpu.Decode()
// 	inst()

// 	overflow := cpu.p & (1 << 6)

// 	fmt.Printf("OVERFLOW: %d\n", overflow)

// 	if overflow != 0 {
// 		t.Errorf("Expexted %d, Actual %d\n", 0, overflow)
// 	}

// }

func TestANDZeroPage(t *testing.T) {
	cpu := newCpu()

	pcStart := cpu.pc
	cpu.a = 3
	cpu.mem[0] = andZp
	cpu.mem[1] = 4
	cpu.mem[4] = 7
	inst := cpu.Decode()
	inst()

	if cpu.a != 3 {
		t.Errorf("Expexted %d, Actual %d\n", 3, cpu.a)
	}

	expected := uint16(infoArray[andZp][Size]) + pcStart

	if cpu.pc != expected {
		t.Errorf("Expected %d, Actual %d\n", expected, cpu.pc)
	}

}

func TestANDZeroPageX(t *testing.T) {
	cpu := newCpu()
	cpu.a = 7
	cpu.x = 2

	cpu.mem[0] = andZpX
	cpu.mem[1] = 8
	cpu.mem[10] = 11

	inst := cpu.Decode()
	inst()

	if cpu.a != 3 {
		t.Errorf("Expected %d, Actual %d\n", 3, cpu.a)
	}

}

func TestANDAbs(t *testing.T) {
	cpu := newCpu()
	cpu.a = 10

	cpu.mem[0] = andAbs
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03

	cpu.mem[1000] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}

}

func TestANDAbsX(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = andAbsX
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03
	cpu.x = 2

	cpu.mem[1002] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}

}

func TestANDAbsY(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = andAbsY
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03
	cpu.y = 2

	cpu.mem[1002] = 3

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}
}

func TestANDIndX(t *testing.T) {

	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = andIndX
	cpu.mem[1] = 4 //Zero Page Index

	cpu.mem[4] = 0xe8
	cpu.mem[5] = 0x03

	cpu.mem[1000] = 3

	cpu.x = 0

	inst := cpu.Decode()
	inst()

	if cpu.a != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.a)
	}

}

func TestANDIndY(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10
	cpu.y = 1

	cpu.mem[0] = andIndY
	cpu.mem[1] = 4 //Zero Page Index

	cpu.mem[4] = 0xe8
	cpu.mem[5] = 0x03

	cpu.mem[1001] = 9

	inst := cpu.Decode()
	inst()

	if cpu.a != 8 {
		t.Errorf("Expected %d, Actual %d\n", 8, cpu.a)
	}

}

func TestASLAcc(t *testing.T) {
	cpu := newCpu()
	cpu.mem[0] = aslAcc

	cpu.a = 0xC0
	inst := cpu.Decode()
	inst()

	if cpu.a != 0x80 {
		t.Errorf("Expected %d, Actual %d\n", 0x80, cpu.a)
	}

	if !cpu.isCarry() {
		t.Errorf("Expected Flag (Carry) C == true, Actual C == %v\n", cpu.c)
	}

	if !cpu.isNegative() {
		t.Errorf("Expected Flag (Negative) N == true, Actual N == %v\n", cpu.n)
	}

	if cpu.isZero() {
		t.Errorf("Expected Flag (Zero) Z == true, Actual Z == %v\n", cpu.z)
	}
}

func TestASLZeroPage(t *testing.T) {
	cpu := newCpu()
	cpu.mem[0] = aslZp
	cpu.mem[1] = 4
	cpu.mem[4] = 0x82 //1000 0010

	inst := cpu.Decode()
	inst()

	if cpu.mem[4] != 4 {
		t.Errorf("Expected %d, Actual %d\n", 4, cpu.mem[4])
	}

	if !cpu.isCarry() {
		t.Errorf("Expected Flag (Carry) C == true, Actual C == %v\n", cpu.c)
	}

	if cpu.isNegative() {
		t.Errorf("Expected Flag (Negative) N == false, Actual N == %v\n", cpu.n)
	}

	if cpu.isZero() {
		t.Errorf("Expected Flag (Zero) Z == false, Actual Z == %v\n", cpu.z)
	}

}

func TestASLZeroPageX(t *testing.T) {
	cpu := newCpu()
	cpu.x = 1
	cpu.mem[0] = aslZpX
	cpu.mem[1] = 4
	cpu.mem[5] = 0x82 //1000 0010

	inst := cpu.Decode()
	inst()

	if cpu.mem[5] != 4 {
		t.Errorf("Expected %d, Actual %d\n", 4, cpu.mem[5])
	}

	if !cpu.isCarry() {
		t.Errorf("Expected Flag (Carry) C == true, Actual C == %v\n", cpu.c)
	}

	if cpu.isNegative() {
		t.Errorf("Expected Flag (Negative) N == false, Actual N == %v\n", cpu.n)
	}

	if cpu.isZero() {
		t.Errorf("Expected Flag (Zero) Z == false, Actual Z == %v\n", cpu.z)
	}

}

func TestASLAbsolute(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = aslAbs
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03

	cpu.mem[1000] = 128 // 1000 0000

	inst := cpu.Decode()
	inst()

	if cpu.mem[1000] != 0 {
		t.Errorf("Expected %d, Actual %d\n", 0, cpu.mem[1000])
	}

	if !cpu.isCarry() {
		t.Errorf("Expected Flag (Carry) C == true, Actual C == %v\n", cpu.c)
	}

	if cpu.isNegative() {
		t.Errorf("Expected Flag (Negative) N == false, Actual N == %v\n", cpu.n)
	}

	if !cpu.isZero() {
		t.Errorf("Expected Flag (Zero) Z == true, Actual Z == %v\n", cpu.z)
	}

}

func TestASLAbsoluteX(t *testing.T) {
	cpu := newCpu()

	cpu.a = 10

	cpu.mem[0] = aslAbsX
	//Value 1000 Dec
	cpu.mem[1] = 0xe8
	cpu.mem[2] = 0x03
	cpu.x = 2

	cpu.mem[1002] = 128 // 1000 0000

	inst := cpu.Decode()
	inst()

	if cpu.mem[1000] != 0 {
		t.Errorf("Expected %d, Actual %d\n", 0, cpu.mem[1000])
	}

	if !cpu.isCarry() {
		t.Errorf("Expected Flag (Carry) C == true, Actual C == %v\n", cpu.c)
	}

	if cpu.isNegative() {
		t.Errorf("Expected Flag (Negative) N == false, Actual N == %v\n", cpu.n)
	}

	if !cpu.isZero() {
		t.Errorf("Expected Flag (Zero) Z == true, Actual Z == %v\n", cpu.z)
	}

}

func TestBCC(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 0x80
	cpu.mem[19] = aslAcc
	cpu.mem[20] = bcc
	cpu.mem[21] = 4
	cpu.mem[22] = aslAcc
	cpu.mem[23] = bcc
	cpu.mem[24] = 0xfa

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 22 {
		t.Errorf("Expected %d, Actual %d\n", 22, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 17 {
		t.Errorf("Expected %d, Actual %d\n", 17, cpu.pc)
	}

}

func TestBCS(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 0x80
	cpu.mem[19] = aslAcc
	cpu.mem[20] = bcs
	cpu.mem[21] = 2
	cpu.mem[22] = aslAcc
	cpu.mem[23] = bcs
	cpu.mem[24] = 0xfa

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 22 {
		t.Errorf("Expected %d, Actual %d\n", 22, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 25 {
		t.Errorf("Expected %d, Actual %d\n", 25, cpu.pc)
	}

}

func TestBEQ(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 0x80
	cpu.mem[19] = aslAcc
	cpu.mem[20] = beq
	cpu.mem[21] = 2
	cpu.mem[22] = aslAcc
	cpu.mem[23] = beq
	cpu.mem[24] = 0xfa

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 22 {
		t.Errorf("Expected %d, Actual %d\n", 22, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 17 {
		t.Errorf("Expected %d, Actual %d\n", 17, cpu.pc)
	}

}

func TestBIT(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 10
	cpu.a = 3
	cpu.mem[20] = 2
	cpu.mem[10] = bitZp
	cpu.mem[11] = 20

	inst := cpu.Decode()
	cycles := inst()

	if cpu.z != false {
		t.Errorf("Expected %v, Actual %v\n", false, cpu.z)
	}

	m7 := cpu.i2b((cpu.mem[20] & (1 << 7)))
	if cpu.n != m7 {
		t.Errorf("Expected %v, Actual %v\n", m7, cpu.n)
	}

	m6 := cpu.i2b((cpu.mem[20] & (1 << 6)))
	if cpu.n != m6 {
		t.Errorf("Expected %v, Actual %v\n", m6, cpu.v)
	}

	expectedCycles := infoArray[bitZp][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.mem[0x0104] = 4
	cpu.mem[12] = bitAbs
	cpu.mem[13] = 0x04
	cpu.mem[14] = 0x01

	inst = cpu.Decode()
	cycles = inst()

	if cpu.z != true {
		t.Errorf("Expected %v, Actual %v\n", true, cpu.z)
	}

}

func TestBMI(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 0xC0
	cpu.mem[19] = aslAcc
	cpu.mem[20] = bmi
	cpu.mem[21] = 2
	cpu.mem[22] = aslAcc
	cpu.mem[23] = bmi
	cpu.mem[24] = 0xfa

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 22 {
		t.Errorf("Expected %d, Actual %d\n", 22, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 25 {
		t.Errorf("Expected %d, Actual %d\n", 25, cpu.pc)
	}

}

func TestBNE(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 0x80
	cpu.mem[19] = aslAcc
	cpu.mem[20] = bne
	cpu.mem[21] = 2
	cpu.mem[22] = aslAcc
	cpu.mem[23] = bne
	cpu.mem[24] = 0xfa

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 22 {
		t.Errorf("Expected %d, Actual %d\n", 22, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 25 {
		t.Errorf("Expected %d, Actual %d\n", 25, cpu.pc)
	}

}

func TestBPL(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 0xC0
	cpu.mem[19] = aslAcc
	cpu.mem[20] = bpl
	cpu.mem[21] = 2
	cpu.mem[22] = aslAcc
	cpu.mem[23] = bpl
	cpu.mem[24] = 0xfa

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 22 {
		t.Errorf("Expected %d, Actual %d\n", 22, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 17 {
		t.Errorf("Expected %d, Actual %d\n", 17, cpu.pc)
	}

}

func TestBRK(t *testing.T) {
	cpu := newCpu()

	//ISR Vector
	cpu.mem[0xFFFE] = uint8(0x00) //pcl
	cpu.mem[0xFFFF] = uint8(0x10) //pch

	//ISR
	cpu.mem[0x1000] = adcImm
	cpu.mem[0x1001] = 4
	cpu.mem[0x1002] = rti

	cpu.pc = 10
	cpu.a = 2
	cpu.mem[10] = adcImm
	cpu.mem[11] = 5

	cpu.mem[12] = brk
	cpu.mem[13] = nop

	cpu.mem[14] = adcImm
	cpu.mem[15] = 1

	//Execute adcImm 5
	inst := cpu.Decode()
	inst()

	if cpu.a != 7 {
		t.Errorf("Expected %d, Actual %d\n", 7, cpu.a)
	}

	//Execute brk
	inst = cpu.Decode()
	inst()

	if cpu.pc != 0x1000 {
		t.Errorf("Expected %d, Actual %d\n", 4096, cpu.pc)
	}

	//Execute adcImm 4
	inst = cpu.Decode()
	inst()

	if cpu.a != 11 {
		t.Errorf("Expected %d, Actual %d\n", 11, cpu.a)
	}

	//Execute rti
	inst = cpu.Decode()
	inst()

	if cpu.pc != 14 {
		t.Errorf("Expected %d, Actual %d\n", 14, cpu.pc)
	}

	//Execute adcImm 1
	inst = cpu.Decode()
	inst()

	if cpu.a != 12 {
		t.Errorf("Expected %d, Actual %d\n", 12, cpu.a)
	}

}

func TestBVC(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 125

	cpu.mem[19] = adcImm
	cpu.mem[20] = 1 //No overflow, expect to branch
	cpu.mem[21] = bvc
	cpu.mem[22] = 19

	cpu.mem[40] = adcImm
	cpu.mem[41] = 10 //Overflow, expect not to branch
	cpu.mem[42] = bvc
	cpu.mem[43] = 7

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 40 {
		t.Errorf("Expected %d, Actual %d\n", 40, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 44 {
		t.Errorf("Expected %d, Actual %d\n", 44, cpu.pc)
	}

}

func TestBVS(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 19
	cpu.a = 125

	cpu.mem[19] = adcImm
	cpu.mem[20] = 1 //No overflow, expect not to branch
	cpu.mem[21] = bvs
	cpu.mem[22] = 19

	cpu.mem[23] = adcImm
	cpu.mem[24] = 10 //Overflow, expect to branch
	cpu.mem[25] = bvs
	cpu.mem[26] = 5

	inst := cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 23 {
		t.Errorf("Expected %d, Actual %d\n", 23, cpu.pc)
	}

	inst = cpu.Decode()
	inst()

	inst = cpu.Decode()
	inst()

	if cpu.pc != 30 {
		t.Errorf("Expected %d, Actual %d\n", 30, cpu.pc)
	}

}

func TestCMP(t *testing.T) {
	println(" --- EXeCUTING CMP TEST ---")
	cpu := newCpu()
	cpu.a = 10
	cpu.pc = 20
	cpu.mem[20] = cmpImm
	cpu.mem[21] = 5

	cpu.mem[22] = cmpImm
	cpu.mem[23] = 10

	cpu.mem[24] = cmpImm
	cpu.mem[25] = 15

	cpu.x = 2
	cpu.mem[26] = cmpAbsX
	cpu.mem[27] = 0xff
	cpu.mem[28] = 0x01
	cpu.mem[0x201] = 10

	inst := cpu.Decode()
	cycles := inst()

	if cpu.c || cpu.z || cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", false, false, false, cpu.c, cpu.z, cpu.n)
	}

	expectedCycles := infoArray[cmpImm][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	inst = cpu.Decode()
	cycles = inst()

	if cpu.c || !cpu.z || cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", false, true, false, cpu.c, cpu.z, cpu.n)
	}

	inst = cpu.Decode()
	cycles = inst()

	if !cpu.c || cpu.z || !cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", true, false, true, cpu.c, cpu.z, cpu.n)
	}

	inst = cpu.Decode()
	cycles = inst()

	expectedCycles = infoArray[cmpAbsX][Cycles]
	if cycles != expectedCycles+1 {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles+1, cycles)
	}

	cpu.mem[29] = cmpAbsX
	cpu.mem[30] = 04
	cpu.mem[31] = 01

	inst = cpu.Decode()
	cycles = inst()

	expectedCycles = infoArray[cmpAbsX][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestCPX(t *testing.T) {
	println(" --- EXeCUTING CPX TEST ---")
	cpu := newCpu()
	cpu.x = 10
	cpu.pc = 20
	cpu.mem[20] = cpxImm
	cpu.mem[21] = 5

	cpu.mem[22] = cpxImm
	cpu.mem[23] = 10

	cpu.mem[24] = cpxImm
	cpu.mem[25] = 15

	inst := cpu.Decode()
	inst()

	if cpu.c || cpu.z || cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", false, false, false, cpu.c, cpu.z, cpu.n)
	}

	inst = cpu.Decode()
	inst()

	if cpu.c || !cpu.z || cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", false, true, false, cpu.c, cpu.z, cpu.n)
	}

	inst = cpu.Decode()
	inst()

	if !cpu.c || cpu.z || !cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", true, false, true, cpu.c, cpu.z, cpu.n)
	}

	cpu.mem[26] = cpxZp
	cpu.mem[27] = 10
	inst = cpu.Decode()
	cycles := inst()

	expectedCycles := infoArray[cpxZp][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.mem[28] = cpxAbs
	cpu.mem[29] = 10
	inst = cpu.Decode()
	cycles = inst()

	expectedCycles = infoArray[cpxAbs][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestCPY(t *testing.T) {
	println(" --- EXeCUTING CPX TEST ---")
	cpu := newCpu()
	cpu.y = 10
	cpu.pc = 20
	cpu.mem[20] = cpyImm
	cpu.mem[21] = 5

	cpu.mem[22] = cpyImm
	cpu.mem[23] = 10

	cpu.mem[24] = cpyImm
	cpu.mem[25] = 15

	inst := cpu.Decode()
	inst()

	if cpu.c || cpu.z || cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", false, false, false, cpu.c, cpu.z, cpu.n)
	}

	inst = cpu.Decode()
	inst()

	if cpu.c || !cpu.z || cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", false, true, false, cpu.c, cpu.z, cpu.n)
	}

	inst = cpu.Decode()
	inst()

	if !cpu.c || cpu.z || !cpu.n {
		t.Errorf("Expected Status C: %v, Z: %v, N: %v - Actual C: %v, Z: %v, N: %v\n", true, false, true, cpu.c, cpu.z, cpu.n)
	}

	cpu.mem[26] = cpyZp
	cpu.mem[27] = 10
	inst = cpu.Decode()
	cycles := inst()

	expectedCycles := infoArray[cpyZp][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.mem[28] = cpyAbs
	cpu.mem[29] = 10
	inst = cpu.Decode()
	cycles = inst()

	expectedCycles = infoArray[cpyAbs][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestDEC(t *testing.T) {

	cpu := newCpu()

	cpu.mem[4] = 10
	cpu.pc = 20
	cpu.mem[20] = decAbs
	cpu.mem[21] = 4

	inst := cpu.Decode()
	cycles := inst()

	if cpu.mem[4] != 9 {
		t.Errorf("Expected %d, Actual %d\n", 9, cpu.mem[4])
	}

	expectedCycles := infoArray[decAbs][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.x = 2
	cpu.mem[23] = decAbsX
	cpu.mem[24] = 40
	cpu.mem[25] = 0
	cpu.mem[42] = 3

	inst = cpu.Decode()
	cycles = inst()

	if cpu.mem[42] != 2 {
		t.Errorf("Expected %d, Actual %d\n", 2, cpu.mem[42])
	}

	expectedCycles = infoArray[decAbsX][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}
}

func TestDEX(t *testing.T) {
	cpu := newCpu()
	cpu.x = 4
	cpu.pc = 20
	cpu.mem[20] = dex

	inst := cpu.Decode()
	cycles := inst()

	if cpu.x != 3 {
		t.Errorf("Expected %d, Actual %d\n", 3, cpu.x)
	}

	expectedCycles := infoArray[dex][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestDEY(t *testing.T) {
	cpu := newCpu()
	cpu.y = 4
	cpu.pc = 20
	cpu.mem[20] = dey

	inst := cpu.Decode()
	cycles := inst()

	if cpu.y != 3 {
		t.Errorf("Expected %d, Actual %d\n", 3, cpu.y)
	}

	expectedCycles := infoArray[dex][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestEOR(t *testing.T) {
	cpu := newCpu()
	cpu.a = 0xF8
	//cpu.mem[4] = 0xFF

	cpu.pc = 20
	cpu.mem[20] = eorImm
	cpu.mem[21] = 0xFF

	inst := cpu.Decode()
	cycles := inst()

	if cpu.a != 7 {
		t.Errorf("Expected %d, Actual %d\n", 7, cpu.a)
	}

	expectedCycles := infoArray[eorImm][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.x = 4
	cpu.a = 0xFF
	cpu.mem[22] = eorAbsX
	cpu.mem[23] = 0xFF
	cpu.mem[24] = 0x01
	cpu.mem[0x0203] = 3

	inst = cpu.Decode()
	cycles = inst()

	if cpu.a != 0xfc {
		t.Errorf("Expected %d, Actual %d\n", 0xfc, cpu.a)
	}

	expectedCycles = infoArray[eorAbsX][Cycles] + 1
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	//cpu.pc = 25
	cpu.a = 42
	cpu.y = 5
	cpu.mem[25] = eorIndY
	cpu.mem[26] = 40
	cpu.mem[40] = 0x00
	cpu.mem[41] = 0x01
	cpu.mem[0x105] = 15

	inst = cpu.Decode()
	cycles = inst()

	if cpu.a != 37 {
		t.Errorf("Expected %d, Actual %d\n", 37, cpu.a)
	}

	expectedCycles = infoArray[eorIndY][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestINC(t *testing.T) {
	cpu := newCpu()
	cpu.mem[20] = incAbs
	cpu.mem[21] = 0
	cpu.mem[22] = 0x10
	cpu.mem[0x1000] = 7
	cpu.pc = 20

	inst := cpu.Decode()
	cycles := inst()

	if cpu.mem[0x1000] != 8 {
		t.Errorf("Expected %d, Actual %d\n", 6, cpu.mem[0x1000])
	}

	expectedCycles := infoArray[incAbs][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestINX(t *testing.T) {
	cpu := newCpu()
	cpu.x = 4
	cpu.pc = 20
	cpu.mem[20] = inx

	inst := cpu.Decode()
	cycles := inst()

	if cpu.x != 5 {
		t.Errorf("Expected %d, Actual %d\n", 5, cpu.x)
	}

	expectedCycles := infoArray[inx][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestINY(t *testing.T) {
	cpu := newCpu()
	cpu.y = 4
	cpu.pc = 20
	cpu.mem[20] = iny

	inst := cpu.Decode()
	cycles := inst()

	if cpu.y != 5 {
		t.Errorf("Expected %d, Actual %d\n", 5, cpu.y)
	}

	expectedCycles := infoArray[iny][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestJMP(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20

	cpu.mem[20] = jmpAbs
	cpu.mem[21] = 0x00
	cpu.mem[22] = 0x10

	cpu.mem[0x1000] = 0x00
	cpu.mem[0x1001] = 0x20

	inst := cpu.Decode()
	cycles := inst()

	if cpu.pc != 0x1000 {
		t.Errorf("Expected %v, Actual %v\n", 0x1000, cpu.pc)
	}

	expectedCycles := infoArray[jmpAbs][Cycles]

	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.pc = 20
	cpu.mem[20] = jmpInd

	inst = cpu.Decode()
	cycles = inst()

	if cpu.pc != 0x2000 {
		t.Errorf("Expected %v, Actual %v\n", 0x2000, cpu.pc)
	}

	expectedCycles = infoArray[jmpInd][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestJSR(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20

	cpu.mem[20] = jsr
	cpu.mem[21] = 0x01
	cpu.mem[22] = 0x10

	inst := cpu.Decode()
	cycles := inst()

	if cpu.pc != 0x1001 {
		t.Errorf("Expected %v, Actual %v\n", 0x1001, cpu.pc)
	}

	//Ensure correct address has been pushed to the stack
	pcl := cpu.mem[cpu.s+1]
	pch := cpu.mem[cpu.s+2]

	if pcl != 22 {
		t.Errorf("Expected %v, Actual %v\n", 22, pcl)
	}

	if pch != 0x00 {
		t.Errorf("Expected %v, Actual %v\n", 0, pch)
	}

	expectedCycles := infoArray[jsr][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestLDA(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20

	cpu.mem[20] = ldaImm
	cpu.mem[21] = 6

	inst := cpu.Decode()
	cycles := inst()

	if cpu.a != 6 {
		t.Errorf("Expected %v, Actual %v\n", 6, cpu.a)
	}

	expectedCycles := infoArray[ldaImm][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.mem[22] = ldaAbsX
	cpu.mem[23] = 0xFF
	cpu.mem[24] = 0x01
	cpu.x = 1

	cpu.mem[0x200] = 42

	inst = cpu.Decode()
	cycles = inst()

	if cpu.a != 42 {
		t.Errorf("Expected %v, Actual %v\n", 42, cpu.a)
	}

	expectedCycles = infoArray[ldaAbsX][Cycles] + 1 //Page boundary crossed, add 1
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestLDX(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20

	cpu.mem[20] = ldxImm
	cpu.mem[21] = 7

	inst := cpu.Decode()
	cycles := inst()

	if cpu.x != 7 {
		t.Errorf("Expected %v, Actual %v\n", 7, cpu.x)
	}

	expectedCycles := infoArray[ldxImm][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.x = 0
	cpu.y = 4
	cpu.mem[0x203] = 8
	cpu.mem[22] = ldxAbsY
	cpu.mem[23] = 0xFF
	cpu.mem[24] = 0x01

	inst = cpu.Decode()
	cycles = inst()

	if cpu.x != 8 {
		t.Errorf("Expected %d, Actual %d\n", 8, cpu.x)
	}

	expectedCycles = infoArray[ldxAbsY][Cycles] + 1
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestLDY(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20

	cpu.mem[20] = ldyImm
	cpu.mem[21] = 8

	inst := cpu.Decode()
	cycles := inst()

	if cpu.y != 8 {
		t.Errorf("Expected %v, Actual %v\n", 8, cpu.y)
	}

	expectedCycles := infoArray[ldyImm][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	cpu.x = 4
	cpu.y = 0
	cpu.mem[0x203] = 20
	cpu.mem[22] = ldyAbsX
	cpu.mem[23] = 0xFF
	cpu.mem[24] = 0x01

	inst = cpu.Decode()
	cycles = inst()

	if cpu.y != 20 {
		t.Errorf("Expected %v, Actual %v\n", 20, cpu.y)
	}

	expectedCycles = infoArray[ldyAbsX][Cycles] + 1
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestLSR(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.a = 4

	cpu.mem[20] = lsrAcc

	cpu.x = 2
	cpu.mem[21] = lsrAbsX
	cpu.mem[22] = 0x00
	cpu.mem[23] = 0x10
	cpu.mem[0x1002] = 16

	inst := cpu.Decode()
	cycles := inst()

	if cpu.a != 2 {
		t.Errorf("Expected %v, Actual %v\n", 2, cpu.a)
	}

	expectedCycles := infoArray[lsrAcc][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

	inst = cpu.Decode()
	cycles = inst()

	if cpu.mem[0x1002] != 8 {
		t.Errorf("Expected %v, Actual %v\n", 8, cpu.mem[0x1002])
	}

	expectedCycles = infoArray[lsrAbsX][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %d, Actual %d\n", expectedCycles, cycles)
	}

}

func TestORA(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.a = 5

	cpu.mem[20] = oraImm
	cpu.mem[21] = 0x02

	inst := cpu.Decode()
	cycles := inst()

	if cpu.a != 7 {
		t.Errorf("Expected %v, Actual %v\n", 7, cpu.a)
	}

	expectedCycles := infoArray[oraImm][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

	cpu.a = 6
	cpu.x = 5
	cpu.mem[22] = oraAbsX
	cpu.mem[23] = 0xFF
	cpu.mem[24] = 0x01
	cpu.mem[0x204] = 8

	inst = cpu.Decode()
	cycles = inst()

	if cpu.a != 14 {
		t.Errorf("Expected %v, Actual %v\n", 14, cpu.a)
	}

	expectedCycles = infoArray[oraAbsX][Cycles] + 1 //Page boundary crossed
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestPHA(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.a = 5

	cpu.mem[20] = pha
	cpu.mem[21] = 0x02

	inst := cpu.Decode()
	cycles := inst()

	stackVal := cpu.readImm(uint16((01 << 4) | (cpu.s + 1)))

	if stackVal != 5 {
		t.Errorf("Expected %v, Actual %v\n", 5, stackVal)
	}

	expectedCycles := infoArray[pha][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestPHP(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.mem[20] = php
	cpu.mem[21] = 0x02

	cpu.v = true
	cpu.n = true
	cpu.z = true

	origStatus := cpu.readStatus()

	inst := cpu.Decode()
	cycles := inst()

	loc := uint16((01 << 4) | (cpu.s + 1))
	stackVal := cpu.readImm(loc)

	if stackVal != origStatus {
		t.Errorf("Expected %v, Actual %v\n", origStatus, stackVal)
	}

	expectedCycles := infoArray[php][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestPLA(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20

	cpu.mem[20] = pla
	cpu.mem[21] = 0x02

	var val byte = 0xBE

	cpu.push(val)

	inst := cpu.Decode()
	cycles := inst()

	if cpu.a != 0xBE {
		t.Errorf("Expected %v, Actual %v\n", 0xBE, cpu.a)
	}

	expectedCycles := infoArray[pla][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestPLP(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20

	cpu.mem[20] = plp
	cpu.mem[21] = 0x02

	cpu.push(194)

	inst := cpu.Decode()
	cycles := inst()

	status := cpu.readStatus()
	if status != 194 {
		t.Errorf("Expected %v, Actual %v\n", 190, status)
	}

	expectedCycles := infoArray[plp][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestROL(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.a = 0xC0
	cpu.c = false

	cpu.mem[20] = rolAcc
	cpu.mem[21] = rolAcc
	cpu.mem[22] = rolAcc

	inst := cpu.Decode()
	cycles := inst()

	if cpu.a != 0x80 {
		t.Errorf("Expected %v, Actual %v\n", 0x80, cpu.a)
	}

	expectedCycles := infoArray[rolAcc][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

	inst = cpu.Decode()
	inst()

	if cpu.a != 0x01 {
		t.Errorf("Expected %v, Actual %v\n", 0x01, cpu.a)
	}

	inst = cpu.Decode()
	inst()

	if cpu.a != 0x03 {
		t.Errorf("Expected %v, Actual %v\n", 0x03, cpu.a)
	}

	cpu.x = 2
	cpu.mem[23] = rolAbsX
	cpu.mem[24] = 0xFF
	cpu.mem[25] = 0x00

	inst = cpu.Decode()
	cycles = inst()

	expectedCycles = infoArray[rolAbsX][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestROR(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.a = 0xC1
	cpu.c = false

	cpu.mem[20] = rorAcc
	cpu.mem[21] = rorAcc
	cpu.mem[22] = rorAcc

	cpu.pc = 20
	inst := cpu.Decode()
	cycles := inst()

	if cpu.a != 0x60 {
		t.Errorf("Expected %v, Actual %v\n", 0x60, cpu.a)
	}

	expectedCycles := infoArray[rorAcc][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

	inst = cpu.Decode()
	inst()

	if cpu.a != 0xB0 {
		t.Errorf("Expected %v, Actual %v\n", 0xB0, cpu.a)
	}

	inst = cpu.Decode()
	inst()

	if cpu.a != 0x58 {
		t.Errorf("Expected %v, Actual %v\n", 0x58, cpu.a)
	}

	cpu.x = 2
	cpu.mem[23] = rorAbsX
	cpu.mem[24] = 0xFF
	cpu.mem[25] = 0x01
	cpu.mem[0x201] = 8

	inst = cpu.Decode()
	cycles = inst()

	if cpu.mem[0x201] != 4 {
		t.Errorf("Expected %v, Actual %v\n", 4, cpu.mem[0x201])
	}

	expectedCycles = infoArray[rorAbsX][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestRTI(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.mem[0x1000] = nop
	cpu.mem[20] = rti

	cpu.push(0x10) //pch
	cpu.push(0x00) //pcl
	cpu.pushStatusToStack()

	inst := cpu.Decode()
	cycles := inst()

	if cpu.pc != 0x1000 {
		t.Errorf("Expected %v, Actual %v\n", 0x1000, cpu.pc)
	}

	expectedCycles := infoArray[rti][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}

}

func TestRTS(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.mem[0x1000] = nop
	cpu.mem[20] = rts

	cpu.push(0x10)
	cpu.push(0x00)

	inst := cpu.Decode()
	cycles := inst()

	if cpu.pc != 0x1001 {
		t.Errorf("Expected %v, Actual %v\n", 0x1001, cpu.pc)
	}

	expectedCycles := infoArray[rts][Cycles]
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}
}

func TestSBC(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.a = 10
	cpu.c = true

	cpu.mem[20] = sbcImm
	cpu.mem[21] = 10

	cpu.mem[22] = sbcImm
	cpu.mem[23] = 1

	inst := cpu.Decode()
	inst()

	if cpu.a != 0 {
		t.Errorf("Expected %v, Actual %v\n", 0, cpu.a)
	}

	if cpu.z != true {
		t.Errorf("Expected %v, Actual %v\n", true, cpu.z)
	}

	if cpu.c != true {
		t.Errorf("Expected %v, Actual %v\n", true, cpu.c)
	}

	inst = cpu.Decode()
	inst()

	if cpu.c != false {
		t.Errorf("Expected %v, Actual %v\n", false, cpu.c)
	}

	if cpu.n != true {
		t.Errorf("Expected %v, Actual %v\n", true, cpu.n)
	}

	if cpu.z != false {
		t.Errorf("Expected %v, Actual %v\n", false, cpu.z)
	}

	cpu.c = true
	cpu.y = 5
	cpu.a = 20
	cpu.mem[24] = sbcAbsY
	cpu.mem[25] = 0xFF
	cpu.mem[26] = 0x01
	cpu.mem[0x0204] = 5

	inst = cpu.Decode()
	cycles := inst()

	if cpu.a != 15 {
		t.Errorf("Expected %v, Actual %v\n", 15, cpu.a)
	}

	expectedCycles := infoArray[sbcAbsY][Cycles] + 1 //page boundary crossed
	if cycles != expectedCycles {
		t.Errorf("Expected %v, Actual %v\n", expectedCycles, cycles)
	}
}

func TestCLC(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.c = true

	cpu.mem[20] = clc

	inst := cpu.Decode()
	cycles := inst()

	if cpu.c != false {
		t.Errorf("Expected %v, Actual %v\n", false, cpu.c)
	}

	if cycles != 2 {
		t.Errorf("Expected %v, Actual %v\n", 2, cycles)
	}

}

func TestCLD(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.d = true

	cpu.mem[20] = cld

	inst := cpu.Decode()
	cycles := inst()

	if cpu.d != false {
		t.Errorf("Expected %v, Actual %v\n", false, cpu.d)
	}

	if cycles != 2 {
		t.Errorf("Expected %v, Actual %v\n", 2, cycles)
	}

}

func TestCLI(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.i = true

	cpu.mem[20] = cli

	inst := cpu.Decode()
	cycles := inst()

	if cpu.i != false {
		t.Errorf("Expected %v, Actual %v\n", false, cpu.i)
	}

	if cycles != 2 {
		t.Errorf("Expected %v, Actual %v\n", 2, cycles)
	}

}

func TestCLV(t *testing.T) {
	cpu := newCpu()
	cpu.pc = 20
	cpu.v = true

	cpu.mem[20] = clv

	inst := cpu.Decode()
	cycles := inst()

	if cpu.v != false {
		t.Errorf("Expected %v, Actual %v\n", false, cpu.v)
	}

	if cycles != 2 {
		t.Errorf("Expected %v, Actual %v\n", 2, cycles)
	}

}

// func TestCMP(t *testing.T) {
// 	cpu := newCpu()
// 	cpu.pc = 20
// 	cpu.a = 4

// 	cpu.mem[20] = cmpImm
// 	cpu.mem[21] = 4

// 	inst := cpu.Decode()
// 	cycles := inst()

// 	if cpu.z != true {
// 		t.Errorf("Expected %v, Actual %v\n", true, cpu.z)
// 	}

// }

func newCpu() *Cpu {
	cpu := NewCPU(make([]byte, 1024*64))
	cpu.pc = 0
	return cpu
}
