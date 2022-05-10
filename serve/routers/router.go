package routers

import (
	"esp_webrtc/pkg/setting"
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
		//
		api_user.POST("/register", register.Register)
		//新建标签

	}

	return r
}
