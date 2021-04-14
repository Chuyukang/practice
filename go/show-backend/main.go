package main

import (
	"fmt"
	"log"
	"net"
)

func GetIps() []net.IP{

	var ifaces []net.Interface
	var err error
	ifaces, err = net.Interfaces()
	if err != nil {
		log.Printf("Get Interfaces Error\n")
		return nil
	}

	var ipList []net.IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Printf("Get Addrs Error!\n")
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ipList = append(ipList, ip)
		}
	}

	return ipList
}

func main() {
	ipList := GetIps()
	if ipList==nil {
		fmt.Printf("No IP found\n")
	} else {
		for i, ip := range ipList {
			if ip.To4() == nil {
				continue
			}
			if ip.IsLoopback() {
				continue
			}
			fmt.Printf("ip[%d]: %s\n", i, ip)
		}
	}
}