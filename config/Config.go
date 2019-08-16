package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var FileSystem = http.StripPrefix("/static/", http.FileServer(http.Dir("client/")))
var Data Config
var Cert = "cert/cert.pem"
var Key = "cert/key.pem"

type Config struct {
	ServerPort  string
	Host        string
	DBPort      int
	DBName      string
	DBUser      string
	DBPass      string
	AdminPath   string
	UserPath    string
	CookieHash  string
	CookieBlock string
}

func (config *Config) ParseConfigFile() (*Config, error) {
	jsonConfig, err := ioutil.ReadFile("config/Configuration.json")
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(jsonConfig, &config); err != nil {
		return config, err
	}
	return config, nil
}
