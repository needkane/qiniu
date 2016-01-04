package main

import(
	"fmt"
	"strings"
	"strconv"
)


func main(){
	s:="#EXTINF:10.080000,"
	fmt.Println(strings.Trim(strings.TrimPrefix(s,"#EXTINF:"),","))
	s="face2/20x20/ignore-error/1"
	i:=strings.Index(s,"ignore")
	fmt.Println(s[:i-1])
	fmt.Println(s[:len(s)-1])
	fmt.Println(1<<63 -1)
	
	s="abc"
	s=strings.TrimSuffix(s,"d")
	fmt.Println(s)
	var url="http://www.qiniu.com"
 fmt.Println(	strings.HasPrefix("http://", url)  )
	switch s{
		case "abc":
		fmt.Println("-----case0----")
		case "ab":
		fmt.Println("------case1---------")
		default:
		fmt.Println("--------case2-----")
	}
	ff,_:=strconv.ParseFloat("2.34",10)
	if ff>2{
		fmt.Println("3>2")
	}

s3:="seg01.ts"
	   fmt.Println(s3[:len(s3)-1])
}
