package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvDBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewEnvDBConfig() *EnvDBConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return &EnvDBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_DATABASE"),
	}
}

func (c *EnvDBConfig) GetHost() string {
	return c.host
}

func (c *EnvDBConfig) GetPort() string {
	return c.port
}

func (c *EnvDBConfig) GetUsername() string {
	return c.username
}

func (c *EnvDBConfig) GetPassword() string {
	return c.password
}

func (c *EnvDBConfig) GetDatabase() string {
	return c.database
}
