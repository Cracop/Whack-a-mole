package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	jugadores    = make(map[string]PLAYER)
	jugadoresMux sync.Mutex
)

func main() {

	PortTCP := ":49999"
	multicastAddr := "224.0.0.1:9999"

	tcpListener, err := net.Listen("tcp4", PortTCP)
	if err != nil {
		panic(err)
	}

	defer tcpListener.Close()

	fmt.Println("Server listening in port", PortTCP)

	go handleTCPConnections(tcpListener, &jugadoresMux, &jugadores)

	// Resolve multicast address
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Println("Error resolving multicast address:", err)
		return
	}

	go multicast(addr)

	// Block main goroutine to keep server running
	select {}
}
