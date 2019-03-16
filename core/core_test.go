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

	fmt.Printf("FLAGS: %d\n", cpu.p)

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

	fmt.Printf("OVERFLOW: %d\n", cpu.p&1<<6)

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

}

func newCpu() *Cpu {
	cpu := NewCPU(make([]byte, 1024*64))
	cpu.pc = 0
	return cpu
}
