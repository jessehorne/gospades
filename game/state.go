package game

import (
	"encoding/binary"
	"errors"
	"github.com/BenLubar/df2014/cp437"
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"github.com/jessehorne/gospades/util"
)

var ErrPlayerExists = errors.New("player already exists")

type State struct {
	Players    map[uint8]*Player
	Name       string
	MaxPlayers uint8
	Team1Name  string
	Team2Name  string
}

func NewState(name string, maxPlayers uint8, team1Name string, team2Name string) State {
	return State{
		Players:    map[uint8]*Player{},
		Name:       name,
		MaxPlayers: maxPlayers,
		Team1Name:  team1Name,
		Team2Name:  team2Name,
	}
}

func (gs *State) GetPlayerByIP(ip string) (*Player, error) {
	for _, p := range gs.Players {
		if p.IP == ip {
			return p, nil
		}
	}

	return nil, errors.New("no player exists with that IP address")
}

func (gs *State) GetPlayerByID(id uint8) (*Player, error) {
	p, exists := gs.Players[id]
	if !exists {
		return nil, errors.New("no player exists with that ID")
	}

	return p, nil
}

func (gs *State) AddPlayer(ev enet.Event) (*Player, error) {
	peer := ev.GetPeer()
	playerIP := peer.GetAddress().String()

	// check if player already exists with address
	_, err := gs.GetPlayerByIP(playerIP)
	if err == nil {
		return nil, ErrPlayerExists
	}

	// add new player instance to list of players if not
	newPlayerID := uint8(len(gs.Players))
	newPlayer := NewPlayer(peer, "", newPlayerID, playerIP)
	gs.Players[newPlayerID] = &newPlayer

	return &newPlayer, nil
}

func (gs *State) RemovePlayerByIP(ip string) {
	p, err := gs.GetPlayerByIP(ip)
	if err != nil {
		return
	}

	log.Debug("[PLAYER DISCONNECTED] Player ID: %d", p.ID)

	delete(gs.Players, p.ID)
}

func (gs *State) UpdatePlayer(playerID uint8, name []byte, team uint8, weapon uint8, held uint8, kills []byte, color util.Color) error {
	p, err := gs.GetPlayerByID(playerID)
	if err != nil {
		return errors.New("can't update player because no player exists with that ID")
	}

	p.Username = cp437.String(name)
	p.Team = team
	p.Weapon = weapon
	p.Held = held
	p.Kills = binary.BigEndian.Uint32(kills)
	p.BlockColor = color

	return nil
}

func (gs *State) UpdatePlayerBlockColor(playerID uint8, red uint8, green uint8, blue uint8) error {
	p, err := gs.GetPlayerByID(playerID)
	if err != nil {
		return errors.New("can't update player because no player exists with that ID")
	}

	p.BlockColor.R = red
	p.BlockColor.G = green
	p.BlockColor.B = blue

	return nil
}
