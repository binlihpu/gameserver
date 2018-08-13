package internal

import (
	"reflect"
	"time"

	"github.com/binlihpu/gameserver/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Hello{}, handleHello)
	Oninit()
}

func Oninit() {
	log.Debug("1")

	// 定义变量 res 接收结果
	var res string

	skeleton.Go(func() {
		// 这里使用 Sleep 来模拟一个很慢的操作
		time.Sleep(1 * time.Second)

		// 假定得到结果
		res = "3"
	}, func() {
		log.Debug(res)
	})

	log.Debug("2")
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}
