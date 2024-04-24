package netcat

import (
	// "fmt"
	"net"
	// "os"
	"log"
)

func Netty() string {

	Network, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
    }
	defer Network.Close()
	local := Network.LocalAddr().(*net.UDPAddr)
	return local.IP.String()	
}

	func Address() {

	add, err := net.InterfaceAddrs()
	if err!= nil {
		log.Println("Failed to get local IP addresses:", err)
    }

	for _, address := range add {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Printf("Server accessible at IP: %s", ipnet.IP.String())
				// fmt.Printf("Server accessible at IP: %s\n", ipnet.IP.String())
			// } else {
			// 	log.Printf("Server inaccessible %s", ipnet.IP.String())
			// 	// Handle non-IPv4 address
			}
		}
	}
}