package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	conf *Config
)

type Config struct {
	RunModel string `yaml:"runmode"`
	Port     string `yaml:"port"`
	Gormlog  bool   `yaml:"gormlog"`
	Name     string `yaml:"name"`

	DataBase DataBase    `yaml:"db"`
	Redis    RedisConfig `yaml:"redis"`
	JWT      JWTConfig   `yaml:"jwt"`
}

type DataBase struct {
	Type     string `yaml:"type"`
	DbName   string `yaml:"dbname"`
	Addr     string `yaml:"addr"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type JWTConfig struct {
	Expiration int64  `yaml:"expiration"`
	Privatekey string `yaml:"private_key"`
}

func GetConfig() *Config {
	return conf
}

func Read() *Config {
	// Read config file
	buf, err := ioutil.ReadFile("resource/config/conf.yaml")
	if err != nil {
		panic(err)
	}
	// Unmarshal yml
	conf = &Config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		panic(err)
	}
	return conf
}
