package config

// Config represents the config struct
type Config struct {
	Secrets     []string
	Debug       bool
	VaultTool   string   `mapstructure:"vault_tool"`
	EncryptArgs []string `mapstructure:"encrypt_args"`
	DecryptArgs []string `mapstructure:"decrypt_args"`
}

// NewConfig Returns a New Config
func NewConfig() *Config {
	return &Config{}
}
