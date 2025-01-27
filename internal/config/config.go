package config

type Config struct {
	Debug         bool
	ListenAddress string
	DatabasePath  string
}

var DefaultConfig = Config{
	Debug:         true,
	ListenAddress: ":5050",
	DatabasePath:  "./gorm.db",
}
