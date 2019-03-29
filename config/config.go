package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Host     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "postgres",
			Password: "postgres",
			Name:     "came",
			Host:     "localhost",
			Charset:  "utf8",
		},
	}
}
