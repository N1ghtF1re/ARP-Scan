package main

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
)

/**
	Получение мак адреса компьютера
 */
func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

/**
	Получение массива пар айпи, маска всех доступных сетевых инетрфейсов
 */
func getIpAndMask() ([]IpAndMask, error) {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	var arr []IpAndMask

	if err != nil { return arr, err }

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIp, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {

			ip := networkIp.IP.String()
			mask := net.IP(networkIp.Mask).String()

			var node IpAndMask
			node.ip = ip
			node.mask = mask

			arr = append(arr, node)
		}
	}
	return arr, nil
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

func maskValid(ip string) bool{
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

/**
	Массив IP, соотвествующих текущей подсети
 */
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
