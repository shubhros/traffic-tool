package main

import (
	"flag"
)
var serverPort string
var serverIp string
var serverMode bool
var packetGap string
var clientPort string
var clientId uint64

func init() {
	flag.StringVar(&serverPort, "l", "5301", "port on which server should listen")
	flag.StringVar(&serverPort, "p", "5301", "port on which client should connect")
	flag.BoolVar(&serverMode, "s", false, "running as server and client")
	flag.StringVar(&packetGap, "d", "1s", "delay between packets")
	flag.StringVar(&serverIp, "c", "127.0.0.1", "server ip to connect to")
	flag.Uint64Var(&clientId, "id", 1000, "client id to be sent")
}

func main() {
	flag.Parse()

	if serverMode {
		startServer(serverPort)
		return
	}
	startClient(serverIp+":"+serverPort,packetGap, clientId)
}
