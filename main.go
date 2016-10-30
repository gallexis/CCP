package main

import (
	"CCP/Network"
	"CCP/Packets"
	"CCP/Packets/Payloads"
	"fmt"
	"os"
)

// Car Communication Protocol
func main() {

	// Server
	if len(os.Args) > 1 {

		Network.Start_server()

	} else {

		c, _ := Network.NewConnection("localhost", "6000")

		alert := Payloads.EncodeAlert("Test Alert")
		pkt := Packets.Create_packet(alert)

		fmt.Println(pkt)

		c.SendAll(pkt)
	}
}
