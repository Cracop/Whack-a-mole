package main

import (
	"fmt"
	"net"
)

func LoginTCP(c *CONNECTION) {

	addr := fmt.Sprintf("%s:%s", c.ipAddress, c.puerto)
	var err error
	c.conn, err = net.Dial("tcp", addr) // Replace with the target server and port
	if err != nil {
		fmt.Println("TCP send error:", err)
		return
	}
	// defer conn.Close()

	// Send a message
	message := fmt.Sprintf("r/%s", c.nombre)
	c.conn.Write([]byte(message))
	fmt.Println("TCP package sent:", message)
}

func whackTCP(c *CONNECTION, cell int) {
	message := ""
	if cell == c.cell {
		message = "c/success"
	} else {
		message = "c/fail"
	}
	c.conn.Write([]byte(message))
	fmt.Println("TCP package sent:", message)
}
