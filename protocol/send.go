package protocol

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"github.com/jessehorne/gospades/game"
)

func SendMapStart(ev enet.Event, mapSize []byte) {
	// make map start packet
	packet := NewMapStartPacket(mapSize)

	log.Debug("[PACKET] Sending Map Start. Packet size is %d bytes.", len(packet))

	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
}

func SendMapChunk(ev enet.Event, compressedMap []byte) {
	// build packets by 4096 bytes
	//mapSizeInt := len(compressedMap)
	//chunkCount := (mapSizeInt / 4096) + 1
	//log.Debug("Sending %d chunks...", chunkCount)
	//
	//for i := 0; i < chunkCount; i++ {
	//	var packet []byte
	//	section := i * 4096
	//	end := section + 4096
	//
	//	if i == chunkCount-1 {
	//		log.Debug("Last chunk packet... Section is %d", section)
	//		packet = NewMapChunkPacket(compressedMap[section : len(compressedMap)-1])
	//	} else {
	//		packet = NewMapChunkPacket(compressedMap[section:end])
	//	}
	//
	//	log.Debug("[PACKET] Sending Map Chunk #%d.", i)
	//	fmt.Println(len(packet))
	//
	//	// send packet
	//	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
	//}

	packet := NewMapChunkPacket(compressedMap)
	log.Debug("[PACKET] Sending Map Chunk...")

	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
}

func SendStateData(ev enet.Event, playerID uint8, gs *game.State) {
	packet := NewStateDataPacket(playerID, gs)

	log.Debug("[PACKET] Sending State Data. Total Packet Size: %d", len(packet))

	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
}
