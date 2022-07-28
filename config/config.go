package config

import "github.com/spf13/viper"

var Title string
var Port string

func Conf() {
	var Vi = viper.New()
	Vi.SetConfigFile("test.yaml")
	Vi.ReadInConfig()
	Title = Vi.GetString("title")
	Port = Vi.GetString("port")
}

