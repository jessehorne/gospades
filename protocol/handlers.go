package protocol

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"github.com/jessehorne/gospades/game"
	"github.com/jessehorne/gospades/util"
)

func HandleEventConnect(ev enet.Event, gs *game.State) {
	version := GetVersionString(ev.GetData())

	if version == "unknown" {
		ev.GetPeer().Disconnect(P_DISCONNECT_REASON_PROTOCOL)
		return
	}

	// create player in gamestate
	p, err := gs.AddPlayer(ev)
	if err != nil {
		ev.GetPeer().Disconnect(P_DISCONNECT_REASON_LIMIT_EXCEDED)
	}

	// send newly connected player the Map Start packet
	SendMapStart(ev, gs.CompressedMapSize)

	// send newly connected player the Map Chunk packet (includes the whole map for now :D)
	SendMapInOneChunk(ev, gs.CompressedMap)

	// send newly connected player the State Data packet to let the client know that the map is loaded and to continue on
	SendStateDataToClient(ev, gs, p.ID)
}

func HandleDisconnect(ev enet.Event, gs *game.State) {
	ip := ev.GetPeer().GetAddress().String()
	gs.RemovePlayerByIP(ip)
}

func HandlePacketExistingPlayer(ev enet.Event, gs *game.State, data []byte) {
	playerID := data[0]
	team := data[1]
	weapon := data[2]
	held := data[3]

	kills := data[4:8]

	blue := data[8]
	green := data[9]
	red := data[10]
	color := util.NewColor(red, blue, green)

	name := data[11:]

	// update with gamestate
	if err := gs.UpdatePlayer(playerID, name, team, weapon, held, kills, color); err != nil {
		log.Error("[UPDATE PLAYER] " + err.Error())
		return
	}

	//log.Debug("[EXISTING PLAYER PACKET RECEIVED] PlayerID: %d, Team: %d, Weapon: %d, Held: %d, Kills: %d, R,G,B: (%d, %d, %d), Name: %s",
	//	playerID, team, weapon, held, kills, red, blue, green, cp437.String(name))

	// after updating the players initial state, we need to send the state data to the client
	if newPlayer, err := gs.GetPlayerByIP(ev.GetPeer().GetAddress().String()); err == nil {
		SendCreatePlayerToAllPlayers(gs, newPlayer)
	}
}

func HandlePacketSetBlockColor(ev enet.Event, gs *game.State, data []byte) {
	playerID := data[0]
	blue := data[1]
	green := data[2]
	red := data[3]

	if err := gs.UpdatePlayerBlockColor(playerID, red, green, blue); err != nil {
		log.Error("[BLOCK COLOR] ", err.Error())
	}

	//log.Debug("[BLOCK COLOR] Set for player %d", playerID)
}

func HandlePacketPositionData(ev enet.Event, gs *game.State, data []byte) {
	ip := ev.GetPeer().GetAddress().String()

	p, err := gs.GetPlayerByIP(ip)
	if err != nil {
		log.Error("[POSITION DATA] Couldn't set position on server of player that doesn't exist with IP of %s", ip)
		return
	}

	p.Position.X = util.Float32FromBytes(data[0:4])
	p.Position.Y = util.Float32FromBytes(data[4:8])
	p.Position.Z = util.Float32FromBytes(data[8:12])

	//log.Debug("[POSITION DATA] Updating Player '%s' to XYZ: %d,%d,%d", p.Position.Z, p.Position.Y, p.Position.Z)
}

func HandlePacketOrientationData(ev enet.Event, gs *game.State, data []byte) {
	ip := ev.GetPeer().GetAddress().String()

	p, err := gs.GetPlayerByIP(ip)
	if err != nil {
		log.Error("[ORIENTATION DATA] Couldn't set orientation on server of player that doesn't exist with IP of %s", ip)
		return
	}

	p.Orientation.X = util.Float32FromBytes(data[0:4])
	p.Orientation.Y = util.Float32FromBytes(data[4:8])
	p.Orientation.Z = util.Float32FromBytes(data[8:12])

	//log.Debug("[ORIENTATION DATA] Updating Player '%s' to XYZ: %d,%d,%d", p.Orientation.Z, p.Orientation.Y, p.Orientation.Z)
}

func HandlePacketInputData(ev enet.Event, gs *game.State, data []byte) {
	playerID := data[0]
	keyState := data[1]

	p, err := gs.GetPlayerByID(playerID)
	if err != nil {
		log.Error("[INPUT DATA] Couldn't update key state on server of player #%d", playerID)
		return
	}

	p.KeyState = keyState

	//log.Debug("[INPUT DATA] Updating Player '%d' input data key state", playerID)
}

func HandlePacketWeaponInput(ev enet.Event, gs *game.State, data []byte) {
	playerID := data[0]
	weaponState := data[1]

	p, err := gs.GetPlayerByID(playerID)
	if err != nil {
		log.Error("[WEAPON INPUT] Couldn't update weapon state on server of player #%d", playerID)
		return
	}

	p.WeaponState = weaponState
}

func HandlePacketSetTool(ev enet.Event, gs *game.State, data []byte) {
	playerID := data[0]
	tool := data[1]

	p, err := gs.GetPlayerByID(playerID)
	if err != nil {
		log.Error("[SET TOOL] Couldn't update tool on server of player #%d", playerID)
		return
	}

	p.Tool = tool
}

func HandlePacketBlockAction(ev enet.Event, gs *game.State, data []byte) {
	//playerID := data[0]
	//actionType := data[1] // 0-build, 1-leftButtonDestroy, 2-rightButtonDestroy, 3-grenadeDestroy
	//xPos := binary.BigEndian.Uint32(data[2:6])
	//yPos := binary.BigEndian.Uint32(data[6:10])
	//zPos := binary.BigEndian.Uint32(data[10:14])

	// update block according to action in gamestate
	// TODO

	// send this action to all clients
	var newPacket []byte
	newPacket = append(newPacket, P_BLOCK_ACTION)
	newPacket = append(newPacket, data...)

	SendBlockActionToAllPlayers(gs, newPacket)
}

func HandlePacketBlockLine(ev enet.Event, gs *game.State, data []byte) {
	//playerID := data[0]
	//
	//startXPos := binary.BigEndian.Uint32(data[1:5])
	//startYPos := binary.BigEndian.Uint32(data[5:9])
	//startZPos := binary.BigEndian.Uint32(data[9:13])
	//
	//envXPos := binary.BigEndian.Uint32(data[13:17])
	//envYPos := binary.BigEndian.Uint32(data[17:21])
	//envZPos := binary.BigEndian.Uint32(data[21:25])

	// update block according to action in gamestate
	// TODO

	// send this action to all clients
	var newPacket []byte
	newPacket = append(newPacket, P_BLOCK_LINE)
	newPacket = append(newPacket, data...)

	SendBlockLineToAllPlayers(gs, newPacket)
}
