package main

import (
	"github.com/binlihpu/gameserver/conf"
	"github.com/binlihpu/gameserver/game"
	"github.com/binlihpu/gameserver/gamedata"
	"github.com/binlihpu/gameserver/gate"
	"github.com/binlihpu/gameserver/login"

	"github.com/binlihpu/leaf"
	lconf "github.com/binlihpu/leaf/conf"
)

func main() {
	gamedata.LoadFile()
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
