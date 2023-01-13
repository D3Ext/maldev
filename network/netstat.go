package network

/*

References:
https://github.com/cakturk/go-netstat

*/

import (
  "fmt"
  "github.com/ryanuber/columnize"
  "github.com/cakturk/go-netstat/netstat"
)

type PortsInfo struct {
  Name      string
  Port      int
  Pid       int
  State     netstat.SkState
}

func Netstat() ([]*PortsInfo, error) {
  var info_to_return []*PortsInfo
  var port_process_name string
  var process_pid int

  tabs, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
    return s.State == netstat.Listen
  })

  if err != nil {
    return []*PortsInfo{}, err
  }

  for _, p := range tabs {

    if p.Process == nil { // Check if process has value or not (useful when not running as root)
      port_process_name = ""
      process_pid = 0
    } else {
      port_process_name = p.Process.Name
      process_pid = int(p.Process.Pid)
    }

    data := &PortsInfo{ // Parse data with struct
      Name: port_process_name,
      Port: int(p.LocalAddr.Port),
      Pid: process_pid,
      State: p.State,
    }

    info_to_return = append(info_to_return, data) // Append struct to slice
  }

  return info_to_return, nil // Finally return slice without errors
}

func FormatedNetstat() (string, error) { // Return a string in columns with correct display
  nets, err := Netstat()
  if err != nil {
    return "", err
  }

  var output []string
  var net_format string
  output = append(output, "Port || State || Process")
  output = append(output, "---- || ----- || -------")

  for _, e := range nets {
    if e.Name != "" { // If operation is not running as root/admin some PIDs can't be retrieved so it will show 0
      net_format = fmt.Sprintf("%v || %v || %v/%v", e.Port, e.State, e.Name, e.Pid)
    } else {
      net_format = fmt.Sprintf("%v || %v || %v", e.Port, e.State, e.Pid)
    }
    output = append(output, net_format)
  }

  config := columnize.DefaultConfig() // Create config to work with columns
  config.Delim = "||"
  config.Glue = "  "
  config.Prefix = ""
  config.Empty = ""
  config.NoTrim = false
  
  columns_out := columnize.Format(output, config) // Finally parse output line by line to create table format
  return columns_out, nil // Return string format
}

