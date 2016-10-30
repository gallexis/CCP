package Network

import (
	"net"
	"fmt"
)

type Connection struct{
	Socket net.Conn
	Connected bool
}

func NewConnection(hostName string,port string) (Connection,error){

	connection := Connection{Socket:nil,Connected:false}

	sock, err := net.Dial("tcp", hostName+":"+port)

	if err != nil {
		fmt.Println(err)
		return connection,err
	}

	connection.Connected = true
	connection.Socket = sock

	fmt.Printf("Connection established between %s and localhost.\n", hostName)
	fmt.Printf("Remote Address : %s \n", sock.RemoteAddr().String())
	fmt.Printf("Local Address : %s \n", sock.LocalAddr().String())

	return connection,nil
}

func (connection Connection)SendAll(data []byte) error{
	length_data := len(data)
	cpt := 0

	for length_data > cpt{
		n,err := connection.Socket.Write(data[cpt:])

		if err != nil  {
			connection.Socket.Close()
			fmt.Println("Connection closed")
			return err
		}
		cpt += n
	}

	return nil
}