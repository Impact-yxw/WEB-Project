package main

import (
	"OSS/apiServer/conf"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/objects"
	"OSS/apiServer/temp"
	"OSS/apiServer/tmpl"
	"OSS/apiServer/versions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	confile := "conf/conf.json"
	conf.Conf.Parse(confile)
}

func main() {

	var url string //监听地址:端口
	url = conf.Conf.ListenAddr + ":" + conf.Conf.ListenPort
	log.Println(url)
	engine := gin.Default()

	//启动心跳 返回dataservers中随机一个地址
	go heartbeat.ListenHeartbeat()
	//前端界面
	engine.LoadHTMLGlob("html/index.html")
	engine.GET("/OSS", indexPage)
	engine.GET("/OSS/file/:uid", tmpl.Get)

	//断点续传服务
	engine.PUT("/OSS/temp/*file", temp.Put)
	engine.HEAD("/OSS/temp/*file", temp.Head)

	//OSS存储服务
	engine.PUT("/OSS/objects/:file", objects.Put)
	engine.POST("/OSS/objects/:file", objects.Post)
	engine.GET("/OSS/objects/:file", objects.Get)
	engine.DELETE("/OSS/objects/:file", objects.Delete)

	//版本控制服务
	engine.GET("/OSS/versions/:file", versions.Get)

	//文件定位服务
	engine.GET("/OSS/locate/:file", locate.Get)

	engine.Run(url)
}

func indexPage(c *gin.Context) {
	b, _ := ioutil.ReadFile("../index.html")
	c.Data(http.StatusOK, "text/html", b)
}
