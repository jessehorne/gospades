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

func SendMapInManyChunks(ev enet.Event, compressedMap []byte) {
	// build packets by 4096 bytes
	mapSizeInt := len(compressedMap)
	chunkCount := (mapSizeInt / 4096) + 1

	for i := 0; i < chunkCount; i++ {
		var packet []byte
		section := i * 4096
		end := section + 4096

		if i == chunkCount-1 {
			packet = NewMapChunkPacket(compressedMap[section : len(compressedMap)-1])
		} else {
			packet = NewMapChunkPacket(compressedMap[section:end])
		}

		log.Debug("[PACKET] Sending Map Chunk #%d.", i+1)

		// send packet
		ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
	}
}

func SendMapInOneChunk(ev enet.Event, compressedMap []byte) {
	packet := NewMapChunkPacket(compressedMap)
	log.Debug("[PACKET] Sending Map Chunk...")

	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
}

func SendStateDataToClient(ev enet.Event, gs *game.State, playerID uint8) {
	packet := NewStateDataPacket(playerID, gs)

	log.Debug("[PACKET] Sending State Data. Total Packet Size: %d", len(packet))

	ev.GetPeer().SendBytes(packet, 0, enet.PacketFlagReliable)
}

func SendCreatePlayerToAllPlayers(gs *game.State, newPlayer *game.Player) {
	packet := NewCreatePlayerPacket(newPlayer)

	for p := range gs.Players {
		player := gs.Players[p]
		player.Peer.SendBytes(packet, 0, enet.PacketFlagReliable)
	}

	log.Debug("[PACKET] Sending Create Player to %d players.", len(gs.Players))
}

func SendBlockActionToAllPlayers(gs *game.State, packet []byte) {
	for p := range gs.Players {
		player := gs.Players[p]
		player.Peer.SendBytes(packet, 0, enet.PacketFlagReliable)
	}

	log.Debug("[BROADCAST BLOCK ACTION] Sending Block Action to %d players.", len(gs.Players))
}

func SendBlockLineToAllPlayers(gs *game.State, packet []byte) {
	for p := range gs.Players {
		player := gs.Players[p]
		player.Peer.SendBytes(packet, 0, enet.PacketFlagReliable)
	}

	log.Debug("[BROADCAST BLOCK LINE] Sending Block Line to %d players.", len(gs.Players))
}

func SendWorldUpdate(gs *game.State) {
	// build world update packet
	packet := NewWorldUpdatePacket(gs)

	// send to all players
	for p := range gs.Players {
		player := gs.Players[p]
		player.Peer.SendBytes(packet, 0, enet.PacketFlagUnreliableFragment)
	}
}

func SendSetHP(p *game.Player, damageType uint8) {
	packet := NewSetHPPacket(p, damageType)
	p.Peer.SendBytes(packet, 0, enet.PacketFlagReliable)
}
