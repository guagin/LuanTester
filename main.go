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
	// Start Server
	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		log.Println("listen error:", err)
		return
	}

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

	var connIndex int32 = 0
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept err:", err)
			break
		}
		go server_1.HandleNewConnection(connIndex, c)
		connIndex++
	}

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
