package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var ctx  = context.Background()

func TestLatency500Nodes(t *testing.T) {
	ringsize := 512
	DefaultConfig.ringSize = 9 // this is not ringSize, this should be x, where 2^x = ringsize
	numOfNodes := 500
	jump := int(ringsize / numOfNodes)
	numOfLookup := 200

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chords := make([]*Chord, ringsize)
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 1024) // ip addresses in order
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})
	success := 0
	var err error
	for index := 1; index < numOfNodes; index++ {
		chords[index*jump], err = newChord(DefaultConfig, testAddrs[index*jump], testAddrs[0], "a" + string(rune(1)), BootstrapConn{})
		if err == nil {
			success++
		}
	}

	fmt.Printf("successfully launched %d servers\n", success)

	time.Sleep(5 * time.Second) // let it stabilize

	fmt.Println(chords[0].Print())

	for i := 0; i < numOfLookup; i++ {
		chords[0].lookup(ctx, fmt.Sprintf("%d:%f:%d", r1.Int(), r1.Float64(), r1.Int()))
		chords[0].tracerLock.Lock()
		fmt.Println(chords[0].tracer.Hops() + " " + chords[0].tracer.Latency())
		chords[0].tracerLock.Unlock()
	}
}

func TestLookup256Nodes(t *testing.T) {
	DefaultConfig.ringSize = 8
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000)
	chords := make([]*Chord, 100)
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})

	for index := 1; index < 100; index++ {
		chords[index], _ = newChord(DefaultConfig, testAddrs[index], chords[0].Ip, fmt.Sprintf("a-%d", index), BootstrapConn{})
	}

	time.Sleep(10 * time.Second)

	fmt.Println("Finger tables.....")
	for _, c := range chords {
		if c != nil {
			fmt.Println(c.String())
			c.StopFixFingers()
		}
	}

	fmt.Println()
	time.Sleep(5 * time.Second)

	fmt.Println("Start lookup")
	chords[0].lookup(ctx, chords[99].Ip)
	fmt.Println(chords[0].tracer.String())
	fmt.Println("finish lookup")
}

func TestLatency160Nodes(t *testing.T) {
	ringsize := 256
	DefaultConfig.ringSize = 8
	numOfNodes := 160
	jump := int(ringsize / numOfNodes)
	numOfLookup := 200

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chords := make([]*Chord, ringsize)
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000) // ip addresses in order
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})
	for index := 1; index < numOfNodes; index++ {
		chords[index*jump], _ = newChord(DefaultConfig, testAddrs[index*jump], testAddrs[0], "a" + string(rune(1)), BootstrapConn{})
	}

	time.Sleep(5 * time.Second) // let it stabilize

	fmt.Println(chords[0].String())

	for i := 0; i < numOfLookup; i++ {
		chords[0].lookup(ctx, fmt.Sprintf("%d:%f:%d", r1.Int(), r1.Float64(), r1.Int()))
		chords[0].tracerLock.Lock()
		fmt.Println(chords[0].tracer.Hops() + " " + chords[0].tracer.Latency())
		chords[0].tracerLock.Unlock()
	}
}

func TestLatency100Nodes(t *testing.T) {
	ringsize := 128
	DefaultConfig.ringSize = 8
	numOfNodes := 100
	jump := int(ringsize / numOfNodes)
	numOfLookup := 20

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chords := make([]*Chord, ringsize)
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000) // ip addresses in order
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})
	for index := 1; index < numOfNodes; index++ {
		chords[index*jump], _ = newChord(DefaultConfig, testAddrs[index*jump], testAddrs[0], fmt.Sprintf("a-%d", index), BootstrapConn{})
	}

	time.Sleep(5 * time.Second) // let it stabilize

	fmt.Println(chords[0].Print())

	for i := 0; i < numOfLookup; i++ {
		chords[0].lookup(ctx, fmt.Sprintf("%d:%f:%d", r1.Int(), r1.Float64(), r1.Int()))
		chords[0].tracerLock.Lock()
		fmt.Println(chords[0].tracer.Hops() + " " + chords[0].tracer.Latency())
		chords[0].tracerLock.Unlock()
	}
}

func TestLatency40Nodes(t *testing.T) {
	ringsize := 64
	DefaultConfig.ringSize = 8
	numOfNodes := 40
	jump := int(ringsize / numOfNodes)
	numOfLookup := 20

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chords := make([]*Chord, ringsize)
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000) // ip addresses in order
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})
	for index := 1; index < numOfNodes; index++ {
		chords[index*jump], _ = newChord(DefaultConfig, testAddrs[index*jump], testAddrs[0], fmt.Sprintf("a-%d", index), BootstrapConn{})
	}

	time.Sleep(5 * time.Second) // let it stabilize

	fmt.Println(chords[0].Print())

	for i := 0; i < numOfLookup; i++ {
		chords[0].lookup(ctx, fmt.Sprintf("%d:%f:%d", r1.Int(), r1.Float64(), r1.Int()))
		chords[0].tracerLock.Lock()
		fmt.Println(chords[0].tracer.Hops() + " " + chords[0].tracer.Latency())
		chords[0].tracerLock.Unlock()
	}
}

func TestLookup32Nodes(t *testing.T) {
	numBits := 32
	DefaultConfig.ringSize = 5
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000)
	chords := make([]*Chord, numBits)
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})

	for index := 1; index < numBits; index++ {
		chords[index], _ = newChord(DefaultConfig, testAddrs[index], chords[0].Ip, fmt.Sprintf("a-%d", index), BootstrapConn{})
	}

	time.Sleep(5 * time.Second) // let it stabilize

	fmt.Println("Finger tables.....")
	for _, c := range chords {
		if c != nil {
			fmt.Println(c.Print())
			c.StopFixFingers()
		}
	}
	fmt.Println()
	time.Sleep(5 * time.Second)

	fmt.Println("Start lookup")
	chords[0].lookup(ctx, chords[numBits-1].Ip)
	fmt.Println(chords[0].tracer.String())
	fmt.Println("finish lookup")
}


func TestLatency20Nodes(t *testing.T) {
	ringsize := 256
	DefaultConfig.ringSize = 8
	numOfNodes := 20
	jump := int(ringsize / numOfNodes)
	numOfLookup := 20

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chords := make([]*Chord, ringsize)
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000) // ip addresses in order
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})
	for index := 1; index < numOfNodes; index++ {
		chords[index*jump], _ = newChord(DefaultConfig, testAddrs[index*jump], testAddrs[0], fmt.Sprintf("a-%d", index), BootstrapConn{})
	}

	time.Sleep(5 * time.Second) // let it stabilize

	fmt.Println(chords[0].Print())

	for i := 0; i < numOfLookup; i++ {
		chords[0].lookup(ctx, fmt.Sprintf("%d:%f:%d", r1.Int(), r1.Float64(), r1.Int()))
		chords[0].tracerLock.Lock()
		fmt.Println(chords[0].tracer.Hops() + " " + chords[0].tracer.Latency())
		chords[0].tracerLock.Unlock()
	}
}

func TestLookup16Nodes(t *testing.T) {
	DefaultConfig.ringSize = 4
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000)
	chords := make([]*Chord, 16)
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})

	for index := 1; index < 16; index++ {
		chords[index], _ = newChord(DefaultConfig, testAddrs[index], chords[0].Ip, fmt.Sprintf("a-%d", index), BootstrapConn{})
	}

	time.Sleep(10 * time.Second) // let it stabilize

	fmt.Println("Finger tables.....")
	for _, c := range chords {
		if c != nil {
			fmt.Println(c.Print())
			c.StopFixFingers()
		}
	}
	fmt.Println()
	time.Sleep(10 * time.Second)

	fmt.Println("Start lookup")
	chords[0].lookup(ctx, chords[14].Ip)
	fmt.Println(chords[0].tracer.String())
	fmt.Println("finish lookup")
}

func TestLatency10Nodes(t *testing.T) {
	ringsize := 256
	DefaultConfig.ringSize = 8 // 2^8 = 256
	numOfNodes := 10
	jump := int(ringsize / numOfNodes)
	numOfLookup := 200

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chords := make([]*Chord, ringsize)
	testAddrs := reverseHash(DefaultConfig.ringSize, "127.0.0.1", 5000) // ip addresses in order
	chords[0], _ = newChord(DefaultConfig, testAddrs[0], "", "a", BootstrapConn{})
	for index := 1; index < numOfNodes; index++ {
		chords[index*jump], _ = newChord(DefaultConfig, testAddrs[index*jump], testAddrs[0], fmt.Sprintf("a-%d", index), BootstrapConn{})
	}

	time.Sleep(5 * time.Second) // let it stabilize

	fmt.Println(chords[0].Print())

	for i := 0; i < numOfLookup; i++ {
		chords[0].lookup(ctx, fmt.Sprintf("%d:%f:%d", r1.Int(), r1.Float64(), r1.Int()))
		chords[0].tracerLock.Lock()
		fmt.Println(chords[0].tracer.Hops() + " " + chords[0].tracer.Latency())
		chords[0].tracerLock.Unlock()
	}
}


