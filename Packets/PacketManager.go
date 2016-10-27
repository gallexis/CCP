package Packets

import (
	//"fmt"
	"bytes"
	"encoding/binary"
	payloads "CCP/Packets/Payloads"
	"fmt"
)

type packet_ struct{
	hd header
	pl []byte
}

type header struct {
	command_name [5]byte
	payload_length uint16
}

type payload_ interface{}

func Test(){

	h1 := &header{}
	h1.payload_length = uint16(12)
	copy(h1.command_name[:], []byte("alert"))


	pl := []byte("123456789012")

	fmt.Println(h1)
	fmt.Println(pl)

	b := h1.header_writer()
	fmt.Println(b)

	h2 := &header{}
	h2.header_reader(bytes.NewBuffer(b))
	fmt.Println(h2)
	//-------

	pck := &packet_{}
	pck.hd = *h1
	pck.pl = []byte("alexis")

	fmt.Println(pck)
	fmt.Println(pck.Encode_packet_to_binary())
	fmt.Println(pck.Decode_binary_packet(pck.Encode_packet_to_binary()))


}



func (packet *packet_) Decode_binary_packet(pckt []byte) payload_{
	header := &header{}
	payload := ""

	buff_packet := bytes.NewBuffer(pckt)
	header.header_reader(buff_packet)

	if bytes.HasSuffix(header.command_name[:], []byte("alert")){
		payload = payloads.V()
		//fmt.Println( payload )
	}else {
		payload = "err"
	}
	return payload
}


func (packet *packet_) Encode_packet_to_binary() []byte{
	return append(packet.hd.header_writer()[:], packet.pl[:]...)
}

func (header *header) header_reader(packet *bytes.Buffer) (error){

	copy(header.command_name[:],packet.Next(5) )
	header.payload_length = binary.LittleEndian.Uint16(packet.Next(2))

	//binary.LittleEndian.PutUint16(header.payload_length, binary.LittleEndian.Uint16(packet.Next(2)))

	return  nil
}

func (header *header) header_writer() []byte{
	var bin_buf bytes.Buffer

	binary.Write(&bin_buf, binary.LittleEndian, header)
	return bin_buf.Bytes()
}

/*
func Binary_to_payload(packet []byte) Header{



}


func Payload_to_binary(payload payload) []byte{


}
*/