package main

import (
	packets "CCP/Packets"
	"fmt"
	"CCP/Packets/Payloads"
)


// Car Communication Protocol
func main() {

	alert := Payloads.EncodeAlert("description bla bla bla")
	pkt := packets.Create_packet(alert)

	fmt.Println(pkt)

	decoded := packets.Packet{}
	fmt.Println(decoded.Decode_binary_packet(pkt))
}
