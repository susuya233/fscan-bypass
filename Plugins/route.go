package Plugins

import (
	"net"
	"strings"
)

func Subip() (bool, bool, bool, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false, false, false, err
	}
	var is172, is10, is192 bool
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip := ipNet.IP.String()
				if strings.HasPrefix(ip, "192.168.") {
					is192 = true
				}
				if strings.HasPrefix(ip, "10.") {
					is10 = true
				}
				if strings.HasPrefix(ip, "172.") {
					is172 = true
				}
			}
		}
	}
	return is172, is10, is192, nil
}
