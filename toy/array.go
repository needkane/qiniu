package main

import (
	"fmt"		
	"strconv"
	"strings"
)

func main(){
	query := []string{"avthumb","hls","vn","1"}
	query[1] = "m3u8"
	query = append(query,"noDomain","1")
	fmt.Println(query)
	a,_ := strconv.ParseFloat("0.00",64)
	fmt.Println(a)
	b:="string"
	i:=strings.Index(b,"r")
	fmt.Println(b[:i])
	s:="seg11"
	i,_=strconv.Atoi(s[3:])
	new := fmt.Sprintf("%6d",i)
	new = strings.Replace(new," ","0",-1)
	fmt.Println(new,len(new))
	
}
