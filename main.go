package main

import (
	"log"
	"net"
)

func handleConnection(net.Conn) error {
	return nil
}

func startServer() error {
	log.Println("Starting server")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			continue
		}
		handleConnection(conn)
		if err != nil {
			log.Println("Error handling connection", err)
			continue
		}
		//go func() {
		//	if err := handleConnection(conn, eventChannel); err != nil {
		//		log.Println("Error handling connection", err)
		//		return
		//	}
		//}()
	}
}

func main() {
	err := startServer()
	if err != nil {
		log.Fatal(err)
	}
}
