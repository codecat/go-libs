package log

import "testing"

func TestOpen(t *testing.T) {
	Open(CatDebug, CatError)

	if Config.MinLevel != CatDebug {
		t.Error("MinLevel is not CatDebug")
	}
	if Config.MaxLevel != CatError {
		t.Error("MaxLevel is not CatError")
	}
}
