package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	AuthUrl    string `mapstructure:"AUTH_SERVER_URL"`
	MonitUrl   string `mapstructure:"MONITIZATION_SERVER_URL"`
	Email      string `mapstructure:"EMAIL"`
	AppPass    string `mapstructure:"EMAIL_APP_PASS"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return

}
