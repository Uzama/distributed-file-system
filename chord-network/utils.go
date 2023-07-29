package main

import (
	"log"
)

// to enable Debug, set this to true
const isDebugEnable = true

// Used for debug printing
func DebugPrintf(format string, a ...interface{}) (n int, err error) {
	if isDebugEnable {
		log.Printf(format, a...)
	}
	return
}
