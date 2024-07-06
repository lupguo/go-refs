package snowflakex

import (
	"net"
	"sync"

	"golang.org/x/exp/rand"
)

var once sync.Once
var ipnum uint32

func IPToNumber() uint32 {
	if ipnum > 0 {
		return ipnum
	}
	once.Do(func() {
		ipnum = uint32(rand.Int31n(PrimeNum))
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				ip := ipNet.IP
				if ip.To4() != nil {
					ipnum = ipToUint32(ip.To4()) % PrimeNum
				}
				// if ip.To16() != nil {
				// 	return ipToUint32(ip.To16()), nil
				// }
			}
		}
	})

	return ipnum
}

func ipToUint32(ip net.IP) uint32 {
	ip = ip.To16()
	ipBytes := ip.To16()
	return (uint32(ipBytes[12]) << 24) |
		(uint32(ipBytes[13]) << 16) |
		(uint32(ipBytes[14]) << 8) |
		uint32(ipBytes[15])
}

const PrimeNum = 9973
