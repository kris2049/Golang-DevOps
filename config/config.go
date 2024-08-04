package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DBType     string `mapstructure:"db_type"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
}

var (
	config *Config
	once   sync.Once
)

func InitConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath("config")

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("config info: ", viper.Get("db_type"))
		config = &Config{}
		err = viper.Unmarshal(config)
		if err != nil {
			fmt.Printf("Unable to decode into struct, %v", err)
		}
	})
	return config
}
