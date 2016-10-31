package main

import (
	"CCP/Network"
	"CCP/Packets"
	"CCP/Packets/Payloads"
	"fmt"
	"os"
	"time"
	"log"
)

// Car Communication Protocol
func main() {

	// Server
	if len(os.Args) > 1 {

		Network.Start_server()

	} else {

		c, err := Network.NewConnection("192.168.1.72", "6000")
		if err != nil{
			log.Print("Connection error to server: ")
			log.Println(err)
			return
		}

		alert := Payloads.EncodeAlert("Test Alert: "+c.Socket.LocalAddr().String())
		pkt := Packets.Create_packet(alert)

		fmt.Println(pkt)

		time.Sleep(100 * time.Millisecond)
		c.Send_All(pkt)
		time.Sleep(1000 * time.Millisecond)

		Network.Client_handle_connection(c)
	}
}
