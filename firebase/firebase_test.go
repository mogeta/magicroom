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

func TestGet(t *testing.T) {
	ircode := Get("tv")
	if ircode != data.TVOnOff {
		t.Fail()
	}
}
