package core

import "errors"

const (
	ExpectedOpcode = iota
	Size
	Cycles
	BoundaryCycles
	AddressMode
)

const (
	None = iota
	Imm
	Zp
	ZpX
	ZpY //Not yet implemented
	Abs
	AbsX
	AbsY
	Ind //Not yet implemented
	IndX
	IndY
	Rel //Not yet implemented
	Imp //Not yet implemented
)

var infoArray = [][5]int{
	{brk, 1, 7, 0, Imp},
	{oraIndX, 2, 6, 0, IndX},
	{fe02, 0, 0, 0, 0},
	{fe03, 0, 0, 0, 0},
	{fe04, 0, 0, 0, 0},
	{oraZp, 2, 3, 0, Zp},
	{aslZp, 2, 5, 0, Zp},
	{fe07, 0, 0, 0, 0},

	{php, 1, 3, 0, Imp},
	{oraImm, 2, 2, 1, Imm}, //One Extra cycle to be added if accessing across page boundary
	{aslAcc, 1, 2, 0, Imp},
	{fe0b, 0, 0, 0, 0},
	{fe0c, 0, 0, 0, 0},
	{oraAbs, 3, 4, 0, Abs},
	{aslAbs, 3, 6, 0, Abs},
	{fe0f, 0, 0, 0, 0},

	{bpl, 2, 2, 1, Rel}, //+1 Cycle if branch to same page; +2 if branch to different page
	{oraIndY, 2, 5, 0, IndY},
	{fe12, 0, 0, 0, 0},
	{fe13, 0, 0, 0, 0},
	{fe14, 0, 0, 0, 0},
	{oraZpX, 2, 4, 0, ZpX},
	{aslZpX, 2, 6, 0, ZpX},
	{fe17, 0, 0, 0, 0},

	{clc, 1, 2, 0, Imp},
	{oraAbsY, 3, 4, 1, AbsY}, //+1 cycle for crossing page boundary
	{fe1a, 0, 0, 0, 0},
	{fe1b, 0, 0, 0, 0},
	{fe1c, 0, 0, 0, 0},
	{oraAbsX, 3, 4, 1, AbsX}, //+1 cycle for crossing page boundary
	{aslAbsX, 3, 7, 0, AbsX},
	{fe1f, 0, 0, 0, 0},

	{jsr, 3, 6, 0, Abs},
	{andIndX, 2, 6, 0, IndX},
	{fe22, 0, 0, 0, 0},
	{fe23, 0, 0, 0, 0},
	{bitZp, 2, 3, 0, Zp},
	{andZp, 2, 3, 0, Zp},
	{rolZp, 2, 5, 0, Zp},
	{fe27, 0, 0, 0, 0},

	{plp, 1, 4, 0, Imp},
	{andImm, 2, 2, 0, Imm},
	{rolAcc, 1, 2, 0, Imp},
	{fe2b, 0, 0, 0, 0},
	{bitAbs, 3, 4, 0, Abs},
	{andAbs, 3, 4, 0, Abs},
	{rolAbs, 3, 6, 0, Abs},
	{fe2f, 0, 0, 0, 0},

	{bmi, 2, 2, 1, Rel}, //+1 if branch to same page, +2 otherwise
	{andIndY, 2, 5, 0, IndY},
	{fe32, 0, 0, 0, 0},
	{fe33, 0, 0, 0, 0},
	{fe34, 0, 0, 0, 0},
	{andZpX, 2, 4, 0, ZpX},
	{rolZpX, 2, 6, 0, ZpX},
	{fe37, 0, 0, 0, 0},

	{sec, 1, 2, 0, Imp},
	{andAbsY, 3, 4, 1, AbsY}, //+1 if page boundary is crossed.
	{fe3a, 0, 0, 0, 0},
	{fe3b, 0, 0, 0, 0},
	{fe3c, 0, 0, 0, 0},
	{andAbsX, 3, 4, 1, AbsX}, //+1 if page boundary is crossed
	{rolAbsX, 3, 7, 0, AbsX},
	{fe3f, 0, 0, 0, 0},

	{rti, 1, 6, 0, Imp},
	{eorIndX, 2, 6, 0, IndX},
	{fe42, 0, 0, 0, 0},
	{fe43, 0, 0, 0, 0},
	{fe44, 0, 0, 0, 0},
	{eorZp, 2, 3, 0, Zp},
	{lsrZp, 2, 5, 0, Zp},
	{fe47, 0, 0, 0, 0},

	{pha, 1, 3, 0, Imp},
	{eorImm, 2, 2, 0, Imm},
	{lsrAcc, 1, 2, 0, Imp},
	{fe4b, 0, 0, 0, 0},
	{jmpAbs, 3, 3, 0, Abs},
	{eorAbs, 3, 4, 0, Abs},
	{lsrAbs, 3, 6, 0, Abs},
	{fe4f, 0, 0, 0, 0},

	{bvc, 2, 2, 1, Rel},      // +1 if branch to same page, +2 otherwise
	{eorIndY, 2, 5, 1, IndY}, // +1 if page boundary is crossed
	{fe52, 0, 0, 0, 0},
	{fe53, 0, 0, 0, 0},
	{fe54, 0, 0, 0, 0},
	{eorZpX, 2, 4, 0, ZpX},
	{lsrZpX, 2, 6, 0, ZpX},
	{fe57, 0, 0, 0, 0},

	{cli, 1, 2, 0, Imp},
	{eorAbsY, 3, 4, 1, AbsY}, // +1 if page boundary crossed
	{fe5a, 0, 0, 0, 0},
	{fe5b, 0, 0, 0, 0},
	{fe5c, 0, 0, 0, 0},
	{eorAbsX, 3, 4, 1, AbsX}, // +1 if page boundary crossed
	{lsrAbsX, 3, 7, 0, AbsX},
	{fe5f, 0, 0, 0, 0},

	{rts, 1, 6, 0, Imp},
	{adcIndX, 2, 6, 0, IndX},
	{fe62, 0, 0, 0, 0},
	{fe63, 0, 0, 0, 0},
	{fe64, 0, 0, 0, 0},
	{adcZp, 2, 3, 0, Zp},
	{rorZp, 2, 5, 0, Zp},
	{fe67, 0, 0, 0, 0},

	{pla, 1, 4, 0, Imp},
	{adcImm, 2, 2, 0, Imm},
	{rorAcc, 1, 2, 0, Imp},
	{fe6b, 0, 0, 0, 0},
	{jmpInd, 3, 5, 0, Ind},
	{adcAbs, 3, 4, 0, Abs},
	{rorAbs, 3, 6, 0, Abs},
	{fe6f, 0, 0, 0, 0},

	{bvs, 2, 2, 1, Rel}, // +1 if branch to same a page, +2 otherwise
	{adcIndY, 2, 6, 0, IndY},
	{fe72, 0, 0, 0, 0},
	{fe73, 0, 0, 0, 0},
	{fe74, 0, 0, 0, 0},
	{adcZpX, 2, 4, 0, ZpX},
	{rorZpX, 2, 6, 0, ZpX},
	{fe77, 0, 0, 0, 0},

	{sei, 1, 2, 0, 0},
	{adcAbsY, 3, 4, 1, AbsY}, // +1 if page boundary is crossed
	{fe7a, 0, 0, 0, 0},
	{fe7b, 0, 0, 0, 0},
	{fe7c, 0, 0, 0, 0},
	{adcAbsX, 3, 4, 1, AbsX}, //+1 if page boundary crossed
	{rorAbsX, 3, 7, 0, AbsX},
	{fe7f, 0, 0, 0, 0},

	{fe80, 0, 0, 0, 0},
	{staIndX, 2, 6, 0, IndX},
	{fe82, 0, 0, 0, 0},
	{fe83, 0, 0, 0, 0},
	{styZp, 2, 3, 0, Zp},
	{staZp, 2, 3, 0, Zp},
	{stxZp, 2, 3, 0, Zp},
	{fe87, 0, 0, 0, 0},

	{dey, 1, 2, 0, Imp},
	{fe89, 0, 0, 0, 0},
	{txa, 1, 2, 0, Imp},
	{fe8b, 0, 0, 0, 0},
	{styAbs, 3, 4, 0, Abs},
	{staAbs, 3, 4, 0, Abs},
	{stxAbs, 3, 4, 0, Abs},
	{fe8f, 0, 0, 0, 0},

	{bcc, 2, 2, 1, Rel}, // +1 if branch to same page, +2 otherwise
	{staIndY, 2, 6, 0, IndY},
	{fe92, 0, 0, 0, 0},
	{fe93, 0, 0, 0, 0},
	{styZpX, 2, 4, 0, ZpX},
	{staZpX, 2, 4, 0, ZpX},
	{stxZpY, 2, 4, 0, ZpY},
	{fe97, 0, 0, 0, 0},

	{tya, 1, 2, 0, Imp},
	{staAbsY, 3, 5, 0, AbsY},
	{txs, 1, 2, 0, Imp},
	{fe9b, 0, 0, 0, 0},
	{fe9c, 0, 0, 0, 0},
	{staAbsX, 3, 5, 0, AbsX},
	{fe9e, 0, 0, 0, 0},
	{fe9f, 0, 0, 0, 0},

	{ldyImm, 2, 2, 0, Imm},
	{ldaIndX, 2, 6, 0, IndX},
	{ldxImm, 2, 2, 0, Imm},
	{fea3, 0, 0, 0, 0},
	{ldyZp, 2, 3, 0, Zp},
	{ldaZp, 2, 3, 0, Zp},
	{ldxZp, 2, 3, 0, Zp},
	{fea7, 0, 0, 0, 0},

	{tay, 1, 2, 0, Imp},
	{ldaImm, 2, 2, 0, Imm},
	{tax, 1, 2, 0, Imp},
	{feab, 0, 0, 0, 0},
	{ldyAbs, 3, 4, 0, Abs},
	{ldaAbs, 3, 4, 0, Abs},
	{ldxAbs, 3, 4, 0, Abs},
	{feaf, 0, 0, 0, 0},

	{bcs, 2, 2, 1, Rel},      // +1 if branch to same page, +2 otherwise
	{ldaIndY, 2, 5, 1, IndY}, // +1 if page boundary is crossed
	{feb2, 0, 0, 0, 0},
	{feb3, 0, 0, 0, 0},
	{ldyZpX, 2, 4, 0, ZpX},
	{ldaZpX, 2, 4, 0, ZpX},
	{ldxZpY, 2, 4, 0, ZpY},
	{feb7, 0, 0, 0, 0},

	{clv, 1, 2, 0, Imp},
	{ldaAbsY, 3, 4, 1, AbsY}, // +1 if page boundary is crossed.
	{tsx, 1, 2, 0, Imp},
	{febb, 0, 0, 0, 0},
	{ldyAbsX, 3, 4, 1, AbsX}, // +1 if page boundary is crossed.
	{ldaAbsX, 3, 4, 1, AbsX}, // +1 if page boundary is crossed.
	{ldxAbsY, 3, 4, 1, AbsY}, // +1 if page boundary is crossed.
	{febf, 0, 0, 0, 0},

	{cpyImm, 2, 2, 0, Imm},
	{cmpIndX, 2, 6, 0, IndX},
	{fec2, 0, 0, 0, 0},
	{fec3, 0, 0, 0, 0},
	{cpyZp, 2, 3, 0, Zp},
	{cmpZp, 2, 4, 0, Zp},
	{decZp, 2, 5, 0, Zp},
	{fec7, 0, 0, 0, 0},

	{iny, 1, 2, 0, Imp},
	{cmpImm, 2, 2, 0, Imm},
	{dex, 1, 2, 0, Imp},
	{fecb, 0, 0, 0, 0},
	{cpyAbs, 3, 4, 0, Abs},
	{cmpAbs, 3, 4, 0, Abs},
	{decAbs, 3, 6, 0, Abs},
	{fecf, 0, 0, 0, 0},

	{bne, 2, 2, 1, Rel}, // +1 if branch to same page, +2 otherwise
	{cmpIndY, 2, 6, 0, IndY},
	{fed2, 0, 0, 0, 0},
	{fed3, 0, 0, 0, 0},
	{fed4, 0, 0, 0, 0},
	{cmpZpX, 2, 4, 0, ZpX},
	{decZpX, 2, 6, 0, ZpX},
	{fed7, 0, 0, 0, 0},

	{cld, 1, 2, 0, Imp},
	{cmpAbsY, 3, 4, 1, AbsY}, // +1 if page boundary is crossed
	{feda, 0, 0, 0, 0},
	{fedb, 0, 0, 0, 0},
	{fedc, 0, 0, 0, 0},
	{cmpAbsX, 3, 4, 1, AbsX}, // +1 if page boundary crossed
	{decAbsX, 3, 7, 0, AbsX},
	{fedf, 0, 0, 0, 0},

	{cpxImm, 2, 2, 0, Imm},
	{sbcIndX, 2, 6, 0, IndX},
	{fee2, 0, 0, 0, 0},
	{fee3, 0, 0, 0, 0},
	{cpxZp, 2, 3, 0, Zp},
	{sbcZp, 2, 3, 0, Zp},
	{incZp, 2, 5, 0, Zp},
	{fee7, 0, 0, 0, 0},

	{inx, 1, 2, 0, Imp},
	{sbcImm, 2, 2, 0, Imm},
	{nop, 1, 2, 0, Imp},
	{feeb, 0, 0, 0, 0},
	{cpxAbs, 3, 4, 0, Abs},
	{sbcAbs, 3, 4, 0, Abs},
	{incAbs, 3, 6, 0, Abs},
	{feef, 0, 0, 0, 0},

	{beq, 2, 2, 1, Rel},      // +1 for branch to same page, +2 otherwise
	{sbcIndY, 2, 5, 1, IndY}, // +1 if page boundary is crossed.
	{fef2, 0, 0, 0, 0},
	{fef3, 0, 0, 0, 0},
	{fef4, 0, 0, 0, 0},
	{sbcZpX, 2, 4, 0, ZpX},
	{incZpX, 2, 6, 0, ZpX},
	{fef7, 0, 0, 0, 0},

	{sed, 1, 2, 0, Imp},
	{sbcAbsY, 3, 4, 1, AbsY}, // +1 if page boundary crossed
	{fefa, 0, 0, 0, 0},
	{fefb, 0, 0, 0, 0},
	{fefc, 0, 0, 0, 0},
	{sbcAbsX, 3, 4, 1, AbsX}, // +1 if page boundary crossed.
	{incAbsX, 3, 7, 0, AbsX},
	{feff, 0, 0, 0, 0},
}

func InstInfo(opcode int) (size, cycles int, err error) {

	data := infoArray[opcode]

	if opcode == data[ExpectedOpcode] {
		return data[Size], data[Cycles], nil
	}

	return 0, 0, errors.New("The opcode and size/cycle lookup are out of sync")

}
