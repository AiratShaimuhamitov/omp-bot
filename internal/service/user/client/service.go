package client

import (
	"cmp"
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/user"
	"golang.org/x/exp/maps"
	"slices"
)

type ClientService interface {
	Describe(clientID uint64) (*user.Client, error)
	List(cursor uint64, limit uint64) ([]user.Client, error)
	Create(client user.Client) (uint64, error)
	Update(clientID uint64, client user.Client) error
	Remove(clientID uint64) (bool, error)
}

type DummyClientService struct {
	Clients map[uint64]user.Client
}

func NewDummyClientService() *DummyClientService {
	clients := make(map[uint64]user.Client)

	clients[1] = user.Client{ID: 1, FirstName: "Jason", SecondName: "Statham"}
	clients[2] = user.Client{ID: 2, FirstName: "Ryan", SecondName: "Gosling"}
	clients[3] = user.Client{ID: 3, FirstName: "Sylvester", SecondName: "Stallone"}
	clients[4] = user.Client{ID: 4, FirstName: "Matt", SecondName: "Damon"}
	clients[5] = user.Client{ID: 5, FirstName: "Bruce", SecondName: "Willis"}

	return &DummyClientService{
		Clients: clients,
	}
}

func (s *DummyClientService) Describe(clientID uint64) (*user.Client, error) {
	client, ok := s.Clients[clientID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("client with ID %d not found", clientID))
	}

	return &client, nil
}

func (s *DummyClientService) List(cursor uint64, limit uint64) ([]user.Client, error) {
	count := uint64(len(s.Clients))

	if cursor > count {
		return []user.Client{}, nil
	}

	clients := maps.Values(s.Clients)
	slices.SortFunc(clients, func(a, b user.Client) int {
		return cmp.Compare(a.ID, b.ID)
	})

	high := cursor + limit

	if high > count {
		high = count
	}

	return clients[cursor:high], nil
}

func (s *DummyClientService) Create(client user.Client) (uint64, error) {
	newID := s.createNewID()

	client.ID = newID
	s.Clients[newID] = client

	return newID, nil
}

func (s *DummyClientService) Update(clientID uint64, client user.Client) error {
	if _, exists := s.Clients[clientID]; !exists {
		return errors.New(fmt.Sprintf("client with ID %d not found", clientID))
	}

	client.ID = clientID
	s.Clients[clientID] = client

	return nil
}

func (s *DummyClientService) Remove(clientID uint64) (bool, error) {
	if _, exists := s.Clients[clientID]; !exists {
		return false, errors.New(fmt.Sprintf("client with ID %d not found", clientID))
	}

	delete(s.Clients, clientID)

	return true, nil
}

func (s *DummyClientService) createNewID() uint64 {
	return slices.Max(maps.Keys(s.Clients)) + 1
}
