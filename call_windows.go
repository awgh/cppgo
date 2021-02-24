package cpp

import (
	"github.com/awgh/cppgo/asmcall/stdcall"
	"github.com/awgh/cppgo/asmcall/thiscall"
)

func (p ptr) cdeclcall(offset int, a ...uintptr) (uintptr, error) {
	return p.thiscall(offset, a...)
}

func (p ptr) stdcall(offset int, a ...uintptr) (uintptr, error) {
	addr := p.getaddr(offset)
	return stdcall.Call(addr, a...)
}

func (p ptr) thiscall(offset int, a ...uintptr) (uintptr, error) {
	addr := p.getaddr(offset)
	return thiscall.Call(addr, a...)
}
