/* IPv4Mask
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// go run ipv4mask.go 192.168.1.3
func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		log.Fatalln("nil Invalid address")
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		"\nDefault mask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
	derivedMask := net.IPv4Mask(255, 255, 0, 0) // working on mask
	fmt.Printf("Network using %s: %s\n", derivedMask, addr.Mask(derivedMask))
}
