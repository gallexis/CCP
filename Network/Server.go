package Network

import (
	"CCP/Packets"
	"CCP/Packets/Payloads"
	"bufio"
	"fmt"
	"log"
	"net"
)

var socket_pool = make(map[net.Conn]net.Addr)
var HEADER_SIZE int = 7

func recv_all(length int, c net.Conn) ([]byte, error) {

	reader := bufio.NewReader(c)
	buf := make([]byte, length)

	for length > 0 {

		n, err := reader.Read(buf)
		if err != nil || n == 0 {
			return nil, err
		}

		length -= n
	}
	return buf, nil
}

func handleConnection(c net.Conn) {

	log.Printf("Client %v connected.", c.RemoteAddr())

	packet := make([]byte, HEADER_SIZE)

	for {
		//Get header
		n, err := c.Read(packet)
		if err != nil || n != HEADER_SIZE {
			log.Print("Disconnected: ", err)
			close_connection(c)
			return
		}

		//Parse the header
		parsed_header, err := Packets.Decode_binary_header(packet)
		if err != nil {
			log.Print("Problem parsing header: ", err)
			close_connection(c)
			return
		}

		//Get the payload
		payload_size := int(parsed_header.Payload_length)
		payload, err := recv_all(payload_size, c)
		if err != nil || len(payload) != payload_size {
			log.Print("Problem getting payload: ", err)
			close_connection(c)
			return
		}

		//Decode the payload
		decoded_payload, err := Packets.Decode_binary_payload(parsed_header, payload)
		if err != nil {
			log.Print("Problem decoding the payload: ", err)
			close_connection(c)
			return
		}

		switch payload := decoded_payload.(type) {
		case Payloads.Alert:
			fmt.Println("Alert message :D")
			fmt.Println(string(payload.Description))

		default:
			fmt.Print(":/")

		}

	}
}

func close_connection(c net.Conn) {
	log.Printf("Connection from %v closed.", c.RemoteAddr())
	c.Close()
	delete(socket_pool, c)
}

func Start_server() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server up and listening on port 6000")

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Print("Error incoming connection: ")
			log.Println(err)
			close_connection(conn)
			continue
		}

		socket_pool[conn] = conn.RemoteAddr()
		go handleConnection(conn)
	}
}
