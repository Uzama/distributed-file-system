package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// get node ip and port


	// register with bootstrap server


	// join with chord network or create new network


	// strat the background routines

	cfg := defaultConfig()

	h, err := newChord(cfg, "localhost:8004", "localhost:8001")
	if err != nil {
		log.Fatalln(err)
	}

	go func() { 
		for {
			time.Sleep(5 * time.Second)
			fmt.Println(h.Print())
		}
	} ()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	<-c

	h.leave(context.Background())

	os.Exit(0)

}