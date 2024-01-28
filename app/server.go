package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const WeeRedisServerAddress = "0.0.0.0:6379"

func main() {
	l, err := net.Listen("tcp", WeeRedisServerAddress)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	log.Printf("Wee-Redis server is listening on %s", WeeRedisServerAddress)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go func(conn net.Conn) {
			log.Printf("Received connection from %s", conn.RemoteAddr())
			HandleClient(conn)
		}(conn)
	}
}
