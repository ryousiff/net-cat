package netcat

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
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
	if name == "" || name == " " || strings.Contains(name, " ") || strings.Contains(name, "\033") || strings.Contains(name, "\n") {
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

	Broadcast(fmt.Sprintf("%s has joined the chat... \n", clinet.Name))

	muhistory.Lock()
	for _, message := range history {
		_, err := clinet.Network.Write([]byte(message))
		log.Print(message)
		if err != nil {
			log.Printf("Error sending past message to %s: %v", clinet.Name, err)
			break
		}
	}
	muhistory.Unlock()
	for {
		network.Write([]byte(fmt.Sprintf("\n[%s]:", name)))
		msg, err := Reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Connection failed: %v", err)
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" || strings.Contains(msg, "\033") || strings.Contains(msg, "\n") {
			continue
		}

		Broadcast(fmt.Sprintf("\n[%s] [%s]: %s\n", time.Now().Format("02-01-2006: 15:04:05"), name, msg))
		log.Print(msg)
	}
	RemoveClient(Clients)
	Broadcast(fmt.Sprintf("%s has left our chat...\n", name))

}

func Broadcast(message string) {
	for _, client := range allClients {
		client.Network.Write([]byte(message))
		muhistory.Lock()
		history = append(history, message)
		muhistory.Unlock()
	}
}

func RemoveClient(clients client) {
	for i, c := range arrayofclient {
		if c == clients {
			muclient.Lock()
			arrayofclient = append(arrayofclient[:i], arrayofclient[i+1:]...)
			muclient.Unlock()
			break
		}
	}
}
