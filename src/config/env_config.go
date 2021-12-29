package config

import (
	"fmt"
	"log"
	"os"

	"github.com/golobby/dotenv"
)

type ConfigEnv struct {
	Debug bool `env:"DEBUG"`
	App   struct {
		Name string `env:"APP_NAME"`
		Port string `env:"APP_PORT"`
	}
	Database struct {
		Host    string `env:"DB_HOST"`
		Port    string `env:"DB_PORT"`
		User    string `env:"DB_USER"`
		Pass    string `env:"DB_PASS"`
		Name    string `env:"DB_NAME"`
		SslMode string `env:"DB_SSLMODE"`
	}
	Hash struct {
		Salt    string `env:"HASH_SALT"`
		SignKey string `env:"HASH_SIGN_KEY"`
	}
}

func InitEnvConfig() *ConfigEnv {

	c := &ConfigEnv{}
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println("Failed to load .env file")
		log.Fatalln(err)
	}
	err = dotenv.NewDecoder(file).Decode(c)
	if err != nil {
		fmt.Println("Failed to read .env file")
		log.Fatalln(err)
	}

	return c
}
