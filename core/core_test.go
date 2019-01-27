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
	// cpu.mem[2] = adcImm
	// cpu.mem[3] = 2

	cpu.p |= 1

	//cpu.setCarryStatus(254, 1, 0)

	inst := cpu.Decode()
	inst()

	if cpu.getCarry() != 1 {
		t.Errorf("Expexted %d, Actual %d\n", 1, cpu.p&1)
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

	if cpu.p&1 == 0 {
		t.Errorf("Expected Flag (Carry) C == 1, Actual C == %d\n", cpu.p&1)
	}

	if cpu.p&0x80 == 0 {
		t.Errorf("Expected Flag (Negative) N == 1, Actual N == %d\n", (cpu.p&0x80)>>7)
	}

	if cpu.p&2 != 0 {
		t.Errorf("Expected Flag (Zero) Z == 0, Actual Z == %d\n", cpu.p&2)
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

	if cpu.p&1 == 0 {
		t.Errorf("Expected Flag (Carry) C == 1, Actual C == %d\n", cpu.p&1)
	}

	if cpu.p&0x80 == 1 {
		t.Errorf("Expected Flag (Negative) N == 0, Actual N == %d\n", (cpu.p&0x80)>>7)
	}

	if cpu.p&2 != 0 {
		t.Errorf("Expected Flag (Zero) Z == 0, Actual Z == %d\n", cpu.p&2)
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

	if cpu.p&1 == 0 {
		t.Errorf("Expected Flag (Carry) C == 1, Actual C == %d\n", cpu.p&1)
	}

	if cpu.p&0x80 == 1 {
		t.Errorf("Expected Flag (Negative) N == 0, Actual N == %d\n", (cpu.p&0x80)>>7)
	}

	if cpu.p&2 != 0 {
		t.Errorf("Expected Flag (Zero) Z == 0, Actual Z == %d\n", cpu.p&2)
	}

}

func newCpu() *Cpu {
	cpu := NewCPU(make([]byte, 1024*16))
	cpu.pc = 0
	return cpu
}
