package scanning

import (
	"github.com/go-ping/ping"
	"runtime"
	"time"
)

const (
	PingTimeout = 1200
)

func CheckIfUp(ip_to_check string, timeout int) (bool, error) { // Check if an IP is up using native golang
	var check bool
	pinger, err := ping.NewPinger(ip_to_check) // Create ping
	if err != nil {
		return false, err // Handle error
	}

	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
	pinger.Count = 2 // Set ping config
	pinger.Timeout = time.Duration(timeout) * time.Millisecond
	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketsRecv == 2 { // Check if ip has received both packets
			check = true
		} else {
			check = false
		}
	}

	err = pinger.Run()
	if err != nil {
		return false, err
	}

	return check, nil
}
