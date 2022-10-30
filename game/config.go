package game

import (
	"errors"
	"fmt"
	"github.com/yuin/gopher-lua"
	"os"
)

const gamemodesPath = "./gamemodes/"

var Config = map[string]string{
	"gamemode":   "creative",
	"serverName": "gospades default name",
	"maxPlayers": "64",
	"team1Name":  "Red",
	"team2Name":  "Blue",
}

// SetConfigVariable is called from conf.lua as 'config(string string)` to override defaults
func SetConfigVariable(l *lua.LState) int {
	key := l.ToString(1)
	value := l.ToString(2)

	if len(key) == 0 || len(value) == 0 {
		fmt.Println(fmt.Sprintf("Lua Error: Error while setting '%s'...", key))
		return -1
	}

	Config[key] = value
	fmt.Println(fmt.Sprintf("[conf.lua] Set '%s'", key))

	return 0
}

func LoadConfigFromLua(gamemode string) error {
	path := gamemodesPath + gamemode
	confPath := path + "/conf.lua"

	// check if gamemode files exist
	if _, err := os.Stat(confPath); errors.Is(err, os.ErrNotExist) {
		return errors.New("conf.lua doesn't exist for that gamemode")
	}

	l := lua.NewState()
	defer l.Close()

	l.SetGlobal("config", l.NewFunction(SetConfigVariable))
	if err := l.DoFile(confPath); err != nil {
		return err
	}

	return nil
}
