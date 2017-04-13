package main

import (
	"log"
	"net"
	"time"
)

const UDP_PORT_NO = "12000"
const TCP_PORT_NO = "12001"

func main() {

	var rlen int

	remote, err := net.ResolveUDPAddr("udp", "255.255.255.255:"+UDP_PORT_NO)

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	//localAddr, err := net.ResolveUDPAddr("udp", ":0")

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	conn, err := net.DialUDP("udp", nil, remote)

	conn.SetDeadline(time.Now().Add(5 * time.Second))

	defer conn.Close()

	s := "CON_REQ"

	rlen, err = conn.Write([]byte(s))

	if err != nil {
		log.Printf("Send Error: %v\n", err)
		return
	}

	log.Printf("Send %v\n", s)

	buf := make([]byte, 1024)

	rlen, addr, err := conn.ReadFrom(buf)

	if err != nil {
		log.Printf("Receive Error: %v\n", err)
		return
	}

	log.Printf("Receive: %v\n", string(buf[:rlen]), " from ", addr)

}
