package protocol

import (
	"encoding/binary"
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
)

func SendMapStart(ev enet.Event, mapSize []byte) {
	// make map start packet
	packet := NewMapStartPacket(mapSize)

	log.Debug("[PACKET] Sending Map Start. Packet size is %d bytes.", len(packet))

	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
}

func SendMapChunk(ev enet.Event, compressedMap []byte, mapSize []byte) {
	// build packet
	packet := NewMapChunkPacket(compressedMap)

	data := binary.BigEndian.Uint32(mapSize)

	log.Debug("[PACKET] Sending Map Chunk. Total Little Endian Compressed Size: %d", data)

	// send packet
	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
}
