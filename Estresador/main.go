package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	// tcpAddress := "10.10.1.190:5050"
	tcpAddress := "127.0.0.1:5050"
	udpAddress := "224.0.0.1:9999"
	baseName := "POGO"

	var csvMUX sync.Mutex
	// var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	// n := 10000

	maxClients := 50

	for i := 50; i <= maxClients; i = i + 50 {
		for j := 0; j < 10; j++ {
			respuesta_juego := [][]float64{}
			respuesta_registro := [][]float64{}
			fmt.Println("Empieza nuevo juego")
			for k := 0; k < i; k++ {
				wg2.Add(1)
				newName := baseName + strconv.Itoa(k)
				p := POGO{over: false}

				go func() {
					defer wg2.Done() // Decrement the WaitGroup counter when the goroutine completes
					p.run(newName, tcpAddress, udpAddress)
					csvMUX.Lock()
					respuesta_juego = append(respuesta_juego, p.elapsed)
					respuesta_registro = append(respuesta_registro, p.register)
					csvMUX.Unlock()
				}()
			}
			wg2.Wait()
			exportar(respuesta_juego, "juego/", strconv.Itoa(i)+"_clientes_"+strconv.Itoa(j))
			exportar(respuesta_registro, "registro/", strconv.Itoa(i)+"_clientes_"+strconv.Itoa(j))
			// fmt.Println("Vamos al siguiente juego")
		}
	}

	// fmt.Println(respuesta_juego)
	// fmt.Println(respuesta_registro)

}
