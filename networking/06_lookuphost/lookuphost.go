/* LookupHost
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// go run lookuphost.go go.dev
// dig go.dev A go.dev AAAA +short
func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	cname, _ := net.LookupCNAME(name)
	fmt.Println(cname)
	addrs, err := net.LookupHost(cname)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
