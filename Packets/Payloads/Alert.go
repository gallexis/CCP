package Payloads

import (
	"bytes"
	"encoding/binary"
)

var command_name string = "alert"

type Alert struct{
	Description []byte
}

func (alert Alert) Get_command_name() string{
	return command_name
}

func EncodeAlert(description string) *Alert {
	alert := Alert{}

	alert.Description = []byte(description)
	return &alert
}

func DecodeAlert(payload *bytes.Buffer) *Alert {
	alert := Alert{}
	alert.Description = payload.Bytes()

	return &alert
}

func (alert *Alert) Forge() []byte{
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, alert.Description)
	return buffer.Bytes()
}