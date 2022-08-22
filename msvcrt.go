package msvcrt

import "syscall"

type Pointer = uintptr
type SIZE_T = uintptr

var msvcrtdll = syscall.NewLazyDLL("msvcrt.dll")

var (
	memcpy = msvcrtdll.NewProc("memcpy")
	memset = msvcrtdll.NewProc("memset")
	malloc = msvcrtdll.NewProc("malloc")
	free   = msvcrtdll.NewProc("free")
	wcslen = msvcrtdll.NewProc("wcslen")
	strlen = msvcrtdll.NewProc("strlen")
)

func Memset(src Pointer, val int, count SIZE_T) {
	memset.Call(src, uintptr(val), count)
}

func Memcpy(dest, src Pointer, count SIZE_T) Pointer {
	r, _, _ := memcpy.Call(dest, src, count)
	return r
}

func Malloc(size SIZE_T) Pointer {
	r, _, _ := malloc.Call(size)
	return r
}
func Free(size SIZE_T) {
	free.Call(size)
}

func Wcslen(ws Pointer) SIZE_T {
	r, _, _ := wcslen.Call(ws)
	return r
}
func Strlen(s Pointer) SIZE_T {
	r, _, _ := strlen.Call(s)
	return r
}
