package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Print the hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s\n\n", hostname)

	// Get a list of all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces:", err)
		os.Exit(1)
	}

	// Print table header
	fmt.Printf("%-20s %-10s %-20s %-30s\n", "Interface", "MTU", "MAC Address", "IP Address")
	fmt.Println(strings.Repeat("-", 80))

	for _, iface := range interfaces {
		addresses, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error getting addresses:", err)
			continue
		}

		// Check if there are any addresses
		if len(addresses) > 0 {
			// Print first address along with interface details
			fmt.Printf("%-20s %-10d %-20s %-30s\n", iface.Name, iface.MTU, iface.HardwareAddr, addresses[0].String())

			// Print remaining addresses
			for _, addr := range addresses[1:] {
				fmt.Printf("%-20s %-10s %-20s %-30s\n", "", "", "", addr.String())
			}
		} else {
			// Print interface details with no IP address
			fmt.Printf("%-20s %-10d %-20s %-30s\n", iface.Name, iface.MTU, iface.HardwareAddr, "No IP Address")
		}
	}
}
