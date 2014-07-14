package main

import (
	"fmt"
)

func main() {
	var j int = 5
	a := func() func() {
		var i int = 10 //i's safety is ensured
		return func() {
			fmt.Printf("i ,j:%d,%d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()

	var v1 int = 3
	var v2 string = "kane"
	var v3 float64 = 5.4
	var v4 bool = true
	myFunc(v1, v2, v3, v4)
}
func myFunc(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println("int type ")
		case string:
			fmt.Println("string type")
		case float64:
			fmt.Println("float64 type")
		default:
			fmt.Println("unknown type")
		}
	}
}
