package scanning

import (
  "github.com/go-ping/ping"
  "runtime"
  "time"
)

func CheckIfUp(ip_to_check string) (bool, error) { // Check if an IP is up using native golang
  var check bool
  pinger, err := ping.NewPinger(ip_to_check) // Create ping
  if err != nil {
    return false, err // Handle error
  }

  if runtime.GOOS == "windows" {
    pinger.SetPrivileged(true)
  }
  pinger.Count = 2 // Set ping config
  pinger.Timeout = 1200 * time.Millisecond
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

