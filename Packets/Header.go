package Packets

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	command_name [5]byte
	payload_length uint16
}

func (header *Header) read_header(packet *bytes.Buffer) (error){

	copy(header.command_name[:],packet.Next(5) )
	header.payload_length = binary.LittleEndian.Uint16(packet.Next(2))

	return  nil
}

func (header *Header) write_header() []byte{
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, header)
	return buffer.Bytes()
}