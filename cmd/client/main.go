package main

import "github.com/sonikq/gophkeeper/internal/app/client"

func main() {
	clientManager := client.New()
	for {
		clientManager.Run()
	}
}
