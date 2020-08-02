package settings

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Load will deserialize the given yaml file into the given object
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

// Save will serialize the given object into the given yaml file
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
