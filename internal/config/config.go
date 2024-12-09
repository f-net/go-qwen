package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"qwen/internal/types"
)

var db *gorm.DB
var config = types.Config{}
var configFileName = "config.yaml"

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(configFileName)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configIn file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}
}
