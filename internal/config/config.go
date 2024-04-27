package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Telegram telegram
	Postgres postgres
	Rabbit   rabbitMQ
}

type telegram struct {
	TelegramToken string `env:"TELEGRAM_TOKEN"`
	WebhookURL    string `env:"WEBHOOK_URL"`
}

type postgres struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	Username string `env:"POSTGRES_USERNAME" env-default:"user"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"password"`
	Database string `env:"POSTGRES_DATABASE" env-default:"blazingMoney"`
}

type rabbitMQ struct {
	Host     string `env:"RABBIT_HOST" env-default:"localhost"`
	Port     string `env:"RABBIT_PORT" env-default:"15672"`
	Username string `env:"RABBIT_USERNAME" env-default:"rmuser"`
	Password string `env:"RABBIT_PASSWORD" env-default:"rmpassword"`
}

// Load config from env
func NewConfig() *Config {
	_ = godotenv.Load()

	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Database,
	)
}
