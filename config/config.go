package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig(){
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./config")
	if err :=viper.ReadInConfig();err != nil {
		panic("read config is error")
	}else{
		fmt.Println("read config is success")
		viper.WatchConfig()
	}
}