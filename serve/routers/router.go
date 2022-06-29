package routers

import (
	"esp_webrtc/pkg/setting"
	"esp_webrtc/routers/article"
	register "esp_webrtc/routers/user/register"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	api_user := r.Group("/user")
	{
		//注册
		api_user.POST("/register", register.Register)
		//新建标签
		api_user.POST("/login", register.Login)
		//获取用户信息
		api_user.GET("/getUserEmailById/:id", register.GetUserEmailById)
		//获取文章列表
	}
	api_article := r.Group("/article")
	{
		api_article.GET("/getAllArticleList", article.GetAllArticleList)
	}
	return r
}
