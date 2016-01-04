package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	as := ""
	for _, char := range []rune{'A', 0xE6, 0346, 330, '\xE6', '\u00E6'} {
		fmt.Printf("[0x%x '%c']", char, char)
		as += string(char)
	}
	fmt.Println(as)
	line := "Θ￡ 123 Θ ￡ ￡Θ" //Θ and ￡  not only one char
	i := strings.Index(line, " ")
	i2 := strings.IndexFunc(line, unicode.IsSpace)
	firstword := line[:i]
	j := strings.LastIndex(line, " ")
	j2 := strings.LastIndexFunc(line, unicode.IsSpace)
	lastword := line[j+1:]
	item, size := utf8.DecodeRuneInString(line[j+1:]) //返回第一个字符(可能超过一般的ascii码)和字节数
	fmt.Println(i, ",", i2, ",", firstword, ",", j, ",", j2, lastword, item, size)
	bs:="{\"error\":\"thumbnail's arg is out of range: 650001456812x\"}"
	i=strings.Index(bs,":")
	j=strings.LastIndex(bs,"}")
	fmt.Println(bs[i+2:j-1])
	

	var iss []int{1,2,3}
	fmt.Println()
}
