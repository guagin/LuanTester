package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	quit := make(chan (bool))
	for i := 0; i < 10; i++ {
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
	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("foo_%d", i)
		conn.Write([]byte(data))

		time.Sleep(time.Second * 5)
	}
	quit <- true
}
