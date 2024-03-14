package main

import (
	"fmt"
	"net"
	"strings"
)

// Handle TCP connections
func handleTCPConnections(listener net.Listener, mem *MEMORY) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("TCP accept error:", err)
			continue
		}
		go handleTCPConnection(conn, mem)
	}
}

func handleTCPConnection(conn net.Conn, mem *MEMORY) {
	defer conn.Close()
	defer fmt.Println("adios conn")
	buffer := make([]byte, 1024)
	for {
		// fmt.Println("==================")
		// fmt.Println("Voy a leer el tcp")

		n, err := conn.Read(buffer)
		// fmt.Println("Acabo de leer el tcp")
		if err != nil {
			// fmt.Println("Error reading1:", err)

			return
		}

		receivedData := string(buffer[:n])
		separated_data := strings.Split(receivedData, "/")
		remoteAddr := conn.RemoteAddr().String()
		var message string
		fmt.Println("Casilla: " + separated_data[1])
		if separated_data[0] == "r" {
			// fmt.Println("Registro:", separated_data[1]+remoteAddr)
			addPlayer(separated_data[1], remoteAddr, mem, &conn)
			//el problema está aquí

		} else if separated_data[1] == "success" {

			addPoint(remoteAddr, mem)
			// conn.Write([]byte(message))
		}
		mem.jugadoresMux.Lock()
		message = fmt.Sprintf("%v", mem.jugadores[remoteAddr].score)
		// fmt.Println(message)
		conn.Write([]byte(message))

		if mem.jugadores[remoteAddr].score > 4 {
			// mem.winner <- mem.jugadores[remoteAddr].nombre
			mem.winner = mem.jugadores[remoteAddr].nombre
			// fmt.Println(mem.jugadores[remoteAddr].nombre + " ganó")
			mem.jugadoresMux.Unlock()
			return
		}
		mem.jugadoresMux.Unlock()

		mem.winnerMux.Lock()
		if mem.winner != "NULL" {
			mem.winnerMux.Unlock()
			return
		}
		mem.winnerMux.Unlock()

	}

}
