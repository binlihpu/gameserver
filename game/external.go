package game

import (
	"github.com/binlihpu/gameserver/game/internal"
)

var (
	//Module 实例化game模块，用于实现游戏逻辑
	Module = new(internal.Module)
	//ChanRPC 暴露ChanRPC，用于模块间通信
	ChanRPC = internal.ChanRPC
)
