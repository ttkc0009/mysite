package main

import (
	"./routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// template files
	router.LoadHTMLGlob("templates/*.html")

	// static files
	router.Static("/assets", "./assets")

	// path handler
	router.GET("/home", routes.GetHome)
	router.GET("/hello/:name", routes.HelloParam)

	// グルーピング
	user := router.Group("/api")
	{
		user.GET("/hello", routes.HelloJson)
		user.GET("/hello/:name", routes.HelloJsonPram)
	}

	router.NoRoute(routes.Get404) // どのルーティングにも当てはまらなかった場合に処理
	router.Run(":8080")
}
