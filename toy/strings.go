package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	s := "#EXTINF:10.080000,"
	fmt.Println(strings.Trim(strings.TrimPrefix(s, "#EXTINF:"), ","))
	s = "face2/20x20/ignore-error/1"
	i := strings.Index(s, "ignore")
	fmt.Println(s[:i-1])
	fmt.Println(s[:len(s)-1])
	fmt.Println(1<<63 - 1)

	s = "abc"
	s = strings.TrimSuffix(s, "d")
	fmt.Println(s)
	var url = "http://www.qiniu.com"
	fmt.Println(strings.HasPrefix("http://", url))
	switch s {
	case "abc":
		fmt.Println("-----case0----")
	case "ab":
		fmt.Println("------case1---------")
	default:
		fmt.Println("--------case2-----")
	}
	ff, _ := strconv.ParseFloat("2.34", 10)
	if ff > 2 {
		fmt.Println("3>2")
	}

	s3 := "seg01.ts"
	fmt.Println(s3[:len(s3)-1])

	fmt.Println("-------------------------")
	s = "rtmp://127.0.0.0.1/hub/stream2"
	StreamdUrl := strings.Replace(s, "rtmp://", "http://", 1)
	if i = strings.Index(StreamdUrl[7:], ":"); i == -1 {
		i = strings.Index(StreamdUrl[7:], "/") //set rtmp port is unnecessary ,rtmp default 1935
		fmt.Println(StreamdUrl[7:])
	}
	fmt.Println(StreamdUrl[7:])
	fmt.Println(i)
	StreamdUrl = StreamdUrl[:7+i] + ":2803" + "/stream/status?streamId="
	fmt.Println(StreamdUrl)
	fmt.Println("-------------#EXTINF:10.080000,#EXTINF:9.080000,------------")
	ss := []string{"#EXTINF:10.080000,", "#EXTINF:9.080000,"}
	var duration float64
	var num int
	for _, line := range ss {
		fmt.Println(line[8 : len(line)-1])
		d, err := strconv.ParseFloat((line[8 : len(line)-1]), 64)
		if err != nil {
			log.Println("can't  parsefloat")
		}
		duration += d
		num++
	}
	log.Println(int(duration / float64(num)))
	log.Println("------------------trim----------")
	s = "abc,"
	log.Println(strings.Trim(s, ","))
	s = ""
	log.Println(len([]byte(s)))
	i = strings.Index(s, ":")
	if i < 1 || len([]byte(s)) == i+1 {
		log.Println("i<1")
	}
	s = fmt.Sprintf("%v", ss)
	log.Println(s)
	p := fmt.Println
	p("----------------compare----------")
	rawQuery := "avthumb/mp4/wmImage/dsds/wmText/dsd"
	i1 := strings.Index(rawQuery, "wmImage")
	i2 := strings.Index(rawQuery, "wmText")
	if i1 == -1 || i1 > i2 {
		rawQuery = rawQuery[i2:]

	} else {
		rawQuery = rawQuery[i1:]
	}
	p(rawQuery)
	p(len("ufopavthumb/mp4/wmImage/aHR0cDovL25lZWRrYW5lLnFpbml1ZG4uY29tL3AucG5n"))
	s = "ufopavthumb/mp4/wmImage/aHR0cDovL25lZWRrYW5lLnFpbml1ZG4uY29tL3AucG5n"
	p(s[16:])
}
