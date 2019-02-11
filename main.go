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
// CONSTANTS
const colSize = 20
const rowStr = "|%-20s|%-20s|%-20s|\n"



func atMiddle(s string) string{
	return fmt.Sprintf("%[1]*s", -colSize, fmt.Sprintf("%[1]*s", (colSize + len(s))/2, s))
}

func main() {

	args := os.Args[1:]

	strIp, strMask, err := getIpAndMask()

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

	fmt.Printf("My computer:\nIP: %s, mac-adress: %s\n\n", strIp, getMacAddr())

	ips := getIps(strMask, strIp)

	for _, ip := range ips {
		ping(ip)
		node, err := arp(ip)
		if err == nil{
			fmt.Printf(rowStr, atMiddle(node.ip), atMiddle(node.mac), atMiddle(node.name))
		}
	}

}