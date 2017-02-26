package config

// SMTP smtp接続情報を保持
type SMTP struct {
	Host string `yaml:"host"`
	Port int  `yaml:"port"`
}

// GetSMTP SMTP接続情報を返却
func GetSMTP() *SMTP {
	return cfg.SMTP
}
