package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"strconv"
	"time"
)

func multicast(addr *net.UDPAddr, mem *MEMORY) {
	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}

	defer conn.Close()

	for {

		if mem.winner == "NULL" {
			// fmt.Println("No hay ganador:", mem.winner)
			// Generate a random number between 0 and 8
			mole := rand.IntN(9) // Generates a random number between 0 and 8
			// message := "Hello from UDP multicast! " + strconv.Itoa(mole)
			message := strconv.Itoa(mole)
			mem.pointMux.Lock()
			mem.gotPoint = false
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error sending multicast packet:", err)
				return
			}
			mem.pointMux.Unlock()
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println("Ganador:", mem.winner)
			message := "w/" + mem.winner
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error sending multicast packet:", err)
				return

			}
			return
		}

		// fmt.Println("Sent multicast message:", mem.winner)
		// fmt.Println(message)

	}

}
