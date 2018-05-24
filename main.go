package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func readConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()

			if text != "" {
				fmt.Printf("\b\b** %s\n> ", text)
			}

			if !ok {
				fmt.Println("Server disconnect")
				os.Exit(1)
			}
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Error connecting with Server")
		os.Exit(1)
	}

	fmt.Println("Connected with Server")

	go readConnection(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}
}
