package main

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"github.com/jessehorne/gospades/game"
	"github.com/jessehorne/gospades/protocol"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// TODO: verify that required env vars be set

	// Config Validation
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		log.Error("Incorrect `SERVER_NAME` environment variable.")
		return
	}

	maxPlayersEnv := os.Getenv("MAX_PLAYERS")
	if maxPlayersEnv == "" {
		log.Error("Incorrect `MAX_PLAYERS` environment variable.")
		return
	}
	maxPlayersInt, err := strconv.Atoi(maxPlayersEnv)
	maxPlayers := uint8(maxPlayersInt)
	if maxPlayers < 0 || maxPlayers > 64 {
		log.Error("Incorrect `MAX_PLAYERS` environment variable value. Must be between 0 and 64.")
		return
	}

	team1Name := os.Getenv("TEAM_1_NAME")
	if len(team1Name) < 1 || len(team1Name) > 10 {
		log.Error("Incorrect `TEAM_1_NAME` environment variable.")
		return
	}

	team2Name := os.Getenv("TEAM_2_NAME")
	if len(team2Name) < 1 || len(team2Name) > 10 {
		log.Error("Incorrect `TEAM_2_NAME` environment variable.")
		return
	}

	// init gamestate
	gamestate, err := game.NewState(serverName, maxPlayers, team1Name, team2Name)
	if err != nil {
		log.Error("Error initializing server", err.Error())
		return
	}

	// initialize packet handlers
	protocol.InitPacketHandlers()

	// init enet stuff
	enet.Initialize()

	host, err := enet.NewHost(enet.NewListenAddress(3030), 32, 2, 0, 0)
	if err != nil {
		log.Error("Couldn't create host: %s", err.Error())
		return
	}
	host.CompressWithRangeCoder()

	log.Info("gospades v0.0.1 started on port 3030")

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
			protocol.HandleDisconnect(ev, &gamestate)
		case enet.EventReceive: // A peer sent us some data
			protocol.PacketHandler(ev, &gamestate)
		}
	}

	// Destroy the host when we're done with it
	host.Destroy()

	// Uninitialize enet
	enet.Deinitialize()
}
