package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Env struct {
	Port                    int    `mapstructure:"PORT"`
	AppEnv                  string `mapstructure:"APP_ENV"`
	AppURL                  string `mapstructure:"APP_URL"`
	DatabaseURL             string `mapstructure:"DATABASE_URL"`
	ClientURL               string `mapstructure:"CLIENT_URL"`
	JWTSignKey              string `mapstructure:"JWT_SIGN_KEY"`
	GoogleOAuthClientID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleOAuthClientSecret string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOAuthRedirectURL  string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URL"`
}

func NewEnv(envPath string) (*Env, error) {
	env := Env{}

	if os.Getenv("APP_ENV") == "production" {
		viper.AutomaticEnv()
	} else {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")

		if envPath == "" {
			envPath = "."
		}

		viper.AddConfigPath(envPath)

		err := viper.ReadInConfig()

		if err != nil {
			return nil, err
		}
	}

	err := viper.Unmarshal(&env)

	if err != nil {
		return nil, err
	}

	log.Println(env)

	return &env, nil
}
