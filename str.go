package msvcrt

import (
	"syscall"
	"unsafe"
)

func BytesToCBytes(b []byte) uintptr {
	if len(b) == 0 {
		return 0
	}
	return uintptr(BytesToCBytesPtr(b))
}

func BytesToCBytesPtr(b []byte) unsafe.Pointer {
	if len(b) == 0 {
		return nil
	}
	return unsafe.Pointer(&b[0])
}

func StringToCUTF8String(s string) uintptr {
	return uintptr(StringToCUTF8StringPtr(s))
}

func StringToCUTF8StringPtr(s string) unsafe.Pointer {
	temp := []byte(s)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return unsafe.Pointer(&utf8StrArr[0])
}

func StringToCUTF16String(s string) uintptr {
	return uintptr(StringToCUTF16StringPtr(s))
}

func StringToCUTF16StringPtr(s string) unsafe.Pointer {
	p, _ := syscall.UTF16PtrFromString(s)
	return unsafe.Pointer(p)
}

func CUTF16PtrToString(p uintptr) string {
	l := WCStrlen(p)
	if l == 0 {
		return ""
	}
	buff := make([]uint16, l)

	Memcpy(uintptr(unsafe.Pointer(&buff[0])), p, uintptr(l*2))
	return syscall.UTF16ToString(buff)
}

func CUtf8ToString(src uintptr) string {
	strLen := int(CStrlen(src))
	if strLen == 0 {
		return ""
	}
	str := make([]uint8, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*uint8)(unsafe.Pointer(src + uintptr(i)))
	}
	return string(str)
}

func CBytesToBytes(src uintptr, n int) []byte {
	strLen := int(CStrlen(src))
	if strLen == 0 {
		return nil
	}
	str := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*byte)(unsafe.Pointer(src + uintptr(i)))
	}
	return str
}
