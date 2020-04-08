package main

import (
	"fmt"

	"github.com/golobby/config"
)

func main() {
	conf, err := config.New()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	connection, err := conf.Get("connection")
	if err != nil {
		fmt.Println("Error:", err)
	}

	conf.Set("connection", "ttp://user@password:server:port")
	connection, err = conf.Get("connection")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Connection", connection)

}
