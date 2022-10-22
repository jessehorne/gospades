package util

import (
	"encoding/binary"
	"math"
)

type bit byte

const (
	BIT_SPRINT bit = 0b0000_0001
	BIT_SNEAK  bit = 0b0000_0010
	BIT_CROUCH bit = 0b0000_0100
	BIT_JUMP   bit = 0b0000_1000
	BIT_RIGHT  bit = 0b0001_0000
	BIT_LEFT   bit = 0b0010_0000
	BIT_DOWN   bit = 0b0100_0000
	BIT_UP     bit = 0b1000_0000

	BIT_WEAPON_PRIMARY   bit = 0b0000_0001
	BIT_WEAPON_SECONDARY bit = 0b0000_0010
)

func GetBit(b byte, n bit) bool {
	return b&byte(n) > 0
}

func ReverseBytes(b []byte) []byte {
	reversed := make([]byte, len(b))
	count := len(b) - 1
	for i := 0; i < len(b); i++ {
		reversed[count] = b[i]
		count -= 1
	}

	return reversed
}

func Float32FromBytes(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}
