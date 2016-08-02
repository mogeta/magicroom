package firebase

import (
	"testing"

	"github.com/BurntSushi/toml"
)

type IRData struct {
	TVOnOff string
}

var data IRData

func init() {

	toml.DecodeFile("./test_data.toml", &data)
}

func TestGetList(t *testing.T) {
	ircode := GetList()
	if ircode["tv"] != data.TVOnOff {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	ircode := Get("tv")
	if ircode != data.TVOnOff {
		t.Fail()
	}
}
