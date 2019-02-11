package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/j-keck/arping"
	"net"
	"os"
	"strconv"
)
func resolveHostIpAndMask() (string, string, error) {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil { return "", "", err }

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIp, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {

			ip := networkIp.IP.String()
			mask := net.IP(networkIp.Mask).String()

			return ip, mask, nil
		}
	}
	return "", "", errors.New("IP Not Found")
}

func ip2Long(ip string) uint32 {
	var long uint32
	_ = binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

func long2ip(ipInt int64) string {
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func isMaskValid(ip string) bool{
	mask := net.IPMask(net.ParseIP(ip).To4())
	ones, bits := mask.Size()
	if bits == 0 {
		return false
	}
	if ones == bits {
		return false
	}
	return true
}

func getIps(strMask, strIp string) []string {
	var arr []string
	maskInt := ip2Long(strMask)
	ipInt := ip2Long(strIp)

	broadcastInt := ipInt | ^maskInt
	startIp := ipInt & maskInt

	for ip:=startIp; ip < broadcastInt; ip++ {
		arr = append(arr,  long2ip(int64(ip)))
	}

	return arr
}

func main() {
	args := os.Args[1:]

	strIp, strMask, err := resolveHostIpAndMask()

	switch len(args) {
		case 1: {

		}
		case 2: {

		}
	}



	if !isMaskValid(strMask) {
		fmt.Println("Invalid mask")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	ips := getIps(strMask, strIp)

	for _, ip := range ips {
		dstIP := net.ParseIP(ip)
		if hwAddr, _, err := arping.Ping(dstIP); err == nil {
			fmt.Printf("%s : %s\n", dstIP, hwAddr)
		}
	}




}