package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var icon []byte

func main() {
	var err error
	icon, err = ioutil.ReadFile("static/img/favicon.ico")
	if err != nil {
		log.Println("Open icon file error!", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	initialRouter(router)

	router.Run("0.0.0.0:8080")
}

func initialRouter(router *gin.Engine) {
	router.GET("/static/img/favicon.ico", favicon)
	router.StaticFS("/static/css/", http.Dir("static/css"))
	myIcon(router)

	router.GET("/", index)
	router.GET("/download/:filename", download)
	router.GET("/preView/:filename", preView)
	router.GET("/viewData/:filename", viewData) //预览页面加载的内容
}
