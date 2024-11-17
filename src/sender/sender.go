package sender

import (
	"io"
	"log"
	"net"
	"os"
)

func SendFile(filename string, conn net.Conn) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("error opening file %v", err)
	}

	defer file.Close()
	log.Println("File sent successfully")

	_, err = io.Copy(conn, file)
	if err != nil {
		log.Printf("error writing to conn %v", err)
		return
	}

	conn.Close()
	return
}
