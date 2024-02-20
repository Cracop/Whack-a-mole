package main

import (
	"fmt"
	"net"
)

func receiveUDP(c *CONNECTION) {
	// Multicast address and port
	multicastAddr := "224.0.0.1:9999"

	// Resolve multicast address
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Println("Error resolving multicast address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	// Receive multicast packets
	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading multicast packet:", err)
			return
		}
		fmt.Println("Received multicast message:", string(buffer[:n]))
	}
}
