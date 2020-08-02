package settings

import (
	"os"
	"testing"
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

func TestLoad(t *testing.T) {
	if Load("non-existing.yaml", nil) == nil {
		t.Error("expected an error for non-existing config file")
	}

	if Load("wrong.yaml", &config) == nil {
		t.Error("expected an error for deserializing into nil object")
	}

	if err := Load("test.yaml", &config); err != nil {
		t.Fatalf("failed to load test.yaml: %s", err.Error())
	}

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

func TestSave(t *testing.T) {
	config.SaveTest = "yes"

	if err := Save(".saved.yaml", &config); err != nil {
		t.Errorf("failed to save .saved.yaml: %s", err.Error())
	}

	if _, err := os.Stat(".saved.yaml"); err != nil {
		t.Error(".saved.yaml does not exist")
	}

	os.Remove(".saved.yaml")
}
