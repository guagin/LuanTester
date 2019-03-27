package main

import (
	"LunaGO/server"
	"LunaGO/server/stub"
	"LunaTester/client"
	"LunaTester/handlers"
	"log"
	"net"
)

func main() {
	isServerOk := make(chan bool)
	go startServer(isServerOk)
	<-isServerOk
	// Start Client
	quit := make(chan (bool))
	for i := 0; i < 1; i++ {
		go client.New(quit)
	}
	// stop := <-quit
	if <-quit {
		log.Println("quit")
	}

}

func startServer(signalChan chan<- bool) {
	// Start Server

	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	log.Println("listen port:", "55555")
	server_1 := server.New()
	server_1.SetConnectionHandler(
		func(cIndex int32, c net.Conn) {
			defer c.Close()
			stub := stub.New(cIndex)
			stub.SetConnection(c)
			stub.Handle(0, handlers.HandlerLogin(server_1))
			stub.Start()
		},
	)
	signalChan <- true
	var connIndex int32 = 0
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept err:", err)
			break
		}
		log.Println("start accepting")
		go server_1.HandleNewConnection(connIndex, c)
		connIndex++
	}
}
