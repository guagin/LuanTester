package models

import (
	"errors"
	"fmt"
	"sync"
)

type Player struct {
	StubID     int32
	ID         string
	Name       string
	chatRoomID int64
}

func NewPlayer(stubID int32, id, name string) *Player {
	instance := &Player{
		StubID: stubID,
		ID:     id,
		Name:   name,
	}
	return instance
}

func (player *Player) ChatRoomID() int64 {
	return player.chatRoomID
}

func (player *Player) SetChatRoomID(roomID int64) {
	player.chatRoomID = roomID
}

type playerRepository struct {
	players map[string]*Player
}

var newPlayerRepoOnce sync.Once
var playerRepo *playerRepository

func PlayerRepository() *playerRepository {
	newPlayerRepoOnce.Do(new)
	return playerRepo
}

func new() {
	playerRepo = &playerRepository{
		players: make(map[string]*Player),
	}
}

func (repo *playerRepository) Get(id string) (*Player, error) {
	p := repo.players[id]
	if p == nil {
		return nil, errors.New(fmt.Sprintf("playerId(%s) not found", id))
	}
	return p, nil
}

func (repo *playerRepository) GetByStubID(id int32) (*Player, error) {
	for _, player := range repo.players {
		if player.StubID == id {
			return player, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("player not found with stubID(%d)", id))
}

//TODO: check if player is initialed
func (repo *playerRepository) Register(player *Player) {
	repo.players[player.ID] = player
}

func (repo *playerRepository) UnRegister(ID string) {
	delete(repo.players, ID)
}
