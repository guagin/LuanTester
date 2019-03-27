package client

import (
	"LunaGO/server/messages"
	TesterMessage "LunaTester/messages"
	"log"
	"net"
	"time"
)

func New(quit chan<- bool) {
	log.Println("begin Dial")
	conn, err := net.Dial("tcp", ":55555")
	if err != nil {
		log.Println("dial error:", err)
	}
	defer conn.Close()
	log.Println("Dial ok")

	for i := 0; i < 1; i++ {
		sendLogin(conn)
		time.Sleep(time.Second * 5)
		sendClose(conn)
		time.Sleep(time.Second * 5)
	}
	quit <- true
}

func sendLogin(conn net.Conn) {
	login := &TesterMessage.Login{
		ID: "123456",
	}
	data, err := login.Marshal()
	if err != nil {
		log.Println("login err:", err)
	}

	messageData, err := messages.Marshal(0, data)
	conn.Write(messageData)
}

func sendClose(conn net.Conn) {
	messageData, err := messages.Marshal(-1, nil)
	if err != nil {
		log.Println("close err:", err)
	}
	conn.Write(messageData)
}
