/* ResolveIP
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// go doc net.Dial
// go run resolveip.go www.google.com
func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		log.Fatalln("Resolution error", err.Error())
	}
	fmt.Println("Resolved address is ", addr.String())
}
