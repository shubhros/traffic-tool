package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const FIFO_SIZE = 30000

var seqMap map[uint64]uint32
var dupListm map[uint64][]uint32
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
	seqMap = make(map[uint64]uint32)
	dupListm = make(map[uint64][]uint32)

	var number uint64
	var seqNo uint32
	var prevSeqNo uint32
	var ok bool
	for {

		err = binary.Read(ln, binary.BigEndian, &number)
		if err != nil {
			panic(err)
		}
		id := number & 0x00000000ffffffff;
		seqNo = uint32(number >> 32);

		if prevSeqNo, ok = seqMap[id]; ok {
			if seqNo != prevSeqNo+1 {
				fmt.Printf("[srvr] Expecting sequence %x received sequence %x cliend id: %x(%d)\n",
					prevSeqNo+1, seqNo, id, int32(seqNo - prevSeqNo))
			}
			seqMap[id] = seqNo
		} else {
			seqMap[id] = seqNo
			fmt.Printf("accepting new stream with id: %x\n", id)
		}
		idx := seqNo % FIFO_SIZE
		if _, ok = dupListm[id]; ok {
			if dupListm[id][idx] == seqNo {
				fmt.Printf("[srvr] Duplicate packet received, sequence: %x client id: %x\n",
					seqNo, id)
			} else {
				dupListm[id][idx] = seqNo
			}
		} else {
			dupListm[id] = make([]uint32, FIFO_SIZE)
			dupListm[id][seqNo%FIFO_SIZE] = seqNo
		}
	}
}
