package config

import (
	"github.com/niwho/logs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf Config

type Config struct {
	AppToken string `yaml:"app_token"`
	BotId    string `yaml:"bot_id"`
	Suffix   string `yaml:"suffix"`
}

func LoadConf(filePath string) error {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logs.Log(logs.F{"err": err.Error()}).Error()
		return err
	}

	err = yaml.Unmarshal([]byte(data), &Conf)
	if err != nil {
		logs.Log(logs.F{"err": err.Error()}).Error()
		return err
	}

	return nil
}
