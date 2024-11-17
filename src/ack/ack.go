package ack

import (
	"log"
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

type Acknowledge struct {
	Ready   bool   `msgpack:"ready"` // send message to sender about status
	Message string `msgpack:"msg"`   // test message
}

func ReceiveAck(conn net.Conn) (Acknowledge, error) {
	ackBytes := make([]byte, 1024)

	_, err := conn.Read(ackBytes)
	if err != nil {
		log.Printf("error reading ack from conn")
		return Acknowledge{}, err
	}

	var ack Acknowledge

	err = msgpack.Unmarshal(ackBytes, &ack)
	if err != nil {
		log.Printf("error unmarshaling ack")
		return Acknowledge{}, nil
	}

	return ack, nil
}
