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
	KeyState                     uint8
	WeaponState                  uint8
}

func NewPlayer(peer enet.Peer, username string, id uint8, ip string) Player {
	return Player{
		Username:                     username,
		Position:                     util.NewVec3Float(100.0, 100.0, -2.0),
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
		KeyState:                     0,
		WeaponState:                  0,
	}
}

func (p *Player) IsHoldingUp() bool {
	return util.GetBit(p.KeyState, util.BIT_UP)
}

func (p *Player) IsHoldingDown() bool {
	return util.GetBit(p.KeyState, util.BIT_DOWN)
}

func (p *Player) IsHoldingLeft() bool {
	return util.GetBit(p.KeyState, util.BIT_LEFT)
}

func (p *Player) IsHoldingRight() bool {
	return util.GetBit(p.KeyState, util.BIT_RIGHT)
}

func (p *Player) IsHoldingJump() bool {
	return util.GetBit(p.KeyState, util.BIT_JUMP)
}

func (p *Player) IsHoldingCrouch() bool {
	return util.GetBit(p.KeyState, util.BIT_CROUCH)
}

func (p *Player) IsHoldingSneak() bool {
	return util.GetBit(p.KeyState, util.BIT_SNEAK)
}

func (p *Player) IsHoldingSprint() bool {
	return util.GetBit(p.KeyState, util.BIT_SPRINT)
}

func (p *Player) IsFiringPrimary() bool {
	return util.GetBit(p.WeaponState, util.BIT_WEAPON_PRIMARY)
}

func (p *Player) IsFiringSecondary() bool {
	return util.GetBit(p.WeaponState, util.BIT_WEAPON_SECONDARY)
}
