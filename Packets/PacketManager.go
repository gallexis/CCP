package Packets

import (
	"bytes"
	"encoding/binary"
	payloads "CCP/Packets/Payloads"
	"fmt"
)

type Packet struct{
	header  Header
	payload []byte
}

type Header struct {
	command_name [5]byte
	payload_length uint16
}

type Payload []byte

func Test(){

	alert := payloads.EncodeAlert("description bla bla bla")
	pkt := Create_packet(alert.Forge(),alert.GetName())

	fmt.Println(pkt)

	decoded := Packet{}
	fmt.Println(decoded.Decode_binary_packet(pkt))


}


//func Decode_binary()interface{} {}

func Create_packet(payload Payload, command_name string) []byte{
	header := Header{}
	header.payload_length = uint16(len(payload))
	copy(header.command_name[:],[]byte(command_name) )

	packet := Packet{}
	packet.header = header
	packet.payload = payload

	return packet.Encode_packet_to_binary()
}

func (packet *Packet) Decode_binary_packet(pckt []byte) interface{} {
	header := &Header{}

	buffer_packet := bytes.NewBuffer(pckt)
	header.header_reader(buffer_packet)

	if bytes.HasSuffix(header.command_name[:], []byte("alert")){
		alert :=  payloads.DecodeAlert(buffer_packet)
		return string(alert.Description)
	}else {
		return []byte("err")
	}
}


func (packet *Packet) Encode_packet_to_binary() []byte{
	return append(packet.header.header_writer()[:], packet.payload[:]...)
}

func (header *Header) header_reader(packet *bytes.Buffer) (error){

	copy(header.command_name[:],packet.Next(5) )
	header.payload_length = binary.LittleEndian.Uint16(packet.Next(2))

	return  nil
}

func (header *Header) header_writer() []byte{
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, header)
	return buffer.Bytes()
}

/*
func Binary_to_payload(packet []byte) Header{



}


func Payload_to_binary(payload payload) []byte{


}
*/