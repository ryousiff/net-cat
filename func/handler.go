package netcat

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"

	// netcat "netcat/func"

	// "os/signal"
	"net"
	"strings"
	"time"
)

func Welcome() string {
	message, err := ioutil.ReadFile("penguin.txt")
	if err != nil {
		log.Fatal("error displaying welcome image: ", err)
	}
	// .Printf("Welcome to TCP-Chat!\n")
	return string(message) + "\n"
}

func PrevChat() string {
	message, err := ioutil.ReadFile("log.txt")
	if err != nil {
		log.Fatal("error displaying welcome image: ", err)
	}
	return string(message) + "\n"
}

func Handler(network net.Conn) {
	defer network.Close()
	// rere := bufio.NewReader(os.Stdin)

	msg := Welcome()
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
	// i want a to check if the name already exists in the clients struct
	// if it does, then it will write to the network of the client and it will goto invalidstatment
	// if it does not, then it will add the client to the clients struct
	// and then it will write to the network of the client and it will goto invalidstatment
	for _, client := range clients {
		if client.Name == name {
			muclient.Unlock()
			network.Write([]byte("name already taken"))
			goto invalidStatment
		}
	}
	if len(clients) >= 10 {
		muclient.Unlock()
		network.Write([]byte("maximum client count exceeded"))
		return
	}
	network.SetReadDeadline(time.Time{})

	client := &client{
		Network: network,
	}
	client.Name = name
	clients = append(clients, client)

	muclient.Unlock()
	log.Printf("%s has join the chat..\n", name)
	network.Write([]byte(fmt.Sprintf("Welcome %s!\n", name)))
	g := PrevChat()
	network.Write([]byte(g + "\n"))

	Broadcast(fmt.Sprintf("\n%s has joined our chat... \n", client.Name))

	// muhistory.Lock()
	// for _, message := range history {
	// 	_, err := client.Network.Write([]byte(message))
	// 	log.Print(message)
	// 	if err != nil {
	// 		log.Printf("Error sending past message to %s: %v", client.Name, err)
	// 		break
	// 	}
	// }
	// muhistory.Unlock()
	for {
		// network.Write([]byte(fmt.Sprintf("\n[%s]:", name)))
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
	RemoveClient(client)
	Broadcast(fmt.Sprintf("\n%s has left our chat...\n", name))

}

func Broadcast(message string) {
	for _, client := range clients {
		client.Network.Write([]byte(message))
		// muhistory.Lock()
		// history = append(history, message)
		// muhistory.Unlock()
	}
}

func RemoveClient(client *client) {
	for i, c := range clients {
		if c == client {
			muclient.Lock()
			clients = append(clients[:i], clients[i+1:]...)
			muclient.Unlock()
			break
		}
	}
}
