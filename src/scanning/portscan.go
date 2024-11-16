package scanning

import (
	"net"
	"strconv"
	"time"
)

func PortscanAll(ip string) ([]int, error) {
	var ports_results []int
	for i := 1; i <= 65535; i++ {
		go func() {
			_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(i), time.Duration(400)*time.Millisecond)
			if err == nil {
				ports_results = append(ports_results, i)
			}
		}()

		if i%2 == 0 {
			time.Sleep(4 * time.Millisecond)
		} else {
			time.Sleep(2 * time.Millisecond)
		}
	}

	time.Sleep(800 * time.Millisecond)
	return ports_results, nil
}

func PortscanCommon(ip string) ([]int, error) {
	// Define array with common ports
	common_ports := []int{21, 22, 23, 24, 25, 43, 79, 80, 88, 135, 139, 143, 194, 389, 443, 445, 512, 513, 554, 593, 623, 636, 873, 993, 1080, 1433, 1521, 1522, 2049, 2222, 2375, 2376, 3000, 3001, 3128, 3268, 3269, 3306, 3389, 4000, 5000, 5555, 5601, 5984, 5985, 5986, 6379, 6984, 8000, 8080, 9200, 27017, 27018}

	var ports_results []int
	for _, port := range common_ports {
		go func() {
			_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Duration(500)*time.Millisecond)
			if err == nil {
				ports_results = append(ports_results, port)
			}
		}()

		time.Sleep(10 * time.Millisecond)
	}

	time.Sleep(800 * time.Millisecond)
	return ports_results, nil
}

func CheckOpenTcpPort(ip string, port int) bool {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Duration(500)*time.Millisecond)
	if err != nil {
		return false
	} else {
		return true
	}
}

func CheckOpenUdpPort(ip string, port int) bool {
	_, err := net.DialTimeout("udp", ip+":"+strconv.Itoa(port), time.Duration(750)*time.Millisecond)
	if err != nil {
		return false
	} else {
		return true
	}
}
