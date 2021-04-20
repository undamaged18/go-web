package main

import (
	server2 "ecommerce/server"
	"ecommerce/services/config"
	"fmt"
	"os"
)

func main() {
	intro()
	server2.Run()
}

func intro() {
	conf := config.New()
	if conf.Environment != "production" {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting Application")
		fmt.Println("Working Directory: ", path)
		fmt.Printf("Listening on: http://localhost%s\n", conf.Ports.HTTP)
		fmt.Printf("Listening on: https://localhost%s\n", conf.Ports.HTTPS)
	}
}