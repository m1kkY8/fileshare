package receiver

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func ReceiveFile(conn net.Conn) {
	file, err := os.Create("amogus")
	if err != nil {
		log.Printf("error creating file %v", err)
		return
	}

	_, err = io.Copy(file, conn)
	if err != nil {
		log.Printf("error writing to file %v", err)
		return
	}

	file.Close()
	fmt.Println("transfer completed")
	conn.Close()
	return
}
