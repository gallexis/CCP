package Packets

import (
	"bytes"
	"CCP/Packets/Payloads"
)

type Packet struct{
	header  Header
	payload Payload
}

func Create_packet(payload Payload) []byte{

	// forge payload now, calculate its size, then pass it to Encode_packet_to_binary.
	// That way it will not be forged twice
	forged_payload := payload.Forge()

	header := Header{}
	copy(header.Command_name[:],[]byte(payload.Get_command_name()) )
	header.Payload_length = uint16(len(forged_payload))

	packet := Packet{}
	packet.header = header
	packet.payload = payload

	return packet.Encode_packet_to_binary(forged_payload)
}

func Decode_binary_header(pckt []byte) (*Header,error) {
	header := &Header{}

	buffer_packet := bytes.NewBuffer(pckt)
	header.read_header(buffer_packet)

	return header,nil
}

func Decode_binary_payload(header *Header, payload []byte) (Payload,error) {
	buffer_payload := bytes.NewBuffer(payload)

	if bytes.HasSuffix(header.Command_name[:], []byte("alert")){
		Alert :=  Payloads.DecodeAlert(buffer_payload)
		return  Alert,nil
	}else {
		return nil,nil
	}
}

func (packet *Packet) Encode_packet_to_binary(forged_payload []byte) []byte{
	return append(packet.header.write_header()[:], forged_payload...)
}


