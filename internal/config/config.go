package config

type Config struct {
	Debug         bool
	ListenAddress string
	SqlitePath    string
	PostgresDSN   string
}

var DefaultConfig = Config{
	Debug:         true,
	ListenAddress: ":5050",
	PostgresDSN:   "host=localhost user=admin password=admin dbname=lexilift port=5432 sslmode=disable",
	//SqlitePath:    "./gorm.db",
}
