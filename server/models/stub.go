package models

import (
	"LunaGO/server/conn"
	"errors"
	"fmt"
)

// Stub to keep the infomation like connection..etc.
type Stub struct {
	ID         int32
	connection *conn.Connection
}

// New will return a new Stub
func New(ID int32, connection *conn.Connection) *Stub {
	instance := &Stub{
		ID:         ID,
		connection: connection,
	}
	return instance
}

type StubRepository struct {
	Stubs map[int32]*Stub
}

func NewStubRepository() *StubRepository {
	instance := &StubRepository{
		Stubs: make(map[int32]*Stub),
	}
	return instance
}

func (stubRepo *StubRepository) register(ID int32, stub *Stub) {
	//TODO: add mux lock
	stubRepo.Stubs[ID] = stub
}

func (stubRepo *StubRepository) UnRegister(ID int32) {
	//TODO: add mux lock
	delete(stubRepo.Stubs, ID)
}

func (stubRepo *StubRepository) Get(ID int32) (*Stub, error) {
	stub := stubRepo.Stubs[ID]
	if stub == nil {
		return nil, errors.New(fmt.Sprintf("stub(%d) is not exist", ID))
	}
	return stub, nil
}
