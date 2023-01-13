package network

/*

References:
https://pkg.go.dev/net
https://stackoverflow.com/questions/23529663/how-to-get-all-addresses-and-masks-from-local-interfaces-in-go

type Interface struct {
	Index        int          // positive integer that starts at one, zero is never used
	MTU          int          // maximum transmission unit
	Name         string       // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
}

*/

import (
  "net"
  "fmt"
  "strings"
  "errors"
)

func ListAllInterfaces() ([]string, error) {
  interfaces, err := net.Interfaces()
  if err != nil {
    return []string{""}, err
  }

  var interface_names []string
  for _, i := range interfaces {
    interface_names = append(interface_names, i.Name)
  }
  return interface_names, nil
}

func GetInterfaceInfo(interface_name string) (net.Interface, error) {
  interfaces, err := net.Interfaces()
  if err != nil {
    return interfaces[0], err
  }

  for _, inter := range interfaces {
    if interface_name == inter.Name {
      return inter, nil
    }
  }
  return interfaces[0], errors.New("ERROR: interface not found!")
}

func FormatInfo(interface_info net.Interface) (string) {
  var interfaces_out string
  interfaces_out += fmt.Sprintf("Index  : %v\n", interface_info.Index) // Simple information display
  interfaces_out += fmt.Sprintf("Name   : %v\n", interface_info.Name)
  interfaces_out += fmt.Sprintf("HWaddr : %v\n", interface_info.HardwareAddr)
  interfaces_out += fmt.Sprintf("MTU    : %v\n", interface_info.MTU)
  interfaces_out += fmt.Sprintf("Flags  : %v\n", interface_info.Flags)

  addrs, _ := interface_info.Addrs()
  for _, ipaddr := range addrs { // Loop because may have more than one address
    interfaces_out += fmt.Sprintf("Addr   : %v\n", ipaddr)
  }
  interfaces_out = strings.TrimSuffix(interfaces_out, "\n") // Remove trailing line

  return interfaces_out // Return formated info
}


