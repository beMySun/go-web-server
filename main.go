package main

import (
	"go-web-server/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupRouter(engine *gin.Engine, reg *handler.Registry) {
	engine.GET("/ping", reg.PingHandler)
	engine.GET("/reportData", reg.ListReportData)
	engine.POST("/reportData", reg.CreateReportData)
	engine.GET("/getTodoList", reg.ListTodoList)
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	reg := handler.NewRegistry()
	setupRouter(r, reg)
	// listen and serve on 0.0.0.0:8080
	r.Run()
	reg.DAO.Close()
}
