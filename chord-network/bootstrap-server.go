package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type BootstrapConn struct {
	ip string
	port int
	portLength int 
	username string
}

func NewBootstrapConn(ip string, port int, portLength int, username string) BootstrapConn {
	return BootstrapConn{
		ip: ip,
		port: port,
		portLength: portLength,
		username: username,
	}
}

func (b BootstrapConn) RegisterWithBootstrapServer() {
	
	conn, err := net.Dial("udp", "localhost:55555")
	if err != nil {
		log.Fatalln(err)
	}

	length := 11 + len(b.ip) + b.portLength + len(b.username)
	formattedNumber := fmt.Sprintf("%04d", length)
	
	text := fmt.Sprintf("%s REG %s %d %s", formattedNumber, b.ip, b.port, b.username)

	fmt.Println(text)
	
	p :=  make([]byte, 2048)
    
    fmt.Fprintf(conn, text)

    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }

    conn.Close()

}

func (b BootstrapConn) UnRegisterWithBootstrapServer() {
	
	conn, err := net.Dial("udp", "localhost:55555")
	if err != nil {
		log.Fatalln(err)
	}

	length := 13 + len(b.ip) + b.portLength + len(b.username)
	formattedNumber := fmt.Sprintf("%04d", length)
	
	text := fmt.Sprintf("%s UNREG %s %d %s", formattedNumber, b.ip, b.port, b.username)

	fmt.Println(text)
	
	p :=  make([]byte, 2048)
    
    fmt.Fprintf(conn, text)

    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
	
    conn.Close()
}