package main
import (
		"fmt"
		"path"
)

func main(){
	a:=path.Join("http://","a.com","/dsds")
	fmt.Println(a)
	ss:=[]string{"qiniu","127.0.0.1"}
	  fmt.Println(ss[1:])
}
