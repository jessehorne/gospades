package protocol

import (
	"github.com/BenLubar/df2014/cp437"
	"github.com/codecat/go-libs/log"
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

func NewStateDataPacket(playerID uint8) []byte {
	buf := make([]byte, 11)

	buf[0] = P_STATE_DATA

	buf[1] = playerID

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
	team1Name := cp437.Bytes("RED TEAM  ")
	buf = append(buf, team1Name...)

	// append 10 character CP437 string for team 2 name
	team2Name := cp437.Bytes("BLUE TEAM ")
	buf = append(buf, team2Name...)

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

	// base1 x - float
	buf = append(buf, leFloatBuf...)

	// base1 y - float
	buf = append(buf, leFloatBuf...)

	// base1 z - float
	buf = append(buf, leFloatBuf...)

	// base2 x - float
	buf = append(buf, leFloatBuf...)

	// base2 y - float
	buf = append(buf, leFloatBuf...)

	// base2 z - float
	buf = append(buf, leFloatBuf...)

	log.Debug("STATE DATA PACKET SIZE: ", len(buf))

	return buf
}
