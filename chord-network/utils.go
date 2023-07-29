package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
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

// Generate random IPs
func generateUniqueIPs(baseIP string, count int) []string {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	uniqueIPs := make(map[string]bool)
	generatedIPs := make([]string, count)

	for i := 0; i < count; {
		newIP := baseIP + ":" + strconv.Itoa(randomizer.Intn(60000)+1535)
		_, exists := uniqueIPs[newIP]
		if !exists {
			uniqueIPs[newIP] = true
			generatedIPs[i] = newIP
			i++
		}
	}

	return generatedIPs
}
