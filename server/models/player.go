package models

import (
	"errors"
	"fmt"
	"sync"
)

type Player struct {
	StubID int32
	ID     string
	Name   string
}

func NewPlayer(stubID int32, id, name string) *Player {
	instance := &Player{
		StubID: stubID,
		ID:     id,
		Name:   name,
	}
	return instance
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

//TODO: check if player is initialed
func (repo *playerRepository) Register(player *Player) {
	repo.players[player.ID] = player
}

func (repo *playerRepository) UnRegister(ID string) {
	delete(repo.players, ID)
}
