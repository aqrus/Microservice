package utils

import "github.com/spf13/viper"

type Config struct {
	PostgresUser string `mapstructure:"POSTGRES_USER"`;
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"DB_NAME"`;
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error ) {
	viper.AddConfigPath(path);
	viper.SetConfigName("app");
	viper.SetConfigType("env");

	viper.AutomaticEnv()

	err = viper.ReadInConfig();

	if err != nil {
		return 
	}

	err = viper.Unmarshal(&config)
	return 
}