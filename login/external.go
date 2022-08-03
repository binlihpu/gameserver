package login

import (
	"github.com/binlihpu/gameserver/login/internal"
)

var (
	//Module 实例化login模块，用于游戏登录
	Module = new(internal.Module)
	//ChanRPC 暴露ChanRPC，用于模块间通信
	ChanRPC = internal.ChanRPC
)
