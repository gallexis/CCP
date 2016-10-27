package Payloads

import (
	"bytes"
	"encoding/binary"
)

var mesage string = "alert"

type Alert struct{
	Message     [5]byte
	Description []byte
}

func (alert Alert) GetName() string{
	return mesage
}

func EncodeAlert(description string) *Alert {
	alert := Alert{}

	copy(alert.Message[:],[]byte(mesage))
	alert.Description = []byte(description)
	return &alert
}

func DecodeAlert(payload *bytes.Buffer) *Alert {
	alert := Alert{}
	copy(alert.Message[:],payload.Next(5))
	alert.Description = payload.Bytes()

	return &alert
}

func (alert *Alert) Forge() []byte{
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, alert.Message)
	binary.Write(&buffer, binary.LittleEndian, alert.Description)
	return buffer.Bytes()
}