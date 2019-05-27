package core

import (
	"errors"
	"testing"
)

func TestMapping(t *testing.T) {
	if error := isValidMapping(opcodes); error != nil {
		t.Error(error)
	}
}

var opcodes = []int{
	brk, oraIndX, fe02, fe03, fe04, oraZp, aslZp, fe07, php, oraImm, aslAcc, fe0b, fe0c, oraAbs, aslAbs, fe0f,
	bpl, oraIndY, fe12, fe13, fe14, oraZpX, aslZpX, fe17, clc, oraAbsY, fe1a, fe1b, fe1c, oraAbsX, aslAbsX, fe1f,
	jsr, andIndX, fe22, fe23, bitZp, andZp, rolZp, fe27, plp, andImm, rolAcc, fe2b, bitAbs, andAbs, rolAbs, fe2f,
	bmi, andIndY, fe32, fe33, fe34, andZpX, rolZpX, fe37, sec, andAbsY, fe3a, fe3b, fe3c, andAbsX, rolAbsX, fe3f,
	rti, eorIndX, fe42, fe43, fe44, eorZp, lsrZp, fe47, pha, eorImm, lsrAcc, fe4b, jmpAbs, eorAbs, lsrAbs, fe4f,
	bvc, eorIndY, fe52, fe53, fe54, eorZpX, lsrZpX, fe57, cli, eorAbsY, fe5a, fe5b, fe5c, eorAbsX, lsrAbsX, fe5f,
	rts, adcIndX, fe62, fe63, fe64, adcZp, rorZp, fe67, pla, adcImm, rorAcc, fe6b, jmpInd, adcAbs, rorAbs, fe6f,
	bvs, adcIndY, fe72, fe73, fe74, adcZpX, rorZpX, fe77, sei, adcAbsY, fe7a, fe7b, fe7c, adcAbsX, rorAbsX, fe7f,
	fe80, staIndX, fe82, fe83, styZp, staZp, stxZp, fe87, dey, fe89, txa, fe8b, styAbs, staAbs, stxAbs, fe8f,
	bcc, staIndY, fe92, fe93, styZpX, staZpX, stxZpY, fe97, tya, staAbsY, txs, fe9b, fe9c, staAbsX, fe9e, fe9f,
	ldyImm, ldaIndX, ldxImm, fea3, ldyZp, ldaZp, ldxZp, fea7, tay, ldaImm, tax, feab, ldyAbs, ldaAbs, ldxAbs, feaf,
	bcs, ldaIndY, feb2, feb3, ldyZpX, ldaZpX, ldxZpY, feb7, clv, ldaAbsY, tsx, febb, ldyAbsX, ldaAbsX, ldxAbsY, febf,
	cpyImm, cmpIndX, fec2, fec3, cpyZp, cmpZp, decZp, fec7, iny, cmpImm, dex, fecb, cpyAbs, cmpAbs, decAbs, fecf,
	bne, cmpIndY, fed2, fed3, fed4, cmpZpX, decZpX, fed7, cld, cmpAbsY, feda, fedb, fedc, cmpAbsX, decAbsX, fedf,
	cpxImm, sbcIndX, fee2, fee3, cpxZp, sbcZp, incZp, fee7, inx, sbcImm, nop, feeb, cpxAbs, sbcAbs, incAbs, feef,
	beq, sbcIndY, fef2, fef3, fef4, sbcZpX, incZpX, fef7, sed, sbcAbsY, fefa, fefb, fefc, sbcAbsX, incAbsX, feff,
}

var opsizes = []int{
	1, 2, 0, 0, 0, 2, 2, 0, 1, 2, 1, 0, 0, 3, 3, 0,
	2, 2, 0, 0, 0, 2, 2, 0, 1, 3, 0, 0, 0, 3, 3, 0,
	3, 2, 0, 0, 2, 2, 2, 0, 1, 2, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 0, 2, 2, 0, 1, 3, 0, 0, 0, 3, 3, 0,
	1, 2, 0, 0, 0, 2, 2, 0, 1, 2, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 0, 2, 2, 0, 1, 3, 0, 0, 0, 3, 3, 0,
	1, 2, 0, 0, 0, 2, 2, 0, 1, 2, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 0, 2, 2, 0, 1, 3, 0, 0, 0, 3, 3, 0,
	0, 2, 0, 0, 2, 2, 2, 0, 1, 0, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 2, 2, 2, 0, 1, 3, 1, 0, 0, 3, 0, 0,
	2, 2, 2, 0, 2, 2, 2, 0, 1, 2, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 2, 2, 2, 0, 1, 3, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 2, 2, 2, 0, 1, 2, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 0, 2, 2, 0, 1, 3, 0, 0, 0, 3, 3, 0,
	2, 2, 0, 0, 2, 2, 2, 0, 1, 2, 1, 0, 3, 3, 3, 0,
	2, 2, 0, 0, 0, 2, 2, 0, 1, 3, 0, 0, 0, 3, 3, 0,
}

var opcycles = []int{
	7, 6, 0, 0, 0, 3, 5, 0, 3, 2, 2, 0, 0, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	6, 6, 0, 0, 3, 3, 5, 0, 4, 2, 2, 0, 4, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	6, 6, 0, 0, 0, 3, 5, 0, 3, 2, 2, 0, 3, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	6, 6, 0, 0, 0, 3, 5, 0, 4, 2, 2, 0, 5, 4, 6, 0,
	2, 6, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	0, 6, 0, 0, 3, 3, 3, 0, 2, 0, 2, 0, 4, 4, 4, 0,
	2, 6, 0, 0, 4, 4, 4, 0, 2, 5, 2, 0, 0, 5, 0, 0,
	2, 6, 2, 0, 3, 3, 3, 0, 2, 2, 2, 0, 4, 4, 4, 0,
	2, 5, 0, 0, 4, 4, 4, 0, 2, 4, 2, 0, 4, 4, 4, 0,
	2, 6, 0, 0, 3, 4, 5, 0, 2, 2, 2, 0, 4, 4, 6, 0,
	2, 6, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	2, 6, 0, 0, 3, 3, 5, 0, 2, 2, 2, 0, 4, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
}

func isValidMapping(opcodes []int) (err error) {

	for i := 0; i < len(opcodes); i++ {
		opcode := opcodes[i]
		if opcode != infoArray[i][ExpectedOpcode] {
			return errors.New("The opcode and size/cycle lookup are out of sync")
		}

		if opsize := opsizes[i]; opsize != infoArray[i][Size] {
			return errors.New("The test opcode size does not match the cpu opcode size")
		}

		if opcycle := opcycles[i]; opcycle != infoArray[i][Cycles] {
			return errors.New("The test opcode cycle time does not match the cpu opcode cycle time")
		}
	}

	return nil

}
