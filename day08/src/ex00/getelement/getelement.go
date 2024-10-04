package getelement

import (
	"errors"
	"unsafe"
)

func GetElement(arr []int, idx int) (int, error) {
	if arr == nil {
		return 0, errors.New("non-existent slice")
	}

	if len(arr) == 0 {
		return 0, errors.New("empty slice")
	}

	if idx < 0 {
		return 0, errors.New("negative index")
	}

	if idx >= len(arr) {
		return 0, errors.New("index out of range")
	}

	first := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(int(0))

	val := *(*int)(unsafe.Pointer(uintptr(first) + size*uintptr(idx)))

	return val, nil
}
