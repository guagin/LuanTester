package handlers

import (
	"LunaGO/server/interfaces"
	"LunaTester/server/models"
	"bytes"
	"encoding/binary"
	"log"
)

func HandleLogin(server interfaces.Server) func([]byte) []byte {
	// clientReposiory :=
	stubRepository := models.StubRepository()
	return func(packet []byte) []byte {

		var clientID int32
		reader := bytes.NewReader(packet)
		err := binary.Read(reader, binary.LittleEndian, &clientID)
		if err != nil {
			log.Println("login parsing clientID failed:", err.Error())
		}
		log.Printf("player(%d) login", clientID)
		// TODO: maybe move this to the block that server accept connection.
		stubRepository.Register(clientID, stub)
		buf := bytes.NewBuffer([]byte{})
		binary.Write(buf, binary.LittleEndian, int32(0))
		return nil
	}
}
