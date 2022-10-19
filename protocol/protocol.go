package protocol

import (
	"github.com/codecat/go-enet"
	"github.com/jessehorne/gospades/game"
	"github.com/jessehorne/gospades/util"
)

// Protocol information for Ace of Spades 0.75 according to (http://www.piqueserver.org/aosprotocol/protocol075.html)

const P_VERSION = 3 // specifies 0.75
const P_DISCONNECT_REASON_BANNED = 1
const P_DISCONNECT_REASON_LIMIT_EXCEDED = 2
const P_DISCONNECT_REASON_PROTOCOL = 3
const P_DISCONNECT_REASON_FULL = 4
const P_DISCONNECT_REASON_KICKED = 10

// Server from/to Client
const P_POSITION = 0               // 13 bytes
const P_ORIENTATION = 1            // 13 bytes
const P_WORLD_UPDATE = 2           // 13 bytes
const P_INPUT = 3                  // 3 bytes
const P_WEAPON_INPUT = 4           // 3 bytes
const P_SET_HP = 5                 // 15 bytes
const P_GRENADE = 6                // 30 bytes
const P_SET_TOOL = 7               // 3 bytes
const P_SET_BLOCK = 8              // 5 bytes
const P_EXISTING_PLAYER = 9        // ???
const P_SHORT_EXISTING_PLAYER = 10 // 4 bytes
const P_MOVE_OBJECT = 11           // 15 bytes
const P_CREATE_PLAYER = 12         // ???
const P_BLOCK_ACTION = 13          // 15 bytes
const P_BLOCK_LINE = 14            // 26 bytes
// const P_CTF_STATE = ??? // 52 bytes
// const P_TC_STATE = ??? // ???
const P_CHAT = 17 // ???

// Server to Client
const P_STATE_DATA = 15        // 52 bytes
const P_KILL = 16              // 5 bytes
const P_MAP_START uint8 = 18   // 5 bytes
const P_MAP_CHUNK = 19         // ???
const P_PLAYER_LEFT = 20       // 2 bytes
const P_TERRITORY_CAPTURE = 21 // 5 bytes
const P_PROGRESS_BAR = 22      // 8 bytes
const P_INTEL_CAPTURE = 23     // 3 bytes
const P_INTEL_PICKUP = 24      // 24 bytes
const P_INTEL_DROP = 25        // 14 bytes
const P_RESTOCK = 26           // 2 bytes
const P_FOG_COLOUR = 27        // 5 bytes

// Client to Server
const P_HIT = 5            // 3 bytes
const P_WEAPON_RELOAD = 28 // 4 bytes
const P_CHANGE_TEAM = 29   // 3 bytes
const P_CHANGE_WEAPON = 30 // 3 bytes
const P_MAP_CACHED = 31    // 2 bytes

var PacketTypes = map[int]string{
	0:  "POSITION",
	1:  "ORIENTATION",
	2:  "WORLD_UPDATE",
	3:  "INPUT",
	4:  "WEAPON_INPUT",
	5:  "SET_HP",
	6:  "GRENADE",
	7:  "SET_TOOL",
	8:  "SET_BLOCK",
	9:  "EXISTING_PLAYER",
	10: "SHORT_EXISTING_PLAYER",
	11: "MOVE_OBJECT",
	12: "CREATE_PLAYER",
	13: "BLOCK_ACTION",
	14: "BLOCK_LINE",
	15: "STATE_DATA",
	16: "KILL",
	17: "CHAT",
	18: "MAP_START",
	19: "MAP_CHUNK",
	20: "PLAYER_LEFT",
	21: "TERRITORY_CAPTURE",
	22: "PROGRESS_BAR",
	23: "INTEL_CAPTURE",
	24: "INTEL_PICKUP",
	25: "INTEL_DROP",
	26: "RESTOCK",
	27: "FOG_COLOR",
	28: "WEAPON_RELOAD",
	29: "CHANGE_TEAM",
	30: "CHANGE_WEAPON",
	31: "MAP_CACHED",
}

func PacketTypeToString(p int) string {
	v, exists := PacketTypes[p]
	if !exists {
		return ""
	}
	return v
}

func ParsePacket(b []byte) ([]byte, interface{}, error) {
	return b, nil, nil
}

func GetVersionString(d uint32) string {
	if d == 3 {
		return "0.75"
	}

	if d == 4 {
		return "0.76"
	}

	return "unknown"
}

func HandleEventConnect(ev enet.Event, gs *game.State) {
	version := GetVersionString(ev.GetData())

	if version == "unknown" {
		ev.GetPeer().Disconnect(P_DISCONNECT_REASON_PROTOCOL)
		return
	}

	// create player in gamestate
	_, err := gs.AddPlayer(ev)
	if err != nil {
		ev.GetPeer().Disconnect(P_DISCONNECT_REASON_LIMIT_EXCEDED)
	}

	// Map stuff
	compressedMap, mapSize, err := util.GetMapAndSize("./maps/cs_assault.vxl")

	// send newly connected player the Map Start packet
	SendMapStart(ev, mapSize)

	// send newly connected player the Map Chunk packet (includes the whole map for now :D)
	SendMapChunk(ev, compressedMap, mapSize)
}

func HandleDisconnect(ev enet.Event) {
	// TODO
}
