package main

import (
	"LunaTester/client"
	"LunaTester/server"
	"log"
	"sync"
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

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go startClient(quit, int32(i))
	}

	go func() {
		for {
			_, ok := <-quit
			if ok {
				wg.Done()
			} else {
				log.Println("quit channel is drain, return.")
				return
			}
		}
	}()

	wg.Wait()
	close(quit)
	log.Println("quit")
}

func startServer(signal chan<- bool) {
	// Start Server
	server.Start(signal)
}

func startClient(quit chan<- bool, ID int32) {
	c := client.New(quit, int32(ID))
	c.SendLogin()
	time.Sleep(time.Second * time.Duration(ID+1))
	c.SendClose()
}
