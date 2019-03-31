package handlers

import (
	"LunaGO/server/interfaces"
	"bytes"
	"encoding/binary"
	"log"
)

func HandleLogin(server interfaces.Server) func([]byte) []byte {
	// clientReposiory :=
	return func(packet []byte) []byte {

		return login(packet)
	}
}

func login(packet []byte) []byte {
	var clientID int32
	reader := bytes.NewReader(packet)
	err := binary.Read(reader, binary.LittleEndian, &clientID)
	if err != nil {
		log.Println("login parsing clientID failed:", err.Error())
	}
	log.Printf("player(%d) login", clientID)
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, int32(0))
	return nil
}
