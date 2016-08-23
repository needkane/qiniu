package main

import (
	"io/ioutil"
	"os"

	"qiniupkg.com/x/log.v7"
)

func main() {
	files, err := ioutil.ReadDir("/home/qboxtest/Desktop/qiniu/")
	if err != nil {
		log.Warn(err)
	}
	fi, _ := os.Stat("httpServerGet.go")
	modtime := fi.ModTime()
	for _, file := range files {
		if err != nil {
			log.Warn(err)
			break
		}
		if file.IsDir() {
			continue
		} else {
			fi, err := os.Stat(file.Name())
			if err != nil {
				log.Warn(err)
				break
			}
			if fi.ModTime().Unix() > modtime.Unix() {
				log.Println(file.Name())
				log.Println(fi.ModTime())

			}
		}
	}
}
