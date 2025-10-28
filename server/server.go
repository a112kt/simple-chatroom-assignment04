package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// ChatServer exported so rpc can find it
type ChatServer struct {
	messages []string
}

// SendMessage appends msg and returns the whole history
func (c *ChatServer) SendMessage(msg string, reply *[]string) error {
	c.messages = append(c.messages, msg)
	*reply = c.messages
	return nil
}

func main() {
	chatServer := new(ChatServer)

	if err := rpc.Register(chatServer); err != nil {
		fmt.Println("Error registering RPC:", err)
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("🚀 Chat Server is running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
