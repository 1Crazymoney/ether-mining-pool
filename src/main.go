package main

import (
    "bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
    "net/http"
	"fmt"
)

const (
	Message       = "Pong"
	StopCharacter = "\r\n\r\n"
)

func SocketServer(port int) {

	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}

	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {

	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:", data)
			if isTransportOver(data) {
				break ILOOP
			}

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}

	}
	w.Write([]byte(Message))
	w.Flush()
	log.Printf("Send: %s", Message)

}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

func    main () {

    go func() {
	    port := 9091
	    SocketServer(port)			
	}()

	conn, _ := net.Dial("tcp", "127.0.0.1:9092")
	for 
	{ 
    	reader := bufio.NewReader(os.Stdin)
    	fmt.Print("Text to send: ")
    	text, _ := reader.ReadString('\n')
    	fmt.Fprintf(conn, text + "\n")
    	message, _ := bufio.NewReader(conn).ReadString('\n')
    	fmt.Print("Message from server: "+message)
  	}

    http.Handle("/", http.FileServer(http.Dir("../www")))
    http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil);
}