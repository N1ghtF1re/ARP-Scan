package main

import (
	"errors"
	"fmt"
	"os"
)

type Node struct {
	name string
	ip string
	mac string
}

type IpAndMask struct{
	ip string
	mask string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseArgs(args []string) (string, string, error){
	newmask := ""
	newfile := ""

	if len(args) != 0 && len(args) % 2 != 0 {
		return "", "", errors.New("Invalid arguments")
	}
	for i := 0; i < len(args); i+=2 {
		switch args[i] {
		case "-m":
			newmask = args[i+1]
		case "-o":
			newfile = args[i+1]
		default:
			panic("Invalid argument: " + args[i])
		}

	}

	return newmask, newfile, nil
}

func main() {
	args := os.Args[1:]

	ipsAndMasks, err := getIpAndMask() // DEFAULT MASK
	check(err)

	file := os.Stdout // DEFAULT FILE

	argMask, argFile, err := parseArgs(args)
	check(err)

	// В аргументах указана маска
	if argMask != "" {
		for i := 0; i < len(ipsAndMasks); i++ {
			ipsAndMasks[i].mask = argMask // Подменяем везде маски
		}
	}

	// В аргументах указан выходной файл
	if argFile != "" {
		file, err = os.Create(argFile)
		check(err)
	}


	_, _ = fmt.Fprintf(file, "\nActive networks count: %d\n\n", len(ipsAndMasks))

	for i, el := range ipsAndMasks { // Перебор сетевых интерфейсов

		_, _ = fmt.Fprintf(file, "Network interface #%d\n\n", i+1)

		strMask := el.mask
		strIp := el.ip

		if !maskValid(strMask) {panic("Invalid mask")}

		_, _ = fmt.Fprintf(file, "My computer:\nIP: %s, mac-address: %s\n\n", strIp, getMacAddr())

		ips := getIps(strMask, strIp)

		drawHeader(file)
		for _, ip := range ips {
			ping(ip)
			node, err := arp(ip)
			if err == nil {
				drawRow(file, node)
			}
		}
		drawSplitter(file)
	}

}