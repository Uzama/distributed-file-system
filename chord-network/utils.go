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

// if error log the error
func logOnError(source string, err error) {
	if err != nil {
		log.Println(source, ": ", err)
	}
}

// if error, log and abort the program
func logAndExitOnError(source string, err error) {
	if err != nil {
		log.Fatal(source+": ", err)
	}
}
