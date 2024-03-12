package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"time"
)

type POGO struct {
	// contador   int
	name       string
	TCPconn    net.Conn
	tcpAddress string
	udpAddress string
	UDPconn    *net.UDPConn
	monster    string
	buffer     []byte
	over       bool
}

func (p *POGO) login() {
	var err error
	p.TCPconn, err = net.Dial("tcp", p.tcpAddress) // Replace with the target server and port
	if err != nil {
		fmt.Println("TCP send error:", err)
		return
	}

	message := fmt.Sprintf("r/%s", p.name)
	p.TCPconn.Write([]byte(message))
	// fmt.Println("TCP package sent:", message)
}

func (p *POGO) logout() {
	p.TCPconn.Close()

}

func (p *POGO) joinMulticast() {
	addr, err := net.ResolveUDPAddr("udp", p.udpAddress)
	if err != nil {
		fmt.Println("Error resolving multicast address:", err)
		return
	}

	p.UDPconn, err = net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer p.UDPconn.Close()

	// Receive multicast packets
	for {
		buffer := make([]byte, 1024)
		n, _, err := p.UDPconn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading multicast packet:", err)
			return
		}
		p.monster = string(buffer[:n])
		if p.monster[0] == 'w' {
			fmt.Println(p.monster)
			p.over = true
			return
		}
		// fmt.Println(p.monster)
		// fmt.Println("Received multicast message:", strconv.Itoa(p.monster))
	}
}

func (p *POGO) whack() {
	c := rand.IntN(1)
	// fmt.Println("Voy a mandar mi success")
	message := ""
	if c == 0 {
		message = "c/success"
	} else {
		message = "c/fail"
	}
	// fmt.Println("Ya lo mandé y ahora lo voy a leer")

	_, err := p.TCPconn.Write([]byte(message))
	start := time.Now()
	if err != nil {
		// fmt.Println("Error writing:", err)
		return
	}
	// fmt.Println(message)
	_, err = p.TCPconn.Read(p.buffer)

	if err != nil {
		// fmt.Println("Error reading:", err)
		return
	}
	elapsed := time.Since(start)
	seconds := elapsed.Seconds()
	fmt.Println("Time elapsed in seconds:", seconds)
	// receivedData := string(p.buffer[:n])
	// fmt.Println(receivedData)

}

// func (p *POGO) success() {

// }

func (p *POGO) run(name string, tcpAddress string, udpAddress string) {
	// fmt.Println("p.name")
	p.name = name
	p.tcpAddress = tcpAddress
	p.udpAddress = udpAddress
	p.buffer = make([]byte, 1024)
	// fmt.Println(p.name)
	p.login()
	// fmt.Println("Update received!")
	go p.joinMulticast()

	for {

		if !p.over {
			p.whack()
			time.Sleep(1 * time.Second)
		} else {
			// fmt.Println("Finito")
			return
		}

		// if !p.over {

		// }

		// Perform actions needed when an update is received

		// break
	}

}
