package entity

import (
	"fileshare-client/src/config"
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

type Status struct {
	Status string `msgpack:"status"`
}

type Handshake struct {
	Intent   string `msgpack:"intent"`   // r for receive, s for send
	Keyword  string `msgpack:"keyword"`  // keyword used for pairing clients
	FileName string `msgpack:"filename"` // r for receive, s for send
	FileSize int64  `msgpack:"filesize"` // r for receive, s for send
}

func CreateHandshake(config config.Config) Handshake {
	handshake := Handshake{
		Intent:   config.Intent,
		Keyword:  config.Keyword,
		FileName: config.FileName,
	}

	return handshake
}

func MarshalHandshake(handshake Handshake) ([]byte, error) {
	handshakeBytes, err := msgpack.Marshal(handshake)
	if err != nil {
		return nil, err
	}

	return handshakeBytes, nil
}

func SendHandshake(conn net.Conn, handshake []byte) error {
	_, err := conn.Write(handshake)
	if err != nil {
		return err
	}

	return nil
}
