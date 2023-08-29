
// Copyright (c) 2017, Tom Thorogood
// Copyright (c) 2021, Carlo Alberto Ferraris
// All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

// +build amd64,!gccgo,!appengine

#include "textflag.h"


TEXT ·CASPair(SB),NOSPLIT,$8-48
	MOVQ pair+0(FP), BP
	MOVQ oldPtr+8(FP), AX
	MOVQ newPtr+16(FP), DX
	MOVQ oldMark+24(FP), BX
	MOVQ newMark+32(FP), CX
	LOCK
	CMPXCHG16B (BP)
	SETEQ swapped+40(FP)
	RET
