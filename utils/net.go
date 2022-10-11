package utils

import (
	"fmt"
	"net"
)

func NetGetLocalIP() string {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	ips := make([]string, 0)
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	return ips[0]
}

// NetServerIP get server ip from external
func NetServerIP() {

}

// NetMacAddrs get mac
func NetMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("fail to get net interfaces: " + err.Error())
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs

}
