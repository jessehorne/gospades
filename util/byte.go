package util

import (
	"encoding/binary"
	"math"
)

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
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}
