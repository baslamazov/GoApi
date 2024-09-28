package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env              string `yaml:"environment" env:"ENV" env-required:"true"`
	ConnectionString `yaml:"connection_string" env-required:"true"`
	HTTPServer       `yaml:"http_server" env:"HTTP_SERVER" env-required:"true"`
}
type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}
type ConnectionString struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	DBName   string `yaml:"db_name" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-required:"true"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("нет пути")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Не существует конфига %s", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Не читается конфиг %s", err)
	}
	return &config
}
