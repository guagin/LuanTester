package handlers

import (
	"LunaGO/server/interfaces"
	"LunaGO/server/stub"
	"LunaTester/server/models"
	"bytes"
	"encoding/binary"
	"log"
)

// create a function that handle chat message
func HandleMessage(server interfaces.Server, stub *stub.Stub) func(packet []byte) []byte {
	chatRoomRepo := models.ChatRoomRepository()
	playerRepo := models.PlayerRepository()
	return func(packet []byte) []byte {
		player, err := playerRepo.GetByStubID(stub.ID())
		if err != nil {
			log.Println(err.Error())
			return nil
		}

		reader := bytes.NewReader(packet)
		var messageLength int32
		err = binary.Read(reader, binary.LittleEndian, &messageLength)
		if err != nil {
			log.Println("parsing messageLength failed:", err.Error())
			return nil
		}
		message := packet[4 : 4+messageLength]
		log.Printf("player(%s) send chat message:%s in %d\n", player.ID, message, player.ChatRoomID())

		room := chatRoomRepo.Get(player.ChatRoomID())
		// write log
		room.PushTask(func() error {
			log.Printf("write log: %s\n", message)

			return nil
		})
		// broadcast
		room.PushTask(func() error {
			log.Printf("broadcast: %s\n", message)

			return nil
		})
		return nil
	}
}
