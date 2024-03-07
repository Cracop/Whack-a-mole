package main

import "net"

type PLAYER struct {
	conn      net.Conn
	nombre    string
	score     int
	casilla   int
	ipAddress string
}
