package main

import (
	"day08/ex02/window"
)

// #include <stdlib.h>
import "C"

func main() {
	window.InitApplication()
	ptr := window.WindowCreate(600, 400, 300, 200, "School 21")
	defer C.free(ptr)
	window.MakeKeyAndOrderFront(ptr)
	window.RunApplication()
}
