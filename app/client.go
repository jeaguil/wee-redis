package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func HandleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	clientAddress := conn.RemoteAddr()
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection closed by client %s", clientAddress)
			}
			log.Printf("Error reading data:%s", err)
			break
		}
		_, err = writer.WriteString("+PONG\r\n")
		if err != nil {
			log.Printf("Error writing data to client %s: %s", clientAddress, err)
			break
		}
		writer.Flush()
	}
}
