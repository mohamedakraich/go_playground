/* LookupPort
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// go run lookupport.go tcp telnet
// go run lookupport.go udp quake
func main() {
	if len(os.Args) != 3 {
		log.Printf("Usage: %s network-type service\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	fmt.Println("Service port ", port)
}
