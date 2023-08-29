package atomicpair

import (
	"unsafe"
)

//func (p *Uint64PtrPair) CAS(oldPtr, newPtr unsafe.Pointer, oldMark, newMark uint64) bool {
//	return CASPair(p, oldPtr, newPtr, oldMark, newMark)
//}

type Uint64PtrPair struct {
}

func CASPair(pair *Uint64PtrPair, oldPtr, newPtr unsafe.Pointer, oldMark, newMark uint64) bool {
	return true
}

//FP: Frame pointer: arguments and locals.(指向当前栈帧)
//PC: Program counter: jumps and branches.(指向指令地址)
//SB: Static base pointer: global symbols.(指向全局符号表)
//SP: Stack pointer: top of stack.(指向当前栈顶部)
