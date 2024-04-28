package netcat

import (
	"bufio"
	"fmt"
	"strings"

	// "io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func welcome() string {
	message, err := ioutil.ReadFile("penguin.txt")
	if err != nil {
		log.Fatal("error displaying welcome image: ", err)
	}
	log.Printf("Welcome to TCP-Chat!\n")
	return string(message) + "\n"
}

func Handler(network net.Conn) {
	defer network.Close()

	msg := welcome()
invalidStatment:
	network.Write([]byte(msg + "\n"))
	network.Write([]byte("Enter your name: "))
	Reader := bufio.NewReader(network)
	deadline := time.Now().Add(30 * time.Second)
	network.SetReadDeadline(deadline)

	name, err := Reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Connection failed: %v", err)
		return
	}
	name = strings.TrimSpace(name)
	if name == "" || name == " " || strings.Contains(name, " ") || strings.Contains(name, "\033") || strings.Contains(name, "\n") || strings.Contains(name, "\t") {
		goto invalidStatment
	}

	muclient.Lock()
	if _, exist := allClients[name]; exist {
		muclient.Unlock()
		network.Write([]byte("This name is already taken.\n"))
		goto invalidStatment
	}
	if len(allClients) >= 10 {
		muclient.Unlock()
		network.Write([]byte("maximum client count exceeded"))
		return
	}
	network.SetReadDeadline(time.Time{})

	clinet := &client{
		Network: network,
		Name:    name,
	}
	allClients[name] = clinet

	muclient.Unlock()
	network.Write([]byte(fmt.Sprintf("Welcome %s!\n", name)))

	for {
		network.Write([]byte(fmt.Sprintf("\n[%s]:", name)))
		msg, err := Reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Connection failed: %v", err)
			return
		}
		msg = strings.TrimSpace(msg)
		if msg == "" || strings.Contains(msg, "\033") || strings.Contains(msg, "\n") {
			continue
		}

		Broadcast(fmt.Sprintf("\n[%s] [%s]: %s\n", time.Now().Format("02-01-2006: 15:04:05"), name, msg))

	}
	fmt.Sprintf("%s has left our chat...", Clients.Name)
}

func Broadcast(message string) {
	for _, client := range allClients {
		client.Network.Write([]byte(message))
	}
}

// Broadcast("Server", fmt.Sprintf("%s has joined the chat.. ", clinet.Name))
// Broadcast(fmt.Sprintf("%s has left our chat...", Clients.Name))

// func Broadcast(sender, message string) {
// 	muclient.Lock()
// 	defer muclient.Unlock()
// 	formatted := ""
// 	if sender == "Server" && strings.Contains(message, "has left the chat..") || strings.Contains(message, "has joined the chat..") || strings.Contains(message, "is now known as") {
// 		formatted = fmt.Sprintf("%s\n", message)
// 		for _, client := range allClients {
// 			if !strings.Contains(message, client.Name) {
// 				_, err := client.Network.Write([]byte(formatted))
// 				if err != nil {
// 					log.Printf("Failed to send message to %s: %v", client.Name, err)
// 				}
// 			}
// 		}

//		} else {
//			formatted = fmt.Sprintf("[%s],[%s]:\n %s \n", time.Now().Format("02-01-2006: 15:04:05"), sender, message)
//			muhistory.Lock()
//			history = append(history, formatted)
//			muhistory.Unlock()
//			for _, client := range allClients {
//				_, err := client.Network.Write([]byte(formatted))
//				if err != nil {
//					log.Printf("Failed to send message to %s: %v", client.Name, err)
//				}
//			}
//		}
//	}
