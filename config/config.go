package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Driver:   "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Name:     "todo_app",
		},
	}
}
