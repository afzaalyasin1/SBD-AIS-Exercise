package main

import (
	"exc8/client"
	"exc8/server"
	"fmt"
	"time"
)

func main() {
	go func() {
		// todo start server
		err := server.StartGrpcServer()
		if err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	// Wait a moment for server to start
	time.Sleep(1 * time.Second)

	// todo start client
	c, err := client.NewGrpcClient()
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}

	err = c.Run()
	if err != nil {
		fmt.Println("Client error:", err)
	}

	println("Orders complete!")
}
