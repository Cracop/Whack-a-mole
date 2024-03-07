package main

import (
	"net"
	"sync"
)

func addPlayer(nombre string, ipAddress string, mem *MEMORY, conn *net.Conn) {

	player := PLAYER{
		conn:      *conn,
		nombre:    nombre,
		score:     0,
		casilla:   10,
		ipAddress: ipAddress,
	}

	jugadoresMux.Lock()
	(mem.jugadores)[ipAddress] = player
	jugadoresMux.Unlock()
	// mapString := fmt.Sprintf("%v", mem.jugadores)

	// fmt.Println(mapString)
}

func removePlayers(jugadoresMux *sync.Mutex, jugadores *map[string]PLAYER) {

}
