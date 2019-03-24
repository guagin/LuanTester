package main

import (
	"LunaGO/server/messages"
	"log"
	"net"
	"time"
)

func main() {
	quit := make(chan (bool))
	for i := 0; i < 1; i++ {
		go startClient(quit)
	}
	// stop := <-quit
	if <-quit {
		log.Println("quit")
	}
}

func startClient(quit chan<- bool) {
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
	}
	quit <- true
}

func sendLogin(conn net.Conn) {
	login := &messages.Login{
		ID: "123456",
	}
	data, err := login.Marshal()
	if err != nil {
		log.Println("login err:", err)
	}

	messageData, err := messages.Marshal(0, data)
	conn.Write(messageData)

}
