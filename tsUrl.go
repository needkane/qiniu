package main

import (
	"strings"
)

func main(){
	tsUrl:="http://127.0.0.1/dsdsads/00000.ts"
	i := strings.Index(tsUrl, "/")
	tsUrl = tsUrl[i+2:]
	println(tsUrl,i)
	j := strings.Index(tsUrl, "/")
	internalURL := "http://" + "ioDomain" + tsUrl[j:]
	println(internalURL)
}
