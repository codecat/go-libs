package settings

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/codecat/go-libs/log"
	"os"
)

func LoadConfig(filename string, out interface{}) bool {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Couldn't read config file: %s", err.Error())
		return false
	}

	err = yaml.Unmarshal([]byte(configData), out)
	if err != nil {
		log.Fatal("Couldn't unmarshal yaml data: %s", err.Error())
		return false
	}

	log.Info("Config loaded")
	return true
}

func SaveConfig(filename string, in interface{}) bool {
	configData, err := yaml.Marshal(in)
	if err != nil {
		log.Fatal("Couldn't marshal yaml data: %s", err.Error())
		return false
	}

	err = ioutil.WriteFile(filename, configData, os.ModePerm)
	if err != nil {
		log.Fatal("Couldn't write config file: %s", err.Error())
		return false
	}

	log.Info("Config saved")
	return true
}
