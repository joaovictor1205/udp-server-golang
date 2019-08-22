package main

import (
	"fmt"
	"log"
	"net"
)

func UDP_SERVER(conn *net.UDPConn) {

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP client: ", addr)
	fmt.Println("Mensagem recebida do Client: ", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	message := []byte("Digite sua mensagem: ")
	_, err = conn.WriteToUDP(message, addr)

	if err != nil {
		log.Println(err)
	}

}

func main() {

	hostName := "localhost"
	portNum := "8080"
	service := hostName + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp4", service)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP SERVER on port: ", portNum)

	defer ln.Close()

	for {
		UDP_SERVER(ln)
	}

}
