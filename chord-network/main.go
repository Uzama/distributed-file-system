package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	// get node ip and port
	arg := os.Args[1:]

	if len(arg) != 2 {
		log.Fatalln("invalid arguments")
	}
	
	address := strings.Split(arg[0], ":")
	username := arg[1]

	if len(address) != 2 {
		log.Fatalln("invalid arguments")
	} 

	if address[0] != "localhost" {
		log.Fatalln("invalid ip address")
	}

	port, err := strconv.Atoi(address[1])
	if err != nil {
		log.Fatalf("invalid arguments: %s\n", err.Error())
	}

	if port < 1000 || port > 60000 {
		log.Fatalln("invalid port")
	}
	
	// register with bootstrap server
	ip := address[0]
	portLength := len(address[1])
	
	b := NewBootstrapConn(ip, port, portLength, username)

	joinNode, err := b.RegisterWithBootstrapServer()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// join with chord network or create new network
	connection := arg[0]

	cfg := defaultConfig()

	h, err := newChord(cfg, connection, joinNode, username, b)
	if err != nil {
		log.Fatalln(err)
	}

	// strat the background rargoutines
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	<-c

	h.leave(context.Background())

	os.Exit(0)

}