package main

import (
	"aquiladb/src/config"
	"aquiladb/src/controller"
	"aquiladb/src/middleware"
	moduledb "aquiladb/src/module_db"
	"aquiladb/src/repository"
	"aquiladb/src/service"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	_ "aquiladb/swagger_docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Aquila DB
// @version 1.0
// @description Aquila DB

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

	// test route for debugging db module
	server.GET("/test", func(ctx *gin.Context) {
		// create aquila db
		moduledb.CreateAquilaDatabase()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "See the result in the console.",
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
		customer.PATCH("", middleware.UserIdentity, controllers.CreatePermanentCustomer)
		customer.GET("", middleware.UserIdentity, controllers.GetCustomer)
		customer.POST("/auth", controllers.Auth)
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":" + envConfig.App.Port)

}

func SetupLogOutput() {
	err := os.MkdirAll("logs", 8644)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("logs/%v_gin.log", time.Now().Format("01_02_2006"))
	f, _ := os.Create(fileName)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
