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
	"errors"
	"fmt"
	"net"
	"strings"
)

type CustomInterface struct {
	Index        int
	MTU          int
	Name         string
	HardwareAddr net.HardwareAddr
	Flags        net.Flags
	Addrs        []string
}

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

func GetInterfaceInfo(interface_name string) (CustomInterface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return CustomInterface{}, err
	}

	for _, inter := range interfaces {
		if interface_name == inter.Name {

			var interface_addrs []string
			all_addrs, _ := inter.Addrs()
			for _, p := range all_addrs {
				interface_addrs = append(interface_addrs, p.String())
			}

			return CustomInterface{
				Index:        inter.Index,
				MTU:          inter.MTU,
				Name:         inter.Name,
				HardwareAddr: inter.HardwareAddr,
				Flags:        inter.Flags,
				Addrs:        interface_addrs,
			}, nil
		}
	}

	return CustomInterface{}, errors.New("ERROR: interface not found!")
}

func (i CustomInterface) GetMac() string {
	return i.HardwareAddr.String()
}

func (i CustomInterface) GetFlags() net.Flags {
	return i.Flags
}

func (i CustomInterface) GetName() string {
	return i.Name
}

func (i CustomInterface) GetMtu() int {
	return i.MTU
}

func (i CustomInterface) Format() string {
	var interfaces_out string
	interfaces_out += fmt.Sprintf("Index  : %v\n", i.Index) // Simple information display
	interfaces_out += fmt.Sprintf("Name   : %v\n", i.Name)
	interfaces_out += fmt.Sprintf("HWaddr : %v\n", i.HardwareAddr)
	interfaces_out += fmt.Sprintf("MTU    : %v\n", i.MTU)
	interfaces_out += fmt.Sprintf("Flags  : %v\n", i.Flags)

	for _, ipaddr := range i.Addrs { // Loop because may have more than one address
		interfaces_out += fmt.Sprintf("Addrs  : %v\n", ipaddr)
	}
	interfaces_out = strings.TrimSuffix(interfaces_out, "\n") // Remove trailing line

	return interfaces_out // Return formated info
}
