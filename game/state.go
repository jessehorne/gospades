package game

import (
	"errors"
	"github.com/codecat/go-enet"
)

var ErrPlayerExists = errors.New("player already exists")

type State struct {
	Players map[string]*Player
}

func NewState() State {
	return State{
		Players: map[string]*Player{},
	}
}

func (gs *State) HasPlayerByAddress(ip string) bool {
	_, exists := gs.Players[ip]
	if !exists {
		return false
	}

	return true
}

func (gs *State) AddPlayer(ev enet.Event) (*Player, error) {
	playerIP := ev.GetPeer().GetAddress().String()

	// check if player already exists with address
	exists := gs.HasPlayerByAddress(playerIP)
	if exists {
		return nil, ErrPlayerExists
	}

	// add new player instance to list of players if not
	newPlayer := NewPlayer("", len(gs.Players), 0, 0, playerIP)
	gs.Players[playerIP] = &newPlayer

	return &newPlayer, nil
}
