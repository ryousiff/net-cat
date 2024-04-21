package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//create the log file
	logFile, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("error opening logfile :", err)
	}
	defer logFile.Close()
	//save the chat
	log.SetOutput(logFile)

	//listening to port
	var portnum string
	if len(os.Args) == 1 {
		portnum = ":8989"
	} else if len(os.Args) == 2 {
		portnum = ":" + os.Args[2]

	} else {
		fmt.Println(" [USAGE]: ./TCPChat $port")
	}
	listen, err := net.Listen("tcp", portnum)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer listen.Close()
	fmt.Sprintf("server is listening on port %s \n", portnum)

}
