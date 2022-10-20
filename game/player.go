package game

import "github.com/jessehorne/gospades/util"

type Player struct {
	Username                     string
	Position                     util.Vec3
	Orientation                  util.Vec3
	ExistingPlayerPacketReceived bool
	IP                           string
	ID                           uint8
	Team                         uint8
	Weapon                       uint8
	Held                         uint8
	Kills                        uint32
	Color                        util.Vec3
}

func NewPlayer(username string, id uint8, ip string) Player {
	return Player{
		Username:                     username,
		Position:                     util.NewVec3(),
		Orientation:                  util.NewVec3(),
		ExistingPlayerPacketReceived: false,
		IP:                           ip,
		ID:                           id,
		Team:                         255,
		Weapon:                       0,
		Held:                         0,
		Kills:                        0,
		Color:                        util.NewVec3(),
	}
}
