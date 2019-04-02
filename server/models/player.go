package models

import (
	"errors"
	"fmt"
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

type PlayerRepository struct {
	players map[string]*Player
}

func NewPlayerRepository() *PlayerRepository {
	instance := &PlayerRepository{
		players: make(map[string]*Player),
	}
	return instance
}

func (repo *PlayerRepository) Get(id string) (*Player, error) {
	p := repo.players[id]
	if p == nil {
		return nil, errors.New(fmt.Sprintf("playerId(%s) not found", id))
	}
	return p, nil
}

//TODO: check if player is initialed
func (repo *PlayerRepository) Register(player *Player) {
	repo.players[player.ID] = player
}

func (repo *PlayerRepository) UnRegister(ID string) {
	delete(repo.players, ID)
}
