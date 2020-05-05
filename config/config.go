package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"time"
)

var (
	Conf              config // holds the global app config.
	defaultConfigFile = "config/config.toml"
)

type config struct {
	ReleaseMode bool   `toml:"release_mode"`
	LogLevel    string `toml:"log_level"`
	Jwt jwt
	App app

	Server server
}

type jwt struct {
	JWTSecretKey    string `toml:"jwt_secret_key"`
	ExpireMinutes time.Duration `toml:"expire_minutes"`
}

type app struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

type server struct {
	Graceful bool   `toml:"graceful"`
	Addr     string `toml:"addr"`
	DomainApi    string `toml:"domain_api"`
}


func init() {
}

// initConfig initializes the app configuration by first setting defaults,
// then overriding settings from the app config file, then overriding
// It returns an error if any.
func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	// Set defaults.
	Conf = config{
		ReleaseMode: false,
		LogLevel:    "DEBUG",
	}

	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	} else {
		log.Infof("load config from file:" + configFile)
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.New("config load err:" + err.Error())
		}
		_, err = toml.Decode(string(configBytes), &Conf)
		if err != nil {
			return errors.New("config decode err:" + err.Error())
		}
	}
	log.Infof("config data:%v", Conf)

	return nil
}

func GetLogLvl() log.Lvl {
	// DEBUG INFO WARN ERROR OFF
	switch Conf.LogLevel {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OF":
		return log.OFF
	}

	return log.DEBUG
}
