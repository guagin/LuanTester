package main

import (
	"LunaTester/client"
	"LunaTester/server"
	"log"
	"time"
)

func main() {
	isServerOk := make(chan bool)
	// start server
	go startServer(isServerOk)
	// wait until server is ok.
	<-isServerOk

	// start client
	quit := make(chan bool, 1)
	// quit <- true
	for i := 0; i < 3; i++ {
		go startClient(quit, int32(i))
	}

	log.Println("Test")
	if <-quit {
		log.Println("quit")
	}
}

func startServer(signal chan<- bool) {
	// Start Server
	server.Start(signal)
}

func startClient(quit chan<- bool, ID int32) {
	c := client.New(quit, int32(ID))
	c.SendLogin()
	time.Sleep(time.Second * 1)
	c.SendClose()
}
