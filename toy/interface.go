package main

import (
	"log"
	"net"
)

func main() {
	in, err := net.Interfaces()
	log.Println(in, "----", err)
	for _, interf := range in {
		addrs, err := interf.Addrs()
		if err != nil {
			log.Println("addr err: ", err)
			break
		}
		for _, ip := range addrs {
			ipParsed, _, err := net.ParseCIDR(ip.String())
			if err != nil {
				log.Println("parse err: ", err)
				break
			}
			log.Println("-----ipParse: ", ipParsed)
			ip4 := ipParsed.To4()
			if ip4 == nil {
				log.Println("ip4 is nil")
			}
			log.Println("ipstring: ", ip4.String())
		}
	}

	log.Println(1 << 0)
}
