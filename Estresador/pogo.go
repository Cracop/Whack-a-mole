package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"strconv"
	"time"
)

type POGO struct {
	// contador   int
	name       string
	TCPconn    net.Conn
	tcpAddress string
	udpAddress string
	UDPconn    *net.UDPConn
	monster    int
	buffer     []byte
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
	fmt.Println("TCP package sent:", message)
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
		p.monster, _ = strconv.Atoi(string(buffer[:n]))
		// fmt.Println(p.monster)
		// fmt.Println("Received multicast message:", strconv.Itoa(p.monster))
	}
}

func (p *POGO) whack() {
	c := rand.IntN(1)
	fmt.Println(c)
	if c == 0 {
		message := "c/success"
		p.TCPconn.Write([]byte(message))
	}

	_, err := p.TCPconn.Read(p.buffer)

	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

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
	go p.joinMulticast()

	for {
		p.whack()
		time.Sleep(1 * time.Second)
	}

}
