package main

import (
	"net"
)

type CONNECTION struct {
	ipAddress string
	puerto    string
	nombre    string
	conn      net.Conn
	mult      *net.UDPConn
	monster   chan int
	cell      int
	start     bool
	puntaje   int
	ganador   string
	gameGUI   *GAMEGUI
	stop      chan bool
}
