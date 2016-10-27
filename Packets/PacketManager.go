package Packets

import (
	"fmt"
	"bytes"
	"encoding/binary"
	payloads "CCP/Packets/Payloads"
)

type Header struct {
	command_name [4]byte
	payload_length uint16
}

type Payload []byte


func (header *Header) PacketManager(packet []byte){

	buff_header := bytes.NewBuffer(packet)

	payload, _ := header_reader(buff_header)

	if header_parsed.command_name == "alert"{
		fmt.Println( payloads.V() )
		fmt.Println( payload )
	}

}

func (header *Header)header_reader(packet *bytes.Buffer) (*bytes.Buffer, error){
	header.command_name = packet.Next(4)
	header.payload_length = packet.Next(2)

	return  packet, nil
}

func (header *Header)header_writer() bytes.Buffer{
	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.BigEndian, header)
	return bin_buf
}

/*
func Binary_to_payload(packet []byte) Header{



}


func Payload_to_binary(payload payload) []byte{


}
*/