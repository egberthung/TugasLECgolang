package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
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
	err := client.SetReadDeadline(time.Now().Add(10 * time.Second))
	var size uint32

	err := binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	bytmsg := make([]byte, size)
	_, err = client.Read(bytmsg)
	textmessage := string(bytmsg)
	fmt.Println("Message : " + textmessage)

	deadline := time.Now().Add(10 * time.Second) // Timeout set to 10 seconds
	err = client.SetReadDeadline(deadline)
	if err != nil {
		panic(err)
	}
}
