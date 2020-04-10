package routers

import (
	"gin_blog/controllers"
	"github.com/gin-contrib/sessions" // session包 定义了一套session操作的接口 类似于 database/sql
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions/cookie"  // session具体存储的介质
	"github.com/gin-contrib/sessions/redis" // session具体存储的介质
	//"github.com/gin-contrib/sessions/memcached"  // session具体存储的介质
	// github.com/go-redis/redis  --> go连接redis的一个第三方库
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("views/*")

	// 设置session midddleware
	//store := cookie.NewStore([]byte("secret"))
	store, err := redis.NewStore(10, "tcp", "54.180.98.165:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("mysession", store))

	//注册：
	r.GET("/register", controllers.RegisterGet)
	r.POST("/register", controllers.RegisterPost)

	r.GET("/login", controllers.LoginGet)
	r.POST("/login", controllers.LoginPost)

	r.GET("/", controllers.HomeGet)

	//路由组
	article := r.Group("/article")
	{
		article.GET("/add", controllers.AddArticleGet)
		article.POST("/add", controllers.AddArticlePost)
	}
	return r

}
