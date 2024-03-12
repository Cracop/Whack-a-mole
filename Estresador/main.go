package main

import (
	"strconv"
)

func main() {

	// tcpAddress := "10.10.1.190:5050"
	tcpAddress := "127.0.0.1:5050"
	udpAddress := "224.0.0.1:9999"
	baseName := "POGO"

	for i := 0; i < 6; i++ {
		newName := baseName + strconv.Itoa(i)
		p := POGO{over: false}

		go p.run(newName, tcpAddress, udpAddress)
	}

	select {}
}
