package config

// import "os"

type Config struct {
	DatabaseHost string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	DatabaseName string
}

func StartConfig() *Config {
	// conf := &Config{
	// 	DatabaseHost: os.Getenv("DB_HOST"),
	// 	DatabasePort: os.Getenv("DB_PORT"),
	// 	DatabaseUser: os.Getenv("DB_USER"),
	// 	DatabasePass: os.Getenv("DB_PASS"),
	// 	DatabaseName: os.Getenv("DB_NAME"),
	// }
	conf := &Config{
		DatabaseHost: "localhost",
		DatabasePort: "3306",
		DatabaseUser: "root",
		DatabasePass: "password123",
		DatabaseName: "db_product",
	}
	return conf
}
