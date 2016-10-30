package main

import (
	"CCP/Network"
	"CCP/Packets"
	"CCP/Packets/Payloads"
	"fmt"
	"os"
	"time"
)

// Car Communication Protocol
func main() {

	// Server
	if len(os.Args) > 1 {

		Network.Start_server()

	} else {

		c, _ := Network.NewConnection("localhost", "6000")

		alert := Payloads.EncodeAlert("Test Alert: "+c.Socket.LocalAddr().String())
		pkt := Packets.Create_packet(alert)

		fmt.Println(pkt)

		time.Sleep(100 * time.Millisecond)
		c.SendAll(pkt)
		time.Sleep(1000 * time.Millisecond)
	}
}
