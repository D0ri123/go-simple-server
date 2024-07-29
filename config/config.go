package config

import (
	"github.com/naoina/toml"
	"os"
)

// 서버의 기본적인 설정을 하는 파일
// env -> toml
type Config struct {
	Server struct {
		Port string
	}
}

// config를 불러올 수 있는 함수
func NewConfig(filePath string) *Config {
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
