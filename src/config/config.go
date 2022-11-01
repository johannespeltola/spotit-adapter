package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/guregu/null.v3"
)

var cfg Config

// Config stores configuration values read from /config/config.json
type Config struct {
	Version       null.String         `json:"version"`
	Title         null.String         `json:"title"`
	LogPath       null.String         `json:"logPath"`
	LoggerConf    LoggerConfiguration `json:"loggerConf"`
	EntsoeURL     null.String         `json:"entsoeURL"`
	SecurityToken null.String         `json:"securityToken"`
	DocumentType  null.String         `json:"documentType"`
	Domain        null.String         `json:"domain"`
	VATAmount     null.Float          `json:"vatAmount"`
	DataEndpoint  null.String         `json:"dataEndpoint"`
}

// ReadConfigJSON was created with the to make the configuration process
// available in JSON format. The reason was to extend the decoder capability
// to distiguish between valid zero value and unset value to avoid
// using unintended zero/default value on the config instances.
// This is critical as YAML default decoder can not distinguish between
// intentionally set null/default value and missing value.
//
// If the config value is not found in the config files, the getter should
// return the default hard-coded value resides inside configuration file or
// the environment value the system is running agaisnt instead.
func ReadConfigJSON() {
	f, err := os.Open("./config/config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
}

func GetConf() *Config {
	return &cfg
}

func GetLogPath() string {
	return cfg.LogPath.String
}

func GetEntsoeBase() string {
	return fmt.Sprintf("%v?securityToken=%v&documentType=%v&in_Domain=%v&out_Domain=%v", cfg.EntsoeURL.String, cfg.SecurityToken.String, cfg.DocumentType.String, cfg.Domain.String, cfg.Domain.String)
}

func GetVat() float32 {
	return float32(cfg.VATAmount.Float64)
}

func GetDataEndpoint() string {
	return cfg.DataEndpoint.String
}
