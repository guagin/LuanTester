package handlers

import (
	"LunaGO/server/interfaces"
	"LunaGO/server/stub"
	"LunaTester/server/models"
	"bytes"
	"encoding/binary"
	"log"
)

func HandleLogin(server interfaces.Server, stub *stub.Stub) func([]byte) []byte {
	playerRepo := models.PlayerRepository()
	playerRepo = models.PlayerRepository()
	return func(packet []byte) []byte {

		var clientID int32
		reader := bytes.NewReader(packet)
		err := binary.Read(reader, binary.LittleEndian, &clientID)
		if err != nil {
			log.Println("login parsing clientID failed:", err.Error())
			return nil
		}
		var playerIDLength int32
		err = binary.Read(reader, binary.LittleEndian, &playerIDLength)

		playerID := packet[8 : 8+playerIDLength]

		if err != nil {
			log.Println("login parsing playerID failed:", err.Error())
			return nil
		}
		log.Printf("player(%s) login", playerID)
		player := models.NewPlayer(clientID, "", "")
		playerRepo.Register(player)
		buf := bytes.NewBuffer([]byte{})
		binary.Write(buf, binary.LittleEndian, int32(0))
		return buf.Bytes()
	}
}
