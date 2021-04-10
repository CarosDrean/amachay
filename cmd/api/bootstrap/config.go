package bootstrap

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/gommon/log"

	"github.com/CarosDrean/amachay/kit/db"
)

type Database struct {
	Name db.DB `json:"name"`
}

type Configuration struct {
	Port     int      `json:"port"`
	Database Database `json:"database"`
}

func newConfiguration(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	conf := Configuration{}
	if err := json.Unmarshal(file, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
