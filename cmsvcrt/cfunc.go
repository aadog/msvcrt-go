package cmsvcrt

func Memcpy(dest, src Pointer, count SIZE_T) Pointer {
	r, _, _ := memcpy.Call(dest, src, count)
	return r
}
func Memset(src Pointer, val int, count SIZE_T) {
	memset.Call(src, uintptr(val), count)
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
