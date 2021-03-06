package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func Hello(c *gin.Context) {
	name := c.DefaultQuery("name", "ONAMAE") // HOGEはデフォルト値?
	//name := c.Query("lastname") // デフォルトがない場合
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"name": name,
	})
}

func HelloParam(c *gin.Context) {
	name := c.Param("name")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"name": name,
	})
}

func Get404(c *gin.Context) {
	// helloに飛ばす
	c.HTML(http.StatusOK, "404.html", gin.H{})
}
