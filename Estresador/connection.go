package main

import (
	"net"
)

type CONNECTION struct {
	ipAddress string
	puerto    string
	nombre    string
	conn      net.Conn
	mult      net.UDPConn
	monster   int
}
