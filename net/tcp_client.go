package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	go func() {
		reader := bufio.NewReader(conn)
		for {
			response, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Read error:", err)
				return
			}
			fmt.Print("Server response: ", response)
		}
	}()

	// 发送数据
	for {
		message := fmt.Sprintf("ping at %s\n", time.Now().Format(time.RFC3339))
		conn.Write([]byte(message))
		time.Sleep(3 * time.Second)
	}
}
