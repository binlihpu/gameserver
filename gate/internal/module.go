package internal

import (
	"github.com/binlihpu/gameserver/conf"
	"github.com/binlihpu/gameserver/game"
	"github.com/binlihpu/gameserver/msg"
	"github.com/binlihpu/leaf/gate"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.ServerConf.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.ServerConf.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.ServerConf.CertFile,
		KeyFile:         conf.ServerConf.KeyFile,
		TCPAddr:         conf.ServerConf.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
