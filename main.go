package main

import "fmt"
import packets "CCP/Packets"


// Car Communication Protocol
func main() {

	payload := []byte("blablablba blablbal done")

	header := packets.NewHeader()

	fmt.Println(payload)
	fmt.Println(header)

}
