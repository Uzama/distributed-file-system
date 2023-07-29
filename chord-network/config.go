package main

import (
	"crypto/sha1"
	"hash"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark/latency"
)

var (
	//Local simulates local network.
	Local = latency.Network{Kbps: 0, Latency: 0, MTU: 0}
	//LAN simulates local area network network.
	LAN = latency.Network{Kbps: 100 * 1024, Latency: 2 * time.Millisecond, MTU: 1500}
	//WAN simulates wide area network.
	WAN = latency.Network{Kbps: 20 * 1024, Latency: 30 * time.Millisecond, MTU: 1500}
	//Longhaul simulates bad network.
	Longhaul = latency.Network{Kbps: 1000 * 1024, Latency: 200 * time.Millisecond, MTU: 9000}
)

// Config contains configs for chord instance
type Config struct {
	ip string

	stabilizeTime time.Duration
	fixFingerTime time.Duration

	Hash     func() hash.Hash
	ringSize int

	timeout time.Duration

	DialOptions []grpc.DialOption

	networkSimulator *latency.Network
}

// DefaultConfig returns a default configuration for chord
func defaultConfig() *Config {
	config := &Config{
		stabilizeTime: 50 * time.Millisecond,
		fixFingerTime: 50 * time.Millisecond,
	}

	config.Hash = sha1.New
	config.ringSize = 4 // sha1: 160 bits
	config.DialOptions = []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(10 * time.Second),
		grpc.FailOnNonTempDialError(true),
		grpc.WithInsecure(),
	}

	config.networkSimulator = &WAN

	return config
}
