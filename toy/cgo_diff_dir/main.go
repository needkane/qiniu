package main

/*
#include "c_files/foo.h"
#include "c_files/foo.c"
#include "c_files/add.c"
*/
import "C"
import "unsafe"
import "fmt"

func Prin(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.fflush((*C.FILE)(C.stdout))
}
func main() {
	fmt.Println("rannum:%x\n", C.random())
	Prin("Hello CC")
	fmt.Println(C.Num)
	C.foo()
	a := C.int(1)
	b := C.int(2)
	value := C.f_add(a, b)
	fmt.Printf("a+b=%v\n", value)
}
