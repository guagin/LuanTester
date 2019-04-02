package models

import (
	"log"
	"sync"
	"time"
)

type ChatMessage struct {
	playerID   string
	playerName string
	content    string
}

type chatRoom struct {
	id      int64
	players map[string]*Player
	logs    map[int64]*ChatMessage // key: unix time.
	tasks   chan func() error
}

func ChatRoom() *chatRoom {
	instance := &chatRoom{
		id:      time.Now().Unix(),
		players: make(map[string]*Player),
		logs:    make(map[int64]*ChatMessage, 100),
		tasks:   make(chan func() error),
	}

	go instance.process()

	return instance
}

func (room *chatRoom) ID() int64 {
	return room.id
}

func (room *chatRoom) Join(player *Player) {
	room.players[player.ID] = player
}

func (room *chatRoom) WriteLog(message string) {
	room.pushTask(func() error {
		log.Printf("write log: %s\n", message)

		return nil
	})
}

func (room *chatRoom) Braodcast(message string) {
	room.pushTask(func() error {
		log.Printf("broadcast: %s\n", message)

		return nil
	})
}

func (room *chatRoom) Close() {
	room.pushTask(func() error {
		close(room.tasks)
		return nil
	})
}

func (room *chatRoom) pushTask(task func() error) {
	room.tasks <- task
}

func (room *chatRoom) process() {
	for {
		task, ok := <-room.tasks
		if !ok {
			log.Println("chat room tasks channel close. stop")
			// room.close()
			return
		}
		err := task()
		if err != nil {
			log.Println(err.Error())
			continue
		}

	}
}

var newChatRoomOnce sync.Once
var chatRoomRepo *chatRoomRepository

type chatRoomRepository struct {
	chatRooms map[int64]*chatRoom
}

func ChatRoomRepository() *chatRoomRepository {
	newChatRoomOnce.Do(newChatRoomRepository)
	return chatRoomRepo
}

func newChatRoomRepository() {
	chatRoomRepo = &chatRoomRepository{
		chatRooms: make(map[int64]*chatRoom),
	}
}

func (repo *chatRoomRepository) Get(ID int64) *chatRoom {
	return repo.chatRooms[ID]
}

// TODO: use channel to finish this?
func (repo *chatRoomRepository) Register(room *chatRoom) {
	repo.chatRooms[room.id] = room
}

// TODO: use channel to finish this?
func (repo *chatRoomRepository) UnRegister(room *chatRoom) {
	delete(repo.chatRooms, room.id)
}

// func (room *ChatRoom)
