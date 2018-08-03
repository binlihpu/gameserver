package gate

import (
	"github.com/binlihpu/gameserver/gate/internal"
)

var (
	//Module 实例化gate模块，用于和客户端接入
	Module = new(internal.Module)
)
