TEXT ·Neg(SB), $0
    MOVQ 	x+0(FP), AX
	NEGQ    AX
	MOVQ    AX, ret+8(FP)
	RET
