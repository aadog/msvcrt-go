package cmsvcrt

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
