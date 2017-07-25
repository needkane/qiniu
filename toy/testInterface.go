package main

import (
	"fmt"
)

type ReadWrite interface {
	Writer(buf []byte) (n int, err error)
	Reader(buf []byte) (n int, err error)
}
type Write interface {
	Writer(buf []byte) (n int, err error)
}
type File struct {
	f int
}

func (f *File) Writer(buf []byte) (n int, err error) {
	n = 0
	err = nil
	return
}
func (f *File) Reader(buf []byte) (n int, err error) {
	n = 1
	err = nil
	return
}
func main() {
	var f1 ReadWrite = new(File)
	var f2 Write = f1 //接口可以赋值，但是只能小接口等于大接口，而不能反过来（小的包含的方法不够）
	fmt.Println(f1, f2)

	//four query and judge
	var f3 Write
	if f4, ok := f3.(ReadWrite); ok {
		fmt.Println(f4, ok)
	} else {
		fmt.Println("else")
	}
	var f5 Write
	if f4, ok := f5.(*File); ok {
		fmt.Println(f4, ok)
	} else {
		fmt.Println("else")
	}
	var f6 Write = new(File)
	if f4, ok := f6.(ReadWrite); ok {
		fmt.Println(f4, ok)
	} else {
		fmt.Println("else")
	}
	var f7 Write = new(File)
	if f4, ok := f7.(*File); ok {
		fmt.Println(f4, ok)
	} else {
		fmt.Println("else")
	}
}
