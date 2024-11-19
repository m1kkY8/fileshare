package main

import (
	"fileshare-client/src/ack"
	"fileshare-client/src/config"
	"fileshare-client/src/entity"
	"fileshare-client/src/receiver"
	"fileshare-client/src/sender"
	"log"
	"net"
)

func main() {
	config := config.LoadConfig()

	// connect to server
	conn, err := net.Dial("tcp", config.TCPAddr)
	if err != nil {
		log.Printf("error connecting to server %v", err)
	}

	// Make handshake
	handshake := entity.CreateHandshake(config)

	// Marshal and send handshake

	// send file
	if config.Intent == "s" {
		fileSize, err := sender.GetSize(handshake.FileName)
		handshake.FileSize = fileSize

		log.Println(fileSize)

		handshakeBytes, err := entity.MarshalHandshake(handshake)
		if err != nil {
			log.Printf("error marshaling handshake %v", err)
		}

		err = entity.SendHandshake(conn, handshakeBytes)
		if err != nil {
			log.Printf("error sending handshake %v", err)
		}

		// wait for ack
		for {
			if conn != nil {
				ack, err := ack.ReceiveAck(conn)
				if err != nil {
					break
				}

				log.Println(ack.Ready)
				log.Println(ack.Message)
				break
			} else {
				return
			}
		}

		sender.SendFile(handshake.FileName, fileSize, conn)
	}

	// receive file
	if config.Intent == "r" {
		handshakeBytes, err := entity.MarshalHandshake(handshake)
		if err != nil {
			log.Printf("error marshaling handshake %v", err)
		}

		err = entity.SendHandshake(conn, handshakeBytes)
		if err != nil {
			log.Printf("error sending handshake %v", err)
		}

		receiver.ReceiveFile(conn)
	}
}
