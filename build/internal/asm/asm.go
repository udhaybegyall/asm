package asm

import (
	"encoding/binary"

	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/operand"
)

func ConstBytes(name string, data []byte) operand.Mem {
	m := GLOBL(name, RODATA|NOPTR)

	switch {
	case len(data)%8 == 0:
		constBytes8(0, data)

	case len(data)%4 == 0:
		constBytes4(0, data)

	default:
		i := (len(data) / 8) * 8
		constBytes8(0, data[:i])
		constBytes1(i, data[i:])
	}

	return m
}

func constBytes8(offset int, data []byte) {
	for i := 0; i < len(data); i += 8 {
		DATA(offset+i, operand.U64(binary.LittleEndian.Uint64(data[i:i+8])))
	}
}

func constBytes4(offset int, data []byte) {
	for i := 0; i < len(data); i += 4 {
		DATA(offset+i, operand.U32(binary.LittleEndian.Uint32(data[i:i+4])))
	}
}

func constBytes1(offset int, data []byte) {
	for i, b := range data {
		DATA(offset+i, operand.U8(b))
	}
}