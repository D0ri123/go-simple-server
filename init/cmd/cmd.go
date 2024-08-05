package cmd

import (
	"go-server/config"
	"go-server/network"
	"go-server/repository"
	"go-server/service"
)

// main은 cmd를 호출시킬 예정, 모든 항목들을 Cmd 스트럭쳐 안에 구성할 예정
type Cmd struct {
	config  *config.Config
	network *network.Network

	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
