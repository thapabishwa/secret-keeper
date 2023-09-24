package config

// Config represents the config struct
type Config struct {
	FilePatterns []string `mapstructure:"secret_files_patterns"`
	Debug        bool     `mapstructure:"debug"`
	VaultTool    string   `mapstructure:"vault_tool"`
	EncryptArgs  []string `mapstructure:"encrypt_args"`
	DecryptArgs  []string `mapstructure:"decrypt_args"`
	ViewArgs     []string `mapstructure:"view_args"`
}

// NewConfig Returns a New Config
func NewConfig() *Config {
	return &Config{}
}
