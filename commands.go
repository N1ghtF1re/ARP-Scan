package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

/**
	Return true if node responds
 */
func ping(ip string) bool{
	//fmt.Printf("pinging %s\n", ip)
	cmd := exec.Command("ping", "-c", "1", "-w", "1", ip)

	bout, _ := cmd.Output()
	_ = cmd.Run()

	out := fmt.Sprintf("%s", bout)
	return !strings.Contains(out, "0 received")
}

func arp(ip string) (Node, error) {
	var node Node

	//fmt.Printf("ARP %s\n", ip)

	cmdArp := exec.Command("arp", "-a", ip)
	out, err := cmdArp.Output()
	if err != nil {return node, err}
	_ = cmdArp.Run()

	strout := strings.Replace(fmt.Sprintf("%s", out) , "\n", "", 1)


	if strings.Contains(strout, "<incomplete>") {
		return node, errors.New("Incomplete mac address ")
	}

	if strings.Contains(strout, "no match found") {
		return node, errors.New(strout)
	}

	cols := strings.Split(strout, " ")
	if cols[0] == "?" {
		node.name = "Undefined"
	} else {
		node.name = cols[0]
	}
	node.ip = ip
	node.mac = cols[3]

	return node, nil
}
