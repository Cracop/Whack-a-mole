package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func multicast(addr *net.UDPAddr) {
	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}

	defer conn.Close()

	log.Println("Cell tapped:", addr.String())

	for {

		// Seed the random number generator
		rand.Seed(time.Now().UnixNano())

		// Generate a random number between 0 and 8
		mole := rand.Intn(9) // Generates a random number between 0 and 8

		message := "Hello from UDP multicast! " + strconv.Itoa(mole)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending multicast packet:", err)
			return
		}
		fmt.Println("Sent multicast message:", message)
		time.Sleep(1 * time.Second) // Send message every 1 second
	}

}
