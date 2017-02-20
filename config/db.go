package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DBConfigs map[string]*DBConfig
type DBConfig struct {
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

func LoadDBConfig(path string) (*DBConfigs, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfgs DBConfigs
	if err := yaml.Unmarshal(buf, &cfgs); err != nil {
		return nil, err
	}
	return &cfgs, nil
}
