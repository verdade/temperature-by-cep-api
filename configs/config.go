package configs

import "github.com/spf13/viper"

type Conf struct {
	ViaCepApiUrl  string `mapstructure:"VIACEP_API_URL"`
	WeatherApiUrl string `mapstructure:"WEATHER_API_URL"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, err
}
