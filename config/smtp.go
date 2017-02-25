package config

type SMTP struct {
	Host string `yaml:"host"`
	Port int  `yaml:"port"`
}

func GetSMTP() *SMTP {
	return cfg.SMTP
}
