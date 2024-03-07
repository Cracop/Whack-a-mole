package main

import (
	"strconv"
)

func main() {

	tcpAddress := "127.0.0.1:49999"
	udpAddress := "224.0.0.1:9999"
	baseName := "POGO"

	for i := 0; i < 2; i++ {
		newName := baseName + strconv.Itoa(i)
		p := POGO{}

		go p.run(newName, tcpAddress, udpAddress)
	}

	select {}
}