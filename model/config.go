package model

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Username string `koanf:"user"`
	Password string `koanf:"password"`
	DBName   string `koanf:"db_name"`
	SSLMode  string `koanf:"ssl_mode"`
}
