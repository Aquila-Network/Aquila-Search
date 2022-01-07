package main

import (
	"aquiladb/src/config"
	"aquiladb/src/controller"
	"aquiladb/src/middleware"
	"aquiladb/src/repository"
	"aquiladb/src/service"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	envConfig := config.InitEnvConfig()

	dbPostgre, err := config.NewPostgresDB(envConfig)
	if err != nil {
		fmt.Println("Error connect to db.")
		log.Fatal(err)
	}

	repos := repository.NewRepository(dbPostgre)
	services := service.NewService(repos)
	controllers := controller.NewController(services)

	SetupLogOutput()
	server := gin.Default()
	// server.Use(gin.Recovery(), gin.Logger())

	server.GET("/", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(400, gin.H{"errorResponsemessage": "kgjhggkg"})
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Aquila DB new technologies.",
		})
	})

	auth := server.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := server.Group("/api", middleware.UserIdentity)
	{
		api.GET("/secret", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Protected page.",
			})
		})
	}

	admin := server.Group("/admin", middleware.UserIdentity, middleware.AdminIdentity)
	{
		admin.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello admin.",
			})
		})
	}

	customer := server.Group("/customer")
	{
		customer.POST("", controllers.CreateTempCustomer)
	}

	server.Run(":" + envConfig.App.Port)

}

func SetupLogOutput() {
	err := os.MkdirAll("logs", 8644)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
