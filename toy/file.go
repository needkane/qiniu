package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "abc"
	fmt.Printf("----%v\n", reflect.TypeOf(len(s)).Kind())
	fmt.Println(reflect.DeepEqual(int64(3), len(s)))
}
