package config

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
)

var db *gorm.DB

func InitMysql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.DB,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(err)
	}
}
func InitSqlite() {
	var err error
	path, _ := filepath.Abs(os.Args[0])
	// 提取目录部分
	execDir := filepath.Dir(path) + "/" + config.Mysql.Dir
	// 如果子目录不存在，则创建它
	err = os.MkdirAll(execDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	//dsn := "file:" + execDir + "\\" + config.Mysql.DB + ".db" + "?cache=shared&_foreign_keys=on"
	dsn := execDir + "/" + config.Mysql.DB + ".db"

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err = gorm.Open(sqlite.Open(dsn), gormConfig)
	if err != nil {
		panic(err)
	}

	//db, err := gorm.Open(sqlite.(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//})
	if err != nil {
		panic(err)
	}
	sqlPath := execDir + "/docs"
	files, err := os.ReadDir(sqlPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		// 获取文件后缀
		if ext := filepath.Ext(file.Name()); ext == ".sql" {
			// 获取文件内容
			readFile, err := os.ReadFile(sqlPath + "/" + file.Name())
			if err != nil {
				panic(err)
			}
			err = db.Debug().Exec(string(readFile)).Error
			fmt.Println(err)
		}
	}

}

func GetDB() *gorm.DB {
	return db
}
