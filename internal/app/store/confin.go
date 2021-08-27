package store

type Config struct {
	DataBaseURL string `yaml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}