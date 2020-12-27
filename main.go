package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	"github.com/go-fx-practice/httphandler"
)

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	ApplicationConfig `yaml:"application"`
}

func main() {
	conf := &Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Println(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	// slogger := logger.Sugar()

	mux := http.NewServeMux()
	httphandler.New(mux)

	http.ListenAndServe(conf.ApplicationConfig.Address, mux)
}
