package Network

import (
	"fmt"
	"log"
	"net"
	"CCP/Packets"
	"CCP/Packets/Payloads"
)

var socket_pool map[net.Conn]net.Addr

func recv_all(length uint16,c net.Conn) ([]byte, error){
	payload := make([]byte, length)
	var tmp []byte

	for length > 0{

		n, err := c.Read(tmp)

		if err != nil || n == 0 {
			return nil, err
		}

		append(payload,tmp)

		payload -= n
	}
	return payload,nil
}

func handleConnection(c net.Conn) {

	log.Printf("Client %v connected.", c.RemoteAddr())

	packet := make([]byte, 6) // Size of header: 6 bytes

	for {
			n, err := c.Read(packet)

			if err != nil || n == 0 {
				close_connection(c)
				return
			}

			//Parse the header
			parsed_header,_ := Packets.Decode_binary_header(packet)

			//Get the payload
			payload, err := recv_all(parsed_header.Payload_length,c)
			if err != nil {
				close_connection(c)
				return
			}

			//Decode the payload
			decoded_payload,_ := Packets.Decode_binary_payload(parsed_header,payload)

			switch payload := decoded_payload.(type) {

			case Payloads.Alert:
				fmt.Println("Alert message :D")
				fmt.Println(string(payload.Description))

			default:
				fmt.Print(":/")

			}

	}
}

func close_connection(c net.Conn){
	log.Printf("Connection from %v closed.", c.RemoteAddr())
	c.Close()
	delete(socket_pool,c)
}

func Server() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server up and listening on port 6000")

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Println(err)
			close_connection(conn)
			continue
		}

		socket_pool[conn] = conn.RemoteAddr()
		go handleConnection(conn)
	}
}