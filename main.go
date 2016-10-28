package main

import (
	"CCP/Packets"
	"fmt"
	"CCP/Packets/Payloads"
)


// Car Communication Protocol
func main() {

	alert := Payloads.EncodeAlert("description bla bla bla")
	pkt := Packets.Create_packet(alert)

	fmt.Println(pkt)

	decoded,_ := Packets.Decode_binary_packet(pkt)

	switch v := decoded.(type) {

	case Payloads.Alert:
		fmt.Println("Alert message :D")
		fmt.Println(string(v.Description))

	default:
		fmt.Print(":/")

	}

}

