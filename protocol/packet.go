package protocol

import (
	"fmt"
	"github.com/BenLubar/df2014/cp437"
	"github.com/jessehorne/gospades/game"
	"github.com/jessehorne/gospades/util"
)

func NewMapStartPacket(mapSize []byte) []byte {
	buf := make([]byte, 1)
	buf[0] = P_MAP_START

	buf = append(buf, mapSize...)

	return buf
}

func NewMapChunkPacket(compressedMap []byte) []byte {
	buf := make([]byte, 1)
	buf[0] = P_MAP_CHUNK

	buf = append(buf, compressedMap...)

	return buf
}

func NewStateDataPacket(playerID uint8, gs *game.State) []byte {
	buf := make([]byte, 11)

	buf[0] = P_STATE_DATA

	buf[1] = playerID
	fmt.Println("PLAYER ID", buf[1], playerID)

	buf[2] = uint8(0) // fog blue
	buf[3] = uint8(0) // fog green
	buf[4] = uint8(0) // fog red

	buf[5] = uint8(255) // team 1 blue
	buf[6] = uint8(0)   // team 1 green
	buf[7] = uint8(0)   // team 1 red

	buf[8] = uint8(0)    // team 2 blue
	buf[9] = uint8(0)    // team 2 green
	buf[10] = uint8(255) // team 2 red

	// append 10 character CP437 string for team 1 name
	team1Name := cp437.Bytes(gs.Config.Team1Name)
	buf = append(buf, team1Name...)
	// padding if necessary
	if len(team1Name) < 10 {
		buf = append(buf, make([]byte, 10-len(team1Name))...)
	}

	// append 10 character CP437 string for team 2 name
	team2Name := cp437.Bytes(gs.Config.Team2Name)
	buf = append(buf, team2Name...)
	// padding if necessary
	if len(team2Name) < 10 {
		buf = append(buf, make([]byte, 10-len(team2Name))...)
	}

	buf = append(buf, uint8(0)) // mode

	// should be 31 bytes total from here

	// CTF MODE DETAILS
	// team 1 score - byte
	buf = append(buf, uint8(0))

	// team 2 score - byte
	buf = append(buf, uint8(0))

	// capture limit - byte
	buf = append(buf, uint8(10))

	// intel flags - byte
	buf = append(buf, uint8(0))

	// 0 float values for following location data for testing
	leFloatBuf := make([]byte, 4)
	leFloatBuf[0] = uint8(0)
	leFloatBuf[1] = uint8(0)
	leFloatBuf[2] = uint8(0)
	leFloatBuf[3] = uint8(0)

	// flag x - float
	buf = append(buf, leFloatBuf...)

	// flag y - float
	buf = append(buf, leFloatBuf...)

	// flag z - float
	buf = append(buf, leFloatBuf...)

	// flag2 x - float
	buf = append(buf, leFloatBuf...)

	// flag2 y - float
	buf = append(buf, leFloatBuf...)

	// flag2 z - float
	buf = append(buf, leFloatBuf...)

	// tent 1 x - float
	buf = append(buf, leFloatBuf...)

	// tent 1 y - float
	buf = append(buf, leFloatBuf...)

	// tent 1 z - float
	buf = append(buf, leFloatBuf...)

	// tent 2 x - float
	buf = append(buf, leFloatBuf...)

	// tent 2 y - float
	buf = append(buf, leFloatBuf...)

	// tent 2 z - float
	buf = append(buf, leFloatBuf...)

	return buf
}

func NewCreatePlayerPacket(p *game.Player) []byte {
	buf := make([]byte, 4)

	buf[0] = P_CREATE_PLAYER
	buf[1] = p.ID
	buf[2] = p.Weapon
	buf[3] = p.Team

	// x position
	buf = append(buf, util.Float32ToBytes(p.Position.X)...)

	// y position
	buf = append(buf, util.Float32ToBytes(p.Position.Y)...)

	// z position
	buf = append(buf, util.Float32ToBytes(p.Position.Z)...)

	// name
	buf = append(buf, cp437.Bytes(p.Username)...)

	return buf
}

func NewWorldUpdatePacket(gs *game.State) []byte {
	buf := make([]byte, 1)
	buf[0] = P_WORLD_UPDATE

	for i := uint8(0); i < gs.Config.MaxPlayers; i++ {
		p, exists := gs.Players[i]
		if !exists {
			buf = append(buf, make([]byte, 24)...) // append 24 bytes (x y z ox oy oz) of padding because this player doesn't exist
			continue
		}

		buf = append(buf, util.Float32ToBytes(p.Position.X)...)
		buf = append(buf, util.Float32ToBytes(p.Position.Y)...)
		buf = append(buf, util.Float32ToBytes(p.Position.Z)...)

		buf = append(buf, util.Float32ToBytes(p.Orientation.X)...)
		buf = append(buf, util.Float32ToBytes(p.Orientation.Y)...)
		buf = append(buf, util.Float32ToBytes(p.Orientation.Z)...)
	}

	return buf
}

func NewSetHPPacket(p *game.Player, damageType uint8) []byte {
	var buf []byte
	buf = append(buf, p.Health)
	buf = append(buf, damageType)
	buf = append(buf, util.Float32ToBytes(p.Position.X)...)
	buf = append(buf, util.Float32ToBytes(p.Position.Y)...)
	buf = append(buf, util.Float32ToBytes(p.Position.Z)...)
	return buf
}
