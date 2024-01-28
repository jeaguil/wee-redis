package main

import (
	"log"
	"net"
)

func HandleClient(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	log.Printf("Received %b bytes from %s", n, conn.RemoteAddr())
	log.Printf("Received the following data: %s", buf[:n])
	conn.Write([]byte("+PONG\r\n"))
	log.Printf("Sent %d bytes to %s", n, conn.RemoteAddr())
}
