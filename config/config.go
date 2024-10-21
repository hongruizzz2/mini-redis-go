package config

import (
	"os"
)

type ServerProperties struct {
	Bind           string `cfg:"bind"`
	Port           int    `cfg:"port"`
	AppendOnly     bool   `cfg:"appendonly"`
	AppendFilename string `cfg:"appendfilename"`
	MaxCilents     int    `cfg:"maxcilents"`
	RequirePass    string `cfg:"requirepass"`
	Databases      int    `cfg:"databases"`

	Peers []string `cfg:"peers"`
	Self  string   `cfg:"self"`
}

var Properties *ServerProperties

func init() {
	// default config
	Properties = &ServerProperties{
		Bind:       "127.0.0.1",
		Port:       6379,
		AppendOnly: false,
	}
}

// SetupConfig read config file and store properties into Properties
func SetupConfig(configFilename string) {
	file, err := os.Open(configFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

}
