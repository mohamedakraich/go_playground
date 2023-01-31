/* IP
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// go run ip.go 127.0.0.1
// go run ip.go 0:0:0:0:0:0:0:1
func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
}
