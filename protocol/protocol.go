package protocol

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"github.com/jessehorne/gospades/game"
	"time"
)

// Protocol information for Ace of Spades 0.75 according to (http://www.piqueserver.org/aosprotocol/protocol075.html)

const P_VERSION = 3 // specifies 0.75
const P_DISCONNECT_REASON_BANNED = 1
const P_DISCONNECT_REASON_LIMIT_EXCEDED = 2
const P_DISCONNECT_REASON_PROTOCOL = 3
const P_DISCONNECT_REASON_FULL = 4
const P_DISCONNECT_REASON_KICKED = 10

// Server from/to Client
const P_POSITION uint8 = 0               // 13 bytes
const P_ORIENTATION uint8 = 1            // 13 bytes
const P_WORLD_UPDATE uint8 = 2           // 13 bytes
const P_INPUT uint8 = 3                  // 3 bytes
const P_WEAPON_INPUT uint8 = 4           // 3 bytes
const P_SET_HP uint8 = 5                 // 15 bytes
const P_GRENADE uint8 = 6                // 30 bytes
const P_SET_TOOL uint8 = 7               // 3 bytes
const P_SET_BLOCK uint8 = 8              // 5 bytes
const P_EXISTING_PLAYER uint8 = 9        // ???
const P_SHORT_EXISTING_PLAYER uint8 = 10 // 4 bytes
const P_MOVE_OBJECT uint8 = 11           // 15 bytes
const P_CREATE_PLAYER uint8 = 12         // ???
const P_BLOCK_ACTION uint8 = 13          // 15 bytes
const P_BLOCK_LINE uint8 = 14            // 26 bytes
// const P_CTF_STATE = ??? // 52 bytes
// const P_TC_STATE = ??? // ???
const P_CHAT uint8 = 17 // ???

// Server to Client
const P_STATE_DATA uint8 = 15        // 52 bytes
const P_KILL uint8 = 16              // 5 bytes
const P_MAP_START uint8 = 18         // 5 bytes
const P_MAP_CHUNK uint8 = 19         // ???
const P_PLAYER_LEFT uint8 = 20       // 2 bytes
const P_TERRITORY_CAPTURE uint8 = 21 // 5 bytes
const P_PROGRESS_BAR uint8 = 22      // 8 bytes
const P_INTEL_CAPTURE uint8 = 23     // 3 bytes
const P_INTEL_PICKUP uint8 = 24      // 24 bytes
const P_INTEL_DROP uint8 = 25        // 14 bytes
const P_RESTOCK uint8 = 26           // 2 bytes
const P_FOG_COLOUR uint8 = 27        // 5 bytes

// Client to Server
const P_HIT uint8 = 5            // 3 bytes
const P_WEAPON_RELOAD uint8 = 28 // 4 bytes
const P_CHANGE_TEAM uint8 = 29   // 3 bytes
const P_CHANGE_WEAPON uint8 = 30 // 3 bytes
const P_MAP_CACHED uint8 = 31    // 2 bytes

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

type PacketFunction func(enet.Event, *game.State, []byte)

var PacketFuncs = map[uint8]PacketFunction{}

func PacketTypeToString(p int) string {
	v, exists := PacketTypes[p]
	if !exists {
		return ""
	}
	return v
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

func AddPacketFunction(id uint8, f PacketFunction) {
	PacketFuncs[id] = f
}

func PacketHandler(ev enet.Event, gs *game.State) {
	// Get the packet
	packet := ev.GetPacket()
	data := packet.GetData()

	// We must destroy the packet when we're done with it
	defer packet.Destroy()

	// get packet id
	packetID := data[0]

	f, exists := PacketFuncs[packetID]
	if exists {
		f(ev, gs, data[1:])
	} else {
		log.Warn("[UNKNOWN PACKET] %d", packetID)
	}
}

func Init(gs *game.State) {
	AddPacketFunction(0, HandlePacketPositionData)
	AddPacketFunction(1, HandlePacketOrientationData)
	AddPacketFunction(3, HandlePacketInputData)
	AddPacketFunction(4, HandlePacketWeaponInput)
	AddPacketFunction(7, HandlePacketSetTool)
	AddPacketFunction(8, HandlePacketSetBlockColor)
	AddPacketFunction(9, HandlePacketExistingPlayer)
	AddPacketFunction(13, HandlePacketBlockAction)
	AddPacketFunction(14, HandlePacketBlockLine)

	// send world update 10 times a second
	go func() {
		SendWorldUpdate(gs)
		time.Sleep(100 * time.Millisecond)
	}()
}
