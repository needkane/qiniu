package main

/*
#include "c_files/foo.h"
#include "c_files/foo.c"
#include "c_files/add.c"

typedef struct {
	int err;
	int width;
	int height;
} A;

int Generate(A *x)
{
	x->err = 0;
	x->width = 1;
	x->height = 2;
    return 0;
}

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
	CStructToGo()
}

func CStructToGo() {
	var a = new(C.A)
	C.Generate(a)

	var err, width, height int
	err = int(a.err)
	width = int(a.width)
	height = int(a.height)
	fmt.Println(a)
	fmt.Println("error:", err)
	fmt.Println("width:", width)
	fmt.Println("height:", height)
}
