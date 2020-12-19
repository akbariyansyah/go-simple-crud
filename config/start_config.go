package config

import (
	"go-simple-crud/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Config struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func StartServer(router *mux.Router) {
	
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = utils.ReadConfig("server.port")
	}
	log.Printf("Server is starting at port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
func StartConfig() *Config {
	conf := &Config{}

	conf.DB_HOST = os.Getenv("DB_HOST")
	if os.Getenv("DB_HOST") == "" {
		conf.DB_HOST = utils.ReadConfig("database.host")
	}
	conf.DB_PORT = os.Getenv("DB_PORT")
	if os.Getenv("DB_HOST") == "" {
		conf.DB_PORT = utils.ReadConfig("database.port")
	}
	conf.DB_USER = os.Getenv("DB_USER")
	if os.Getenv("DB_HOST") == "" {
		conf.DB_USER = utils.ReadConfig("database.user")
	}
	conf.DB_PASS = os.Getenv("DB_PASS")
	if os.Getenv("DB_HOST") == "" {
		conf.DB_PASS = utils.ReadConfig("database.pass")
	}
	conf.DB_NAME = os.Getenv("DB_NAME")
	if os.Getenv("DB_NAME") == "" {
		conf.DB_NAME = utils.ReadConfig("database.name")
	}
	return conf
}
