package game

import (
	"github.com/codecat/go-enet"
	"github.com/jessehorne/gospades/util"
)

type Player struct {
	Username                     string
	Position                     util.Vec3Float
	Orientation                  util.Vec3Float
	ExistingPlayerPacketReceived bool
	IP                           string
	ID                           uint8
	Team                         uint8
	Weapon                       uint8
	Held                         uint8
	Kills                        uint32
	BlockColor                   util.Color
	Peer                         enet.Peer
}

func NewPlayer(peer enet.Peer, username string, id uint8, ip string) Player {
	return Player{
		Username:                     username,
		Position:                     util.NewVec3Float(0.0, 0.0, 0.0),
		Orientation:                  util.NewVec3Float(0.0, 0.0, 0.0),
		ExistingPlayerPacketReceived: false,
		IP:                           ip,
		ID:                           id,
		Team:                         255,
		Weapon:                       0,
		Held:                         0,
		Kills:                        0,
		BlockColor:                   util.NewColor(0, 0, 0),
		Peer:                         peer,
	}
}
