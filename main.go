package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(c net.Conn, num int) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		_, _ = c.Write([]byte("you're " + string(num)))
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

	connections := 0

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		connections++
		go handleConnection(c, connections)
	}
}
