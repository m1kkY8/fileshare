package config

import (
	"flag"
	"fmt"
)

type Config struct {
	TCPAddr  string
	Port     string
	Keyword  string
	Intent   string
	FileName string
}

func LoadConfig() Config {
	serverAddr := flag.String("h", "127.0.0.1", "Server addres")
	serverPort := flag.String("p", "9001", "Server port")
	keyword := flag.String("word", "amogus", "Keyword used for establishing connection")
	intent := flag.String("a", "send", "(s)end or (r)eceive")
	fileName := flag.String("file", "", "File for sending ")

	flag.Parse()

	tcpAddr := fmt.Sprintf("%s:%s", *serverAddr, *serverPort)

	return Config{
		TCPAddr:  tcpAddr,
		Port:     *serverPort,
		Keyword:  *keyword,
		Intent:   *intent,
		FileName: *fileName,
	}
}
