package settings

import (
	"testing"
	"os"
)

var config struct {
	Test struct {
		KeyA string
		KeyB string
	}

	Foo []int

	Bar struct {
		Foo struct {
			Bar string
		}
	}

	SaveTest string
}

func TestLoadConfig(t *testing.T) {
	LoadConfig("test.yaml", &config)

	if config.Test.KeyA != "valuea" {
		t.Error("config.Test.KeyA is not \"valuea\"")
	}

	if config.Test.KeyB != "valueb" {
		t.Error("config.Test.KeyB is not \"valueb\"")
	}

	if len(config.Foo) != 5 {
		t.Error("config.Foo does not contain 5 items")
	}

	if config.SaveTest != "no" {
		t.Error("config.SaveTest is not \"no\"")
	}
}

func TestSaveConfig(t *testing.T) {
	config.SaveTest = "yes"

	if !SaveConfig(".saved.yaml", &config) {
		t.Error("SaveConfig returned false")
	}

	_, err := os.Stat(".saved.yaml")
	if err != nil {
		t.Error(".saved.yaml does not exist")
	}

	os.Remove(".saved.yaml")
}
