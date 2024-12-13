package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"path/filepath"
	"qwen/internal"
	"qwen/internal/config"
	"syscall"
)

func main() {
	config.InitConfig()
	//config.InitMysql()
	config.InitSqlite()
	config.InitOpenaiClient()

	r := gin.Default()

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	r.LoadHTMLGlob("view/templates/**/*.html")

	internal.InitRooter(r)

	go func() {
		r.Run(":3003")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("go 程序已经关闭")

}
