package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "-s":
		startServer("5301")
	case "-c":
		startClient(os.Args[2], "100ms")
	default:
		fmt.Println("Please specify atleast -c or -s")
	}
}
