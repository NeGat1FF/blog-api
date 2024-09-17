package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NeGat1FF/blog-api/internal/handlers"
	"github.com/NeGat1FF/blog-api/internal/models"
	"github.com/gin-gonic/gin"
)

var ErrUnprocessableEntity = errors.New("unprocessable entity")

func AddPostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPost models.Post
		err := c.BindJSON(&newPost)
		if err != nil {
			fmt.Println("Failed to bind json", err)
			c.Error(ErrUnprocessableEntity)
			return
		}

		if newPost.Title == "" || newPost.Content == "" || newPost.Category == "" || len(newPost.Tags) < 1 {
			fmt.Println("Empty fields")
			c.Error(ErrUnprocessableEntity)
			return
		}

		c.Set("newPost", newPost)
		c.Next()
	}
}

func UpdatePostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatePost models.Post
		err := c.BindJSON(&updatePost)
		if err != nil {
			fmt.Println("Failed to bind json")
			c.Error(ErrUnprocessableEntity)
			return
		}

		if updatePost.Title == "" && updatePost.Content == "" && updatePost.Category == "" && len(updatePost.Tags) < 1 {
			fmt.Println("Empty fields")
			c.Error(ErrUnprocessableEntity)
			return
		}

		c.Set("updatePost", updatePost)
		c.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			switch c.Errors[0].Err {
			case handlers.ErrInternalError:
				c.JSON(http.StatusInternalServerError, c.Errors[0].Err)
			case handlers.ErrBadRequest:
				c.JSON(http.StatusBadRequest, c.Errors[0].Err)
			case ErrUnprocessableEntity:
				c.JSON(http.StatusUnprocessableEntity, c.Errors[0].Err)
			default:
				c.JSON(http.StatusInternalServerError, c.Errors[0].Err)

			}
		}
	}
}
