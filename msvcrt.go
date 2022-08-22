package msvcrt

import (
	"github.com/aadog/msvcrt-go/cmsvcrt"
	"unsafe"
)

func CStrlen(s uintptr) int {
	return int(cmsvcrt.Strlen(s))
}
func WCStrlen(ws uintptr) int {
	return int(cmsvcrt.Wcslen(ws))
}

func MallocCString(s string) uintptr {
	temp := []byte(s)
	n := cmsvcrt.SIZE_T(len(temp) + 1)
	ptr := cmsvcrt.Malloc(n)
	cmsvcrt.Memset(ptr, 0, n)
	cmsvcrt.Memcpy(ptr, uintptr(unsafe.Pointer(&temp[0])), cmsvcrt.SIZE_T(len(temp)))
	return ptr
}

func Memcpy(dest, src uintptr, count uintptr) uintptr {
	return cmsvcrt.Memcpy(dest, src, count)
}
func Memset(src uintptr, val int, count uintptr) {
	cmsvcrt.Memset(src, val, count)
}

func Free(p uintptr) {
	cmsvcrt.Free(p)
}
