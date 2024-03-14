package main

import (
	"fmt"
	"net"
	"strconv"
)

func addPlayer(nombre string, ipAddress string, mem *MEMORY, conn *net.Conn) {

	mem.jugadoresMux.Lock()
	defer mem.jugadoresMux.Unlock()

	if player, exists := mem.jugadores[ipAddress]; exists {
		player.conn = *conn
		fmt.Println("Jugador Reconectado")
	} else {
		player := PLAYER{
			conn:      *conn,
			nombre:    nombre,
			score:     0,
			casilla:   10,
			ipAddress: ipAddress,
		}

		(mem.jugadores)[ipAddress] = player
		player.conn.Write([]byte(mem.multicastAddr))
		fmt.Println("Nuevo jugador registrado")
	}

	mapString := fmt.Sprintf("%v", mem.jugadores)

	fmt.Println(mapString)

}

func addPoint(ipAddress string, mem *MEMORY) {
	mem.pointMux.Lock()
	if !mem.gotPoint {
		mem.gotPoint = true
		player, ok := mem.jugadores[ipAddress]
		if ok {
			player.score += 1
			mem.jugadores[ipAddress] = player
			fmt.Println("Player: " + player.nombre + " got the point" + player.ipAddress + " - " + strconv.Itoa(mem.jugadores[ipAddress].score))
			// message = fmt.Sprintf("%v", player.score)

		} else {
			fmt.Println("Player not found in jugadores map")
		}
	}

	mem.pointMux.Unlock()

}

func flush(mem *MEMORY) {
	for key := range mem.jugadores {
		mem.jugadores[key].conn.Close()
		delete(mem.jugadores, key)
	}
	// fmt.Println((mem.jugadores))
}
