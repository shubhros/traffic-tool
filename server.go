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

	var number uint64
	var seqNo uint32
	var prevSeqNo uint32
	for {

		err = binary.Read(ln, binary.BigEndian, &number)
		if err != nil {
			panic(err)
		}
		id := number & 0x00000000ffffffff;
		seqNo = uint32(number >> 32);
		if seqNo != prevSeqNo+1 {
			fmt.Printf("[srvr] Expecting sequence %x received sequence %x cliend id: %x(%d)\n",
				prevSeqNo+1, seqNo, id, (seqNo - prevSeqNo))
		}
		prevSeqNo = seqNo
	}
}
