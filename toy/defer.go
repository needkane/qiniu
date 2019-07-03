package main

import (
	"fmt"
	"time"
)

type Stru struct {
	x int
}

func main() {
	f0()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
	fmt.Println(f7())
	fE()

	ch := make([]int, 1, 2)
	fmt.Println(len(ch), cap(ch))

	//var m map[int]int   // nil
	m := make(map[int]int) // init zero value
	var x Stru
	y := Stru{1}
	var z = new(Stru)
	var e []int
	var f = []int{}
	var g int
	fmt.Println(g, m == nil, x == Stru{}, y, z, z == nil, e == nil, f == nil, e, f)
	m[1] = 1
	a, ok := m[1]
	fmt.Println(m, a, ok)
	var m2 = map[int]bool{}
	m2[1] = true
	b, ok := m2[1]
	fmt.Println(m2, b, ok)
}

func f0() {
	defer fmt.Println("defer0")
	defer fmt.Println("defer1")
	fmt.Println("end")
}

func f1() (str string) {
	str0 := "defer f1"
	defer func() {
		//closure
		str = str0
	}()
	str0 = "second assign f1"
	str = str0
	return
}

func f2() (str string) {
	str0 := "f2"
	defer func(str0 string) {
		str = str0
	}(str0)
	str0 = "second assign f2"
	str = str0
	return
}
func f3() string {
	var str string
	str0 := "f3"
	defer func(str0 string) {
		str = str0
	}(str0)
	str0 = "second assign f3"
	str = str0
	return str
}

func f4() (r int) {

	defer func(r int) {
		r = 5
	}(r)
	return
}
func f5() (r []int) {
	defer func(r []int) {
		r = []int{5}
	}(r)
	return
}

func f6() []byte {
	var str []byte
	defer func() {
		str = []byte("defer f6")
		fmt.Printf("print defer f6: time(%d), addr(%p),content(%v) \n", time.Now().Unix(), &str, str)
	}()
	str = []byte("return f6")
	fmt.Printf("print return f6: time(%d), addr(%p),content(%v) \n", time.Now().Unix(), &str, str)
	str = []byte("return f6 again")
	fmt.Printf("print return f6 again: time(%d), addr(%p),content(%v) \n", time.Now().Unix(), &str, str)
	time.Sleep(time.Second)
	return str
}

func f7() (str []byte) {
	defer func() {
		str = []byte("defer f7")
		fmt.Printf("print defer f7: time(%d), addr(%p),content(%v) \n", time.Now().Unix(), &str, str)
	}()
	str = []byte("return f7")
	fmt.Printf("print return f7: time(%d), addr(%p),content(%v) \n", time.Now().Unix(), &str, str)
	str = []byte("return f7 again")
	fmt.Printf("print return f7 again: time(%d), addr(%p),content(%v) \n", time.Now().Unix(), &str, str)
	time.Sleep(time.Second)
	return
}
func fE() {
	var str string
	defer fmt.Println("defer fE0")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover: ", r)
		}
	}()
	defer func() {
		str = "first defer"
	}()
	panic(str)
	defer fmt.Println("defer fE1")
	fmt.Println("fE end")
}
