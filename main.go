package main

import (
	"./ssl"
	"fmt"
	"math/rand"
	"net"
	"time"
	"encoding/binary"
	"./configuration"
)

func main() {
	initialize()

	config := configuration.Load("config.json")

	tryHandshake(ssl.NewClientRandom(), config)
}

func initialize() {
	rand.Seed(time.Now().UnixNano())
}

func tryHandshake(random ssl.ClientRandom, config configuration.Configuration) {
	helloBody := ssl.NewClientHello(config.Handshake.RecordProtocolVersion, random, config.Client.SessionID)
	helloHandshake := ssl.NewHandshake(ssl.CLIENT_HELLO, helloBody.Serialization)
	helloMessage := ssl.NewSSLPlainText(ssl.HANDSHAKE, config.Handshake.HandshakeProtocolVersion, helloHandshake.Serialization)

	getResponse("tcp", "example.com:443", helloMessage.Serialization.Serialize())
}

func getResponse(network string, address string, message []byte) {
	timeout := 30 * time.Second

	conn, err := net.Dial(network, address)
	if err != nil {
		panic(err)
	}

	binary.Write(conn, binary.BigEndian, message)

	defer conn.Close()

	response := make([]byte, 0, 65536)
	tmp := make([]byte, 65536)
	conn.SetReadDeadline((time.Now().Add(timeout)))

	n, err := conn.Read(tmp)
	if err != nil {
		panic(err)
	}
	response = append(response, tmp[:n]...)

	conn.Close();

	fmt.Printf("Response: %x", response)
}