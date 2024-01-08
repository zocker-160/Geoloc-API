package main

import (
	"errors"
	"net"
)

func ipToDecimal(ip string) (uint32, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return 0, errors.New("invalid IP address")
	}

	ipBytes := parsedIP.To4()
	if ipBytes == nil {
		return 0, errors.New("invalid IPv4 address")
	}

	ipDecimal :=
		uint32(ipBytes[0]) << 24 |
		uint32(ipBytes[1]) << 16 |
		uint32(ipBytes[2]) << 8 |
		uint32(ipBytes[3])

	return ipDecimal, nil
}
