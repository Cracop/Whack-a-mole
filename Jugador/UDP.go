package main

import (
	"fmt"
	"net"
	"strconv"
)

func receiveUDP(c *CONNECTION, multicastAddr string) {
	// Multicast address and port
	// multicastAddr := c.mult.RemoteAddr().String()
	// multicastAddr := "224.239.239.239:5353"
	// multicastAddr := "127.0.0.1:10000"

	// Resolve multicast address
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Println("Error resolving multicast address:", err)
		return
	}

	// Create UDP connection
	c.mult, err = net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}

	defer c.mult.Close()

	// Receive multicast packets
	for {
		buffer := make([]byte, 1024)
		n, _, err := c.mult.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading multicast packet:", err)
			return
		}
		fmt.Println("Received multicast message:", string(buffer[:n]))
		monster, _ := strconv.Atoi(string(buffer[:n]))
		c.start = true
		c.monster <- monster
		fmt.Println(c.monster)
	}
}
