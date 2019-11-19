package msg

import (
	"github.com/binlihpu/leaf/network/json"
)

// var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
}

type Hello struct {
	Name string
}
