package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func startClient(server string, delay string) {

	udpAddr, err := net.ResolveUDPAddr("udp4", server)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to server @ ", udpAddr)

	c, err := net.DialUDP("udp4", nil, udpAddr)

	if err != nil {
		panic(err)
	}

	sleepDuration, err := time.ParseDuration(delay)
	if err != nil {
		fmt.Println("Error parsing delay duration", delay)
		return
	}
	var seqNo int64 = 0
	for {

		err := binary.Write(c, binary.LittleEndian, seqNo)

		if err != nil {
			panic(err)
		}
		time.Sleep(sleepDuration)
		seqNo++

	}
}
