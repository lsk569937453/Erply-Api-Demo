package main

import (
	"erply-api/config"
	_ "erply-api/dao"
	"erply-api/log"
	"erply-api/router"
	_ "erply-api/task"
	"erply-api/util"
	"runtime"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"erply-api/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/fatih/color"
)

const PORT = "9394"

func init() {
	r := router.InitController()
	host, err := config.GetValue("swagger", "host")
	if err != nil {
		panic("Can not get the host of swagger!")
	}

	docs.SwaggerInfo.Host = host + ":" + PORT

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	printBanner()
	err = r.Run(":" + PORT)
	if err != nil {
		log.Error("initController error", err.Error())
	}
}

// @title Erply Api Demo
// @version 1.0
// @description This is a Erply api demo.
// @termsOfService http://swagger.io/terms/

// @contact.name Shikai Liu
// @contact.url https://github.com/lsk569937453
// @contact.email lsk569937453@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
}

func printBanner() error {
	localIp, err := util.GetIp()
	if err != nil {
		return err
	}
	color.Yellow(*localIp + ":" + PORT + "/")

	return nil
}
