package cmd

import (
	"go-server/config"
	"go-server/network"
)

// main은 cmd를 호출시킬 예정, 모든 항목들을 Cmd 스트럭쳐 안에 구성할 예정
type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	c.network.ServerStart(c.config.Server.Port)

	return c
}
