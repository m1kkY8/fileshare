package main

import (
	"fileshare-client/src/ack"
	"fileshare-client/src/config"
	"fileshare-client/src/entity"
	"fileshare-client/src/receiver"
	"log"
	"net"
	"time"
)

func main() {
	config := config.LoadConfig()

	// connect to server
	connection, err := net.Dial("tcp", config.TCPAddr)
	if err != nil {
		log.Printf("error connecting to server %v", err)
	}

	// Make handshake
	handshake := entity.CreateHandshake(config)

	// Marshal and send handshake
	handshakeBytes, err := entity.MarshalHandshake(handshake)
	if err != nil {
		log.Printf("error marshaling handshake %v", err)
	}

	err = entity.SendHandshake(connection, handshakeBytes)
	if err != nil {
		log.Printf("error sending handshake %v", err)
	}

	// send file
	if config.Intent == "s" {
		for {
			if connection != nil {
				ack, err := ack.ReceiveAck(connection)
				if err != nil {
					break
				}

				log.Println(ack.Ready)
				log.Println(ack.Message)

				time.Sleep(5 * time.Second)
			} else {
				return
			}
		}

		// sender.SendFile(config.FileName, connection)
	}

	// receive file
	if config.Intent == "r" {
		receiver.ReceiveFile(connection)
	}
}
