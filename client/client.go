package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("❌ Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("✅ Connected to chatroom! Type your message below:")
	fmt.Println("Type 'exit' to quit.\n")

	for {
		fmt.Print("You: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			fmt.Println("👋 Goodbye!")
			break
		}

		var chatHistory []string
		err = client.Call("ChatServer.SendMessage", msg, &chatHistory)
		if err != nil {
			fmt.Println("⚠️ Error sending message:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range chatHistory {
			fmt.Println(m)
		}
		fmt.Println("--------------------\n")
	}
}
