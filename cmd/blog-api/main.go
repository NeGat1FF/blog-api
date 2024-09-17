package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NeGat1FF/blog-api/cmd/docs"
	"github.com/NeGat1FF/blog-api/config"
	"github.com/NeGat1FF/blog-api/internal/handlers"
	"github.com/NeGat1FF/blog-api/internal/middleware"
	"github.com/NeGat1FF/blog-api/internal/repository"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Blog API
//	@version		1.0
//	@description	This is a simple blog API

//	@host		localhost:8080
//	@BasePath	/

// @schemes	http
func main() {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		return
	}

	db := repository.InitDB(cfg)
	defer db.Close()

	rep := repository.NewPostRepository(db)

	h := handlers.NewPostHandler(log.New(os.Stderr, "[Error] ", log.Ldate|log.Ltime|log.Lshortfile), rep)

	r := gin.Default()

	r.Use(middleware.ErrorMiddleware())

	docs.SwaggerInfo.Host = "localhost:8080"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/posts", middleware.AddPostMiddleware(), h.HandleAddPost)

	r.PUT("/posts/:id", middleware.UpdatePostMiddleware(), h.HandleUpdatePost)

	r.DELETE("/posts/:id", h.HandleDeletePost)

	r.GET("/posts/:id", h.HandleGetPostByID)
	r.GET("/posts", h.HandleGetAllPosts)

	r.Run(":8080")
}
