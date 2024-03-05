package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

// Handle TCP connections
func handleTCPConnections(listener net.Listener, jugadoresMux *sync.Mutex, jugadores *map[string]PLAYER) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("TCP accept error:", err)
			continue
		}
		go handleTCPConnection(conn, jugadoresMux, jugadores)
	}
}

func handleTCPConnection(conn net.Conn, jugadoresMux *sync.Mutex, jugadores *map[string]PLAYER) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		receivedData := string(buffer[:n])
		separated_data := strings.Split(receivedData, "/")
		remoteAddr := conn.RemoteAddr().String()

		if separated_data[0] == "r" {
			fmt.Println("Registro:", separated_data[1]+remoteAddr)
			addPlayer(separated_data[1], remoteAddr, jugadoresMux, jugadores, &conn)

		} else if separated_data[0] == "c" {
			fmt.Println("Casilla: " + separated_data[1])
		}
	}

}
