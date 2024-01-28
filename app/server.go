package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const _wee_redis_server_address = "0.0.0.0:6379"

func main() {
	l, err := net.Listen("tcp", _wee_redis_server_address)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	log.Printf("Wee-Redis server is listening on %s", _wee_redis_server_address)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go func(conn net.Conn) {
			log.Printf("Received connection from %s", conn.RemoteAddr())
			HandleClient(conn)
			conn.Close()
		}(conn)
	}
}
