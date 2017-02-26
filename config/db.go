package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// DBConfigs 複数のDB設定を保持
type DBConfigs map[string]*DBConfig

// DBConfig DB接続情報を保持
type DBConfig struct {
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

// LoadDBConfig yaml pathを元にDB接続情報を読み込む
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
