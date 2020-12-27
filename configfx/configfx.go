package configfx

import (
	"fmt"
	"io/ioutil"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	ApplicationConfig `yaml:"application"`
}

func ProvideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Println(err)
	}

	return &conf
}

var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
