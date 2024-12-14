package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"os/signal"
	"qwen/internal"
	"qwen/internal/config"
	"syscall"
)

//go:embed view/templates/**/*.html
var content embed.FS

func main() {
	config.InitConfig()
	//config.InitMysql()
	config.InitSqlite()
	config.InitOpenaiClient()

	r := gin.Default()

	//r.LoadHTMLGlob("view/templates/**/*.html")
	// 加载嵌入的 HTML 文件
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(content, "view/templates/**/*.html")))

	internal.InitRooter(r)

	go func() {
		r.Run(":3003")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("go 程序已经关闭")

}
