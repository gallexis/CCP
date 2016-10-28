package Packets

import (
	"bytes"
	"encoding/binary"
	"CCP/Packets/Payloads"
)

type Packet struct{
	header  Header
	payload Payload
}

type Header struct {
	command_name [5]byte
	payload_length uint16
}

type Payload interface{
	Get_command_name() string
	Forge() []byte
}

func Create_packet(payload Payload) []byte{
	header := Header{}
	copy(header.command_name[:],[]byte(payload.Get_command_name()) )
	header.payload_length = uint16(len(payload.Forge()))

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
		alert :=  Payloads.DecodeAlert(buffer_packet)
		return string(alert.Description)
	}else {
		return []byte("err")
	}
}

func (packet *Packet) Encode_packet_to_binary() []byte{
	return append(packet.header.header_writer()[:], packet.payload.Forge()...)
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
