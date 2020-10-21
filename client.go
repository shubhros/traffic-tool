package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func startClient(server string, delay string, clientId uint64) {

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
	var seqNo uint32 = 0
	var number uint64
	for {
		number = uint64(seqNo) << 32 | clientId
		err := binary.Write(c, binary.BigEndian, number)
		if err != nil {
			continue
		}
		time.Sleep(sleepDuration)
		seqNo++

	}
}
