package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"p2p-dev/proto"
	"sync"
)

func main() {
	lst, err := net.Listen("tcp", ":1688")
	if err != nil {
		log.Fatalf("error listening %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			conn, err := lst.Accept()
			if err != nil {
				log.Printf("error accepting %v", err)
				continue
			}
			go connHandler(ctx, conn)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := <-c
	fmt.Println("Got signal: ", s)
}

var (
	conns   []*proto.Address
	connMtx = &sync.Mutex{}
)

func appendConn(addr *proto.Address) {
	connMtx.Lock()
	defer connMtx.Unlock()
	conns = append(conns, addr)
}

func connHandler(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <-ctx.Done():
			_ = conn.Close()
			return
		default:
		}

		message, err := proto.Read(conn)
		if err != nil {
			log.Printf("error reading %v", err)
			_ = conn.Close()
			return
		}

		switch message.Method {
		case proto.Conn:
		case proto.Addr:
		}
	}
}
