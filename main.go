package main

import (
	// "bufio"
	"fmt"
	"log"
	"net"
	"os"

	netcat "netcat/func"
)

func main() {
	// create the log file
	logFile, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("error opening logfile :", err)
	}
	defer logFile.Close()
	// save the chat
	log.SetOutput(logFile)
	
	// listening to port
	var portnum string
	if len(os.Args) == 1 {
		portnum = "8989"
	} else if len(os.Args) == 2 {
		portnum = ":" + os.Args[1]
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
	}

	listener, err := net.Listen("tcp", netcat.Netty()+portnum)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	ip := netcat.Netty()
	// Log server's IP addresses
	fmt.Println("Connected to server: ", ip)
	netcat.Address()
	fmt.Printf("server is listening on port%s \n", portnum)

	for {
		var Conn net.Conn
		Conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection failed: %v", err)
			continue
		}
		// log.Printf("%s has join the chat..\n", netcat.Clients.Name)
		go netcat.Handler(Conn)

	}
}
