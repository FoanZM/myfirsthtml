package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type GinStruct struct {
	egine *gin.Engine
	sync.RWMutex
}

// 主页请求
func (g *GinStruct) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Etcd Config Manager"})
}

// 注册接口回调
func (g *GinStruct) RegistHandler() {
	g.egine.LoadHTMLGlob("./web/html/*.html")
	g.egine.Static("/web/js", "./web/js")
	g.egine.Static("/web/html", "./web/html")

	g.egine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Manager",
		})
	})
}

// 初始化gin
func (g *GinStruct) InitAll() {
	g.egine = gin.Default()
	// regist router
	g.egine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong\n",
		})
	})
}

// 开始Listen
func (g *GinStruct) StartRun() {
	g.egine.Run(":8080")
}

func main() {
	newEgine := new(GinStruct)
	newEgine.InitAll()
	newEgine.RegistHandler()

	newEgine.StartRun()
}
