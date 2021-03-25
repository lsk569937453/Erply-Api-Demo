package router

import (
	"erply-api/controller"
	"erply-api/log"
	"erply-api/midware"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitController() *gin.Engine {
	gin.DefaultWriter = log.BaseGinLog()
	r := gin.New()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: int(20 * 60), Path: "/"})
	r.Use(sessions.Sessions("mysession", store))
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %v |%3d| %13v | %15s | %-7s  %#v %s |\"%s\" \n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
			param.Request.UserAgent(),
		)
	}))

	r.GET("/getToken", controller.GetToken)

	v1 := r.Group("/api")
	{
		v1.Use(midware.JWTAuthMiddleware())

		v1.POST("/AddCustomerRewardPoints", controller.AddCustomerRewardPoints)
		v1.POST("/AddCustomerRewardPointsBulk", controller.AddCustomerRewardPointsBulk)
		v1.POST("/GetCustomersBulk", controller.GetCustomersBulk)
		v1.GET("/GetCustomerByCustomerId", controller.GetCustomerByCustomerId)
	}

	return r
}
