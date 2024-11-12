package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

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

func main() {
	// get prgram flags
	serverAddr := flag.String("h", "127.0.0.1", "Server addres")
	serverPort := flag.String("p", "9001", "Server port")
	keyword := flag.String("word", "amogus", "Keyword used for establishing connection")
	intent := flag.String("a", "send", "(s)end or (r)eceive")
	fileName := flag.String("file", "", "File for sending ")

	flag.Parse()

	// parse tcp server address
	fullAdrr := fmt.Sprintf("%s:%s", *serverAddr, *serverPort)

	// connect to server
	connection, err := net.Dial("tcp", fullAdrr)
	if err != nil {
		log.Printf("error connecting to server %v", err)
	}

	// pack handshake
	handshake := Handshake{
		Intent:   *intent,
		Keyword:  *keyword,
		FileName: *fileName,
	}

	handshakeBytes, err := msgpack.Marshal(handshake)
	if err != nil {
		log.Printf("error marshaling handshake %v", err)
	}

	_, err = connection.Write(handshakeBytes)
	if err != nil {
		log.Printf("error writing to connection %v", err)
	}

	// send file
	if *intent == "s" {
		file, err := os.Open(*fileName)
		if err != nil {
			log.Printf("error opening file %v", err)
		}

		defer file.Close()
		log.Println("File sent successfully")

		_, err = io.Copy(connection, file)
		if err != nil {
			log.Printf("error writing to conn %v", err)
			return
		}

		connection.Close()
		return
	}

	// receive file
	if *intent == "r" {
		file, err := os.Create(*fileName)
		if err != nil {
			log.Printf("error creating file %v", err)
			return
		}

		_, err = io.Copy(file, connection)
		if err != nil {
			log.Printf("error writing to file %v", err)
			return
		}

		file.Close()
		fmt.Println("transfer completed")
		connection.Close()
		return
	}
}
