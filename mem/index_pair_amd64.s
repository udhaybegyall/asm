// Code generated by command: go run index_pair_asm.go -pkg mem -out ../mem/index_pair_amd64.s -stubs ../mem/index_pair_amd64.go. DO NOT EDIT.

#include "textflag.h"

// func indexPair1(b []byte) int
TEXT ·indexPair1(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), AX
	MOVQ b_len+8(FP), CX
	CMPQ CX, $0x01
	JBE  done
	MOVQ AX, DX
	MOVQ AX, BX
	ADDQ CX, BX

loop:
	MOVB (DX), SI
	MOVB 1(DX), DI
	CMPB SI, DI
	JE   found
	INCQ DX
	CMPQ DX, BX
	JNE  loop

done:
	MOVQ CX, ret+24(FP)
	RET

found:
	// The delta between the base pointer and how far we advanced is the index of the pair.
	SUBQ AX, DX
	MOVQ DX, ret+24(FP)
	RET
