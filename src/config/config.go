package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type SupabaseDatabaseConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"sslmode"`
	SSLRoot  string `mapstructure:"sslroot"`
	Schema   string `mapstructure:"schema"`
}

type SupabaseClientConfig struct {
	URL string `mapstructure:"supabase_url"`
	Key string `mapstructure:"supabase_key"`
}

type Config struct {
	SupabaseDatabaseConfig SupabaseDatabaseConfig `mapstructure:"supabase_db"`
	SupabaseClientConfig   SupabaseClientConfig   `mapstructure:"supabase_client"`
}

func New() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file: ", err)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Error unmarshalling config: ", err)
		return nil, err
	}
	return &config, nil
}
