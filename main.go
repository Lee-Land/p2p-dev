package main

import (
	"log"
	"net"

	"p2p-dev/proto"
)

func main() {
	lst, err := net.Listen("tcp", ":1688")
	if err != nil {
		log.Fatalf("error listening %v", err)
	}

	for {
		conn, err := lst.Accept()
		if err != nil {
			log.Printf("error accepting %v", err)
			continue
		}

		proto.Parse(conn)
	}
}
