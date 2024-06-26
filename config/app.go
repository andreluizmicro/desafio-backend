package config

import "github.com/spf13/viper"

var cfg *AppConfig

type AppConfig struct {
	DBDriver                      string `mapstructure:"DB_DRIVER"`
	DBHost                        string `mapstructure:"DB_HOST"`
	DBPort                        string `mapstructure:"DB_PORT"`
	DBUser                        string `mapstructure:"DB_USER"`
	DBPassword                    string `mapstructure:"DB_PASSWORD"`
	DBName                        string `mapstructure:"DB_NAME"`
	WebServerPort                 string `mapstructure:"WEB_SERVER_PORT"`
	AuthorizationClientApiUrl     string `mapstructure:"AUTHORIZATION_CLIENT_API_URL"`
	AuthorizationClientApiVersion string `mapstructure:"AUTHORIZATION_CLIENT_API_VERSION"`
	AuthorizationClientApiToken   string `mapstructure:"AUTHORIZATION_CLIENT_API_TOKEN"`
	NotificationClientApiUrl      string `mapstructure:"NOTIFICATION_CLIENT_API_URL"`
	NotificationClientApiVersion  string `mapstructure:"NOTIFICATION_CLIENT_API_VERSION"`
	NotificationClientApiToken    string `mapstructure:"NOTIFICATION_CLIENT_API_TOKEN"`
}

func LoadConfig(path string) (*AppConfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}

func GetAuthorizationConfigClient() *AppConfig {
	return cfg
}
