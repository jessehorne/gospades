package protocol

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
)

func NewMapStartPacket(mapSize []byte) []byte {
	buf := make([]byte, 1)
	buf[0] = P_MAP_START

	buf = append(buf, mapSize...)

	log.Debug("Map Start Packet", buf)

	return buf
}

func NewMapChunkPacket(compressedMap []byte) []byte {
	buf := make([]byte, 1)
	buf[0] = P_MAP_CHUNK

	buf = append(buf, compressedMap...)

	log.Debug("Map Chunk Packet")

	return buf
}

func SendPacket(ev enet.Event, data []byte) {
	ev.GetPeer().SendBytes(data, 0, enet.PacketFlagReliable)
}
