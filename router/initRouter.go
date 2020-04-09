package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/user")
	{
		userGroup.GET("/index", func(context *gin.Context) {
			//context.Get()
			//context.Set()
			context.Abort()
		})
	}

	// 添加 Get 请求路路由
	router.GET("/", func(context *gin.Context) {
		query, _ := context.GetQuery("username")
		fmt.Println(query)
		context.String(http.StatusOK, "hello gin")
	})

	router.GET("/:user/:name", func(context *gin.Context) {
		user := context.Param("user")
		name := context.Param("name")
		context.JSON(http.StatusOK, map[string]interface{}{
			"username": user,
			"name":     name,
		})
	})
	//router.Any()
	router.NoRoute(handler404Func)
	router.NoMethod(func(context *gin.Context) {
		context.String(http.StatusOK, "NoMethod")
	})
	return router
}

func handler404Func(context *gin.Context) {
	context.String(http.StatusOK, "not font")
}
