package settings

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

func Load(filename string, out interface{}) error {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(configData), out)
	if err != nil {
		return err
	}

	return nil
}

func Save(filename string, in interface{}) error {
	configData, err := yaml.Marshal(in)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, configData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
