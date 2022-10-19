package main

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"testing"
)

func TestServer_DummyClient(t *testing.T) {
	enet.Initialize()

	client, err := enet.NewHost(nil, 1, 1, 0, 0)
	if err != nil {
		log.Error("Couldn't create host")
		t.Error(err)
	}

	peer, err := client.Connect(enet.NewAddress("127.0.0.1", 3030), 1, 0)
	if err != nil {
		log.Error("Couldn't connect")
		t.Error(err)
	}

	for true {
		ev := client.Service(0)

		eventType := ev.GetType()

		if eventType == enet.EventNone {
			continue
		}

		switch eventType {
		case enet.EventConnect:
			log.Info("Connected to server")
			peer.SendBytes([]byte("ping"), 0, enet.PacketFlagReliable)
		case enet.EventDisconnect:
			log.Info("Lost connection to server")
		case enet.EventReceive:
			packet := ev.GetPacket()
			packetBytes := packet.GetData()
			log.Info("Received %d bytes from server", len(packetBytes))

			if string(packetBytes) == "pong" {
				log.Info("Server sent back PONG.")
			}

			if string(packetBytes) == "disconnect" {
				log.Info("Server wants us to disconnect. Ending loop and cleaning up.")
				return
			}
			packet.Destroy()
		}
	}

	client.Destroy()
	enet.Deinitialize()
}
