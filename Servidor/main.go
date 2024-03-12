package main

import (
	"fmt"
	"net"
	"sync"
)

type MEMORY struct {
	jugadores    map[string]PLAYER
	jugadoresMux sync.Mutex
	gotPoint     bool
	pointMux     sync.Mutex
	winnerMux    sync.Mutex
	winner       string
}

func main() {

	mem := MEMORY{
		jugadores:    make(map[string]PLAYER),
		gotPoint:     false,
		jugadoresMux: sync.Mutex{},
		pointMux:     sync.Mutex{},
		winner:       "NULL",
		// over: make(chan bool),
	}

	PortTCP := ":5050"
	multicastAddr := "224.0.0.1:9999"

	tcpListener, err := net.Listen("tcp4", PortTCP)
	if err != nil {
		panic(err)
	}

	// defer tcpListener.Close()

	fmt.Println("Server listening in port", PortTCP)

	// Resolve multicast address
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Println("Error resolving multicast address:", err)
		return
	}

	go multicast(addr, &mem)

	handleTCPConnections(tcpListener, &mem)

	// // Block main goroutine to keep server running
	// select {}
}
