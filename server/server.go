package server

import (
	"LunaGO/server"
	"LunaGO/server/conn"
	"LunaGO/server/stub"
	"LunaTester/client"
	"LunaTester/server/handlers"
	"log"
	"net"
)

// Start a server
func Start(inited chan<- bool) {
	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	log.Println("listen port:", "55555")
	server_1 := server.New()
	server_1.SetConnectionHandler(
		func(cIndex int32, c *conn.Connection) {
			stub := stub.New(cIndex)
			stub.SetConnection(c)
			stub.Handle(0, handlers.HandleLogin(server_1))
			stub.Handle(1, handlers.HandleClose(server_1))
			stub.SetProcess(func(packet []byte) {
				code, err := client.GetMessageCode(packet)
				if err != nil {
					return
				}
				handler := stub.GetHandler(code)
				if handler == nil {
					log.Println("handler not exist:", code)
					return
				}
				res := handler(client.GetData(packet))

				stub.Send(res)
			})
			stub.Start()
		},
	)

	inited <- true

	var connIndex int32 = 0
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept err:", err)
			break
		}
		log.Println("start accepting")
		customeConnection := conn.NewConnection(c)
		go server_1.HandleNewConnection(connIndex, customeConnection)
		connIndex++
	}

}
