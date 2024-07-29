package cmd

import (
	"fmt"
	"go-server/config"
)

// main은 cmd를 호출시킬 예정, 모든 항목들을 Cmd 스트럭쳐 안에 구성할 예정
type Cmd struct {
	config *config.Config
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	fmt.Println(c.config.Server.Port)

	return c
}
