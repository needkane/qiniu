package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"qbox.us/servend/account"
	"qbox.us/sstore"
)

const TestKeyHint = 106

var TestKey = []byte("qbox.pub.test")

func main() {
	// usage: ./eth <base64fh> <fsize>
	if len(os.Args) != 3 {
		fmt.Println("usage: ./eth <base64fh> <fsize>")
		return
	}

	fh, err := base64.URLEncoding.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal("invalid base64 encoding of fh")
	}

	fsize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("invalid fsize")
	}

	fi := &sstore.FhandleInfo{
		Fhandle:  fh,
		Fsize:    int64(fsize),
		Uid:      0,
		Deadline: int64(time.Now().UnixNano()) + (1e9 * 60 * 60 * 1), // 1小时
		KeyHint:  TestKeyHint,
		Utype:    account.USER_TYPE_ENTERPRISE,
		Public:   1,
	}
	eth := sstore.EncodeFhandle(fi, TestKey)
	fmt.Println("curl -v http://iovip.qbox.me/file/" + eth)
}
