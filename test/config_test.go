package test

import (
	"fmt"
	"testing"

	. "github.com/eaciit/appconfig"
)

const pathtofile = "/users/ariefdarmawan/Temp/config.json"
const nameTest = "EACIIT - Leveraging Value"

var cfg = new(Config)

func TestCfgWrite(t *testing.T) {
	cfg.SetConfigFile(pathtofile)
	cfg.Set("FullName", nameTest)
	if e := cfg.Write(); e != nil {
		t.Error(e.Error())
	}
}

func TestCfgGet(t *testing.T) {
	cfg = new(Config)
	cfg.SetConfigFile(pathtofile)
	s := cfg.Get("FullName").(string)
	if s != nameTest {
		t.Error("Unable to read value. Expected '" + nameTest + "' got '" + s + "'")
	} else {
		fmt.Printf("Config: %s\n", s)
	}
}
