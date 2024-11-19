package receiver

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func ReceiveFile(conn net.Conn) {
	var fileSize int64

	binary.Read(conn, binary.LittleEndian, &fileSize)
	log.Printf("fileSize %d", fileSize)

	file, err := os.Create("amogus")
	if err != nil {
		log.Printf("error creating file %v", err)
		return
	}

	_, err = io.CopyN(file, conn, fileSize)
	if err != nil {
		log.Printf("error writing to file %v", err)
		return
	}

	file.Close()
	fmt.Println("transfer completed")
	conn.Close()
	return
}
