/* Mask
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

// go run mask.go 103.232.159.187 24 32
// go run mask.go fda3:97c:1eb:fff0:5444:903a:33f0:3a6b 52 128
func main() {
	if len(os.Args) != 4 {
		log.Printf("Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		log.Fatalln("nil Invalid address")
	}
	mask := net.CIDRMask(ones, bits)
	computedOnes, computedBits := mask.Size()
	network := addr.Mask(mask)
	fmt.Println("Address is ", addr.String(),
		"\nMask length is ", computedBits,
		"\nLeading ones count is ", computedOnes,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
}
