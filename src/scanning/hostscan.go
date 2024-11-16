package scanning

import (
	"encoding/binary"
	"fmt"
	"github.com/go-ping/ping"
	"net"
	"time"
)

func Hostscan(ip_range string, ping_timeout int) ([]string, error) { // Scan an ip range and return active ips in less than 30s
	var range_ips []string
	var active_ips []string

	_, ipv4Net, err := net.ParseCIDR(ip_range)
	if err != nil {
		return []string{""}, nil
	}

	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)
	finish := (start & mask) | (mask ^ 0xffffffff)

	for i := start; i <= finish; i++ { // Generate ips slice
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		range_ips = append(range_ips, ip.String())
	}

	// Iterate through ips to check
	for _, ip_to_scan := range range_ips {
		fmt.Println(ip_to_scan)
		pinger, err := ping.NewPinger(ip_to_scan) // Send ping to all ips and check if packet is up
		if err != nil {
			return []string{""}, err
		}

		pinger.SetPrivileged(true)
		pinger.Timeout = time.Duration(ping_timeout) * time.Millisecond
		pinger.Count = 1 // just send one packet
		pinger.OnFinish = func(stats *ping.Statistics) {
			if stats.PacketsRecv == 1 {
				active_ips = append(active_ips, ip_to_scan)
			}
		}

		err = pinger.Run()
		if err != nil {
			return []string{""}, err
		}
	}

	return active_ips, nil
}
