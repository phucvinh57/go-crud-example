package configs

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

type Config struct {
	PgUser     string `mapstructure:"POSTGRES_USER"`
	PgPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PgDB       string `mapstructure:"POSTGRES_DB"`
	PgHost     string `mapstructure:"POSTGRES_HOST"`
	PgPort     int    `mapstructure:"POSTGRES_PORT"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		once.Do(func() {
			viper.SetConfigFile(".env")
			viper.AutomaticEnv()

			err := viper.ReadInConfig()
			if err != nil {
				panic(err)
			}

			err = viper.Unmarshal(&config)
			if err != nil {
				panic(err)
			}
		})
	}

	return config
}

