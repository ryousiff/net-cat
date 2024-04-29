package netcat

import (
	"net"
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
				
			}
		}
	}
}