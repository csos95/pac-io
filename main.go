package main

import "log"

func main() {
	server := newServer("127.0.0.1:8082")

	err := server.run()
	if err != nil {
		log.Println(err)
	}
}
