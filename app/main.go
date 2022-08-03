package main

import (
	"github.com/binlihpu/gameserver/conf"
	"github.com/binlihpu/gameserver/game"
	"github.com/binlihpu/gameserver/gate"
	"github.com/binlihpu/gameserver/login"
	"github.com/binlihpu/leaf"
	lconf "github.com/binlihpu/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.ServerConf.LogLevel
	lconf.LogPath = conf.ServerConf.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.ServerConf.ConsolePort
	lconf.ProfilePath = conf.ServerConf.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
