package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// get node ip and port


	// register with bootstrap server


	// join with chord network or create new network


	// strat the background routines

	cfg := defaultConfig()

	h, err := newChord(cfg, "localhost:8001", "")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		time.Sleep(5 * time.Second)
		fmt.Println(h.String())
	}

}