package main

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"github.com/jessehorne/gospades/game"
	"github.com/jessehorne/gospades/protocol"
)

func main() {
	// init gamestate
	gamestate := game.NewState()

	// init enet stuff
	enet.Initialize()

	host, err := enet.NewHost(enet.NewListenAddress(3030), 32, 2, 0, 0)
	if err != nil {
		log.Error("Couldn't create host: %s", err.Error())
		return
	}
	host.CompressWithRangeCoder()

	log.Info("Server started on port 3030")

	// The event loop
	for true {
		// Wait until the next event
		ev := host.Service(0)

		// Do nothing if we didn't get any event
		if ev.GetType() == enet.EventNone {
			continue
		}

		switch ev.GetType() {
		case enet.EventConnect: // A new peer has connected
			protocol.HandleEventConnect(ev, &gamestate)
		case enet.EventDisconnect: // A connected peer has disconnected
			log.Info("Peer disconnected: %s", ev.GetPeer().GetAddress())

		case enet.EventReceive: // A peer sent us some data
			// Get the packet
			packet := ev.GetPacket()

			// We must destroy the packet when we're done with it
			defer packet.Destroy()

			log.Info("[PACKET]", packet.GetData())
		}
	}

	// Destroy the host when we're done with it
	host.Destroy()

	// Uninitialize enet
	enet.Deinitialize()
}
