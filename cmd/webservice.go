package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/maxzycon/SIB-Golang-Assigment-3/config"
)

func Init() {
	// ----- init mariadb
	db := config.InitMariaDb()
	// ----- ini handler, service
	serviceInit := Service{
		db: db,
	}
	controllerInit := Controller{
		service: serviceInit,
	}

	r := gin.Default()
	r.GET("/auto_reload", controllerInit.AutoReloadData)

	r.Run(":8005") // listen and serve on 0.0.0.0:8085
}
