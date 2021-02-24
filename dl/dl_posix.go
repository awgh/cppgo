// +build !windows

package dl

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"

	"github.com/awgh/cppgo/asmcall/cdecl"
)

var (
	RTLD_NOW    = int(C.RTLD_NOW)
	RTLD_LAZY   = int(C.RTLD_LAZY)
	RTLD_GLOBAL = int(C.RTLD_GLOBAL)
	RTLD_LOCAL  = int(C.RTLD_LOCAL)
)

func open(filename string, flags ...int) *Library {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	flag := 0
	if len(flags) == 0 {
		flag = RTLD_NOW
	}
	for _, f := range flags {
		flag |= f
	}
	return &Library{p: C.dlopen(cfilename, C.int(flag))}
}

func (l Library) load(procname string) *Func {
	cprocname := C.CString(procname)
	defer C.free(unsafe.Pointer(cprocname))
	return &Func{p: uintptr(C.dlsym(l.p, cprocname))}
}

func (l Library) close() {
	C.dlclose(l.p)
}

func (f Func) stdcall(a ...uintptr) (uintptr, error) {
	return cdecl.Call(f.p, a...)
}
