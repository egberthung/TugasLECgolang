package main

import (
	"bufio"
	"encoding/binary"
	"net"
	"os"
	"time"
)

func main() {
	serverconn, err := net.Dial("tcp", "127.0.0.1:1480")
	if err != nil {
		panic(err)
	}
	defer serverconn.Close()
	err = serverconn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	var message string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message = scanner.Text()
	err = binary.Write(serverconn, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		panic(err)
	}
	_, err = serverconn.Write([]byte(message))
	if err != nil {
		panic(err)
	}
	deadline := time.Now().Add(5 * time.Second) // Timeout set to 5 seconds
	err = serverconn.SetDeadline(deadline)
	if err != nil {
		panic(err)
	}
}
