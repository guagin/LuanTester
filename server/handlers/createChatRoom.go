package handlers

import (
	"LunaGO/server/interfaces"
	"LunaGO/server/stub"
	"LunaTester/server/models"
	"log"
)

func HandleCreateChatRoom(server interfaces.Server, stub *stub.Stub) func([]byte) []byte {
	chatRoomRepo := models.ChatRoomRepository()
	playerRepo := models.PlayerRepository()
	return func(packet []byte) []byte {

		player, err := playerRepo.GetByStubID(stub.ID())
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		newRoom := models.ChatRoom()
		newRoom.Join(player)
		chatRoomRepo.Register(newRoom)
		player.SetChatRoomID(newRoom.ID())
		log.Printf("create new Chat Room: %d\n", newRoom.ID())
		return nil
	}
}
