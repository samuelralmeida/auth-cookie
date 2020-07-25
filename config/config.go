package config

import (
	"github.com/spf13/viper"
)

var JWTKey []byte

func init() {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	JWTKey = []byte(viper.GetString("jwtkey"))
}
