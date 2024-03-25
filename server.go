package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1480")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		clientconn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleconnection(clientconn)
	}

}

func handleconnection(client net.Conn) {
	var size uint32
	err := binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	bytmsg := make([]byte, size)
	_, err = client.Read(bytmsg)
	textmessage := string(bytmsg)
	fmt.Println("Message : " + textmessage)
}
