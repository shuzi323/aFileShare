package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

//页面title图标
func favicon(c *gin.Context) {
	c.Data(http.StatusOK, "image/x-icon", icon)
}

//首页
func index(c *gin.Context) {
	dirs, files := readFilesName()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"dirs":  dirs,
		"files": files,
	})
}

//下载专用
func download(c *gin.Context) {
	filename := c.Param("filename")
	haveTwoDot, _ := regexp.MatchString(`\.\.`, filename)
	if haveTwoDot {
		c.AbortWithStatus(404)
		return
	}
	file, err := ioutil.ReadFile("./share/" + filename)
	//这里需要正则匹配..防止访问上一级目录
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.Header("Content-Disposition", "attachment")
	c.Header("filename", filename)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Data(http.StatusOK, "application/octet-stream", file)
}

//预览页面
func preView(c *gin.Context) {
	filename := c.Param("filename")
	haveTwoDot, _ := regexp.MatchString(`\.\.`, filename)
	if haveTwoDot {
		c.AbortWithStatus(404)
		return
	}
	file, err := os.Open("./share/" + filename)
	defer file.Close()
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	var size string
	size = fileSize(file)
	c.HTML(http.StatusOK, "preview.html", gin.H{
		"filename": filename,
		"size":     size,
	})
}

func viewData(c *gin.Context) {
	filename := c.Param("filename")
	haveTwoDot, _ := regexp.MatchString(`\.\.`, filename)
	if haveTwoDot {
		c.AbortWithStatus(404)
		return
	}
	file, err := ioutil.ReadFile("./share/" + filename)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}
	mp4, _ := regexp.MatchString(`\.mp4$`, filename)
	if mp4 {
		c.Data(http.StatusOK, "video/mpeg4", file)
	}
}

//文件图标
func myIcon(router *gin.Engine) {
	var icons []string
	items, err := ioutil.ReadDir("./static/icon")
	if err != nil {
		fmt.Println("read dir error")
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		} else {
			icons = append(icons, "static/icon/"+item.Name())
		}
	}
	for _, i := range icons {
		i_icon, err := ioutil.ReadFile(i)
		if err != nil {
			fmt.Println("Load icons error!", err)
		}
		router.GET("/"+i, func(c *gin.Context) {
			c.Data(http.StatusOK, "image/png", i_icon)
		})
	}
}
