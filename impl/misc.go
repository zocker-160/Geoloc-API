package impl

import (
	"errors"
	"net"
)


func IpToDecimal(ip string) (uint32, error) {
	parsedIP := net.ParseIP(ip)
	ipBytes := parsedIP.To4()
	if ipBytes == nil {
		return 0, errors.New("invalid IP address")
	}

	ipDecimal := 
		uint32(ipBytes[0]) << 24 | 
		uint32(ipBytes[1]) << 16 |
		uint32(ipBytes[2]) << 8 | 
		uint32(ipBytes[3])

	return ipDecimal, nil
}