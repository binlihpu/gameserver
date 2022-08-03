package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/binlihpu/leaf/log"
)

type Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

var ServerConf *Server

func init() {
	ServerConf = &Server{}
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, ServerConf)
	if err != nil {
		log.Fatal("%v", err)
	}
	log.Release("ServerConf:%+v", ServerConf)
}
