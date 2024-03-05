package main

import (
	"fmt"
	"net"
	"sync"
)

func addPlayer(nombre string, ipAddress string, jugadoresMux *sync.Mutex, jugadores *map[string]PLAYER, conn *net.Conn) {

	player := PLAYER{
		conn:    *conn,
		nombre:  nombre,
		score:   0,
		casilla: 10,
	}

	jugadoresMux.Lock()
	(*jugadores)[ipAddress] = player
	jugadoresMux.Unlock()
	mapString := fmt.Sprintf("%v", *jugadores)

	fmt.Println(mapString)
}

func removePlayers(jugadoresMux *sync.Mutex, jugadores *map[string]PLAYER) {

}
