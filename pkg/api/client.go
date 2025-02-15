package api

import "github.com/spf13/viper"

type Client struct {
	Token string
	Host  string
}

func NewClient() Client {
	client := Client{
		viper.GetString("token"),
		viper.GetString("host"),
	}
	return client
}
