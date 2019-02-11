package main

import (
	"fmt"
	"os"
)

type Node struct {
	name string
	ip string
	mac string
}




func main() {

	args := os.Args[1:]

	strIp, strMask, err := getIpAndMask()

	file := os.Stdout

	switch len(args) {
		case 1: {

		}
		case 2: {

		}
	}


	if !maskValid(strMask) {
		fmt.Println("Invalid mask")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = fmt.Fprintf(file, "My computer:\nIP: %s, mac-adress: %s\n\n", strIp, getMacAddr())

	ips := getIps(strMask, strIp)

	drawHeader(file)
	for _, ip := range ips {
		ping(ip)
		node, err := arp(ip)
		if err == nil{
			drawRow(file, node)
		}
	}
	drawSplitter(file)

}