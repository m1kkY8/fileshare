package sender

import (
	"io"
	"log"
	"net"
	"os"
)

func GetSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return -1, err
	}

	return fileInfo.Size(), nil
}

func SendFile(filename string, fileSize int64, conn net.Conn) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("error opening file %v", err)
	}

	defer file.Close()
	log.Println("File sent successfully")

	_, err = io.CopyN(conn, file, fileSize)
	if err != nil {
		log.Printf("error writing to conn %v", err)
		return
	}

	conn.Close()
	return
}
