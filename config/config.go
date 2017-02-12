package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	CertFilePath   string `yaml:"cert_file_path"`
	KeyFilePath string `yaml:"key_file_path"`
	HTTPPort string `yaml:"http_port"`
}

func Load(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(buf, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
