package game

type Player struct {
	Username                     string
	X                            int
	Y                            int
	ExistingPlayerPacketReceived bool
	IP                           string
	ID                           int
}

func NewPlayer(username string, id int, x int, y int, ip string) Player {
	return Player{
		Username:                     username,
		X:                            x,
		Y:                            y,
		ExistingPlayerPacketReceived: false,
		IP:                           ip,
		ID:                           id,
	}
}
