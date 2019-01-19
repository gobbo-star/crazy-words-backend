package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var shard Shard

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		// TODO extract room or create a new one
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		sig := strings.TrimSpace(string(netData))
		if sig == "STOP" {
			break
		} else if sig == "NEW_ROOM" {
			shard.NewRoom()
		}

		_, _ = c.Write([]byte(netData))
	}
	_ = c.Close()
}

func main() {
	fmt.Println("server is starting")
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fmt.Println("server is stopped")
	defer l.Close()

	shard = NewShard()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(c)
	}
}
