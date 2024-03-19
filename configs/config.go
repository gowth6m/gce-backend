package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host string `envconfig:"APP_HOST" default:"localhost"`
		PORT string `envconfig:"APP_PORT" default:":8080"`
	}
	Database struct {
		MongoUri string `envconfig:"MONGODB_URI" default:"mongodb://localhost:27017"`
	}
	Auth struct {
		JWTSecret        string `envconfig:"JWT_SECRET" default:"token-secret"`
		TokenExpire      int    `envconfig:"TOKEN_EXPIRE" default:"60"`
		ShortTokenExpire int    `envconfig:"SHORT_TOKEN_EXPIRE" default:"15"`
	}
}

var appConfig = &Config{}

func AppConfig() *Config {
	return appConfig
}

func LoadConfig() error {
	godotenv.Load()
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	return nil
}
