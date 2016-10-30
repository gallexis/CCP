package main

import (
	"CCP/Packets"
	"CCP/Packets/Payloads"
	"os"
	"CCP/Network"
	"fmt"
)


// Car Communication Protocol
func main() {

	// Client
	if len(os.Args) > 1{

		c,_ := Network.NewConnection("localhost","6000")

		alert := Payloads.EncodeAlert("Test Alert")
		pkt := Packets.Create_packet(alert)

		fmt.Println(pkt)

		c.SendAll(pkt)

	}else{

		Network.Start_server()


	}
}

