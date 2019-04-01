package models

import (
	"LunaGO/server/conn"
	"errors"
	"fmt"
	"sync"
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

type stubRepository struct {
	Stubs map[int32]*Stub
}

var newRepositoryOnce sync.Once
var repository *stubRepository

func StubRepository() *stubRepository {
	newRepositoryOnce.Do(newStubRepository) // make sure repository only get one instance.
	return repository
}

func newStubRepository() {

	repository = &stubRepository{
		Stubs: make(map[int32]*Stub),
	}
}

func (stubRepo *stubRepository) Register(ID int32, stub *Stub) {
	//TODO: add mux lock
	stubRepo.Stubs[ID] = stub
}

func (stubRepo *stubRepository) UnRegister(ID int32) {
	//TODO: add mux lock
	delete(stubRepo.Stubs, ID)
}

func (stubRepo *stubRepository) Get(ID int32) (*Stub, error) {
	stub := stubRepo.Stubs[ID]
	if stub == nil {
		return nil, errors.New(fmt.Sprintf("stub(%d) is not exist", ID))
	}
	return stub, nil
}
