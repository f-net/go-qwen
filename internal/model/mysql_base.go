package model

import (
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"qwen/internal/types"
)

var db *gorm.DB
var configFileName = "config.yaml"

func InitMysql() {
	config := types.Config{}
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

	var db *sql.DB
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port))
	if err != nil {
		return
	}
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
