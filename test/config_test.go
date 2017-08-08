package test

import (
	"fmt"
	"testing"

	"github.com/eaciit/toolkit"

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

func TestSerde(t *testing.T) {
	serdeObj := struct {
		FullName string
		Age      int
	}{}

	cfg.Set("Age", 37)
	if err := cfg.Serde(&serdeObj); err != nil {
		t.Error(err)
	}

	fmt.Printf("Obj: %s\n", toolkit.JsonString(&serdeObj))
}
