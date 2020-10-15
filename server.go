package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func startServer(port string) {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":"+port)
	if err != nil {
		panic("error resolving addr")
	}
	ln, err := net.ListenUDP("udp4", udpAddr)

	if err != nil {
		panic("error listening")
	}

	fmt.Println("Listening on", udpAddr)

	var seqNo int64
	var prevSeqNo int64 = -1
	for {
		err := binary.Read(ln, binary.LittleEndian, &seqNo)
		if err != nil {
			panic(err)
		}
		if seqNo != prevSeqNo+1 {
			fmt.Printf("[srvr] Expecting sequence %x received sequence %x\n",
				prevSeqNo+1, seqNo)
		}
		fmt.Printf("Received sequence number: %d\n", seqNo)
		prevSeqNo = seqNo

	}
}
