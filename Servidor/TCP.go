package main

import (
	"fmt"
	"net"
	"strconv"
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

	for {
		// fmt.Println("==================")
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
			// fmt.Println("Registro:", separated_data[1]+remoteAddr)
			addPlayer(separated_data[1], remoteAddr, mem, &conn)

		} else if separated_data[0] == "c" {
			// fmt.Println("Casilla: " + separated_data[1])
			mem.pointMux.Lock()
			var message string
			if !mem.gotPoint {
				mem.gotPoint = true
				player, ok := mem.jugadores[remoteAddr]
				if ok {
					player.score += 1
					mem.jugadores[remoteAddr] = player
					fmt.Println("Player: " + player.nombre + " got the point" + player.ipAddress + " - " + strconv.Itoa(mem.jugadores[remoteAddr].score))
					message = fmt.Sprintf("%v", player.score)

				} else {
					fmt.Println("Player not found in jugadores map")
				}
			} else {
				message = fmt.Sprintf("%v", mem.jugadores[remoteAddr].score)
			}

			conn.Write([]byte(message))
			mem.pointMux.Unlock()

		}
	}

}
