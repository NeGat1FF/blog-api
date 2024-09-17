package handlers

import (
	"net/http"

	"github.com/NeGat1FF/blog-api/internal/models"
	"github.com/gin-gonic/gin"
)

type AddPostRequest struct {
	Title    string   `json:"title"`
	Tags     []string `json:"tags"`
	Category string   `json:"category"`
	Content  string   `json:"content"`
} //	@name	AddPostRequest

type AddPostResponse struct {
	ID int `json:"id"`
}

// HandleAddPost godoc
//
//	@Summary		Add post
//	@Description	Add post
//	@ID				add-post
//	@Accept			json
//	@Produce		json
//	@Param			post	body		AddPostRequest	true	"Post"
//	@Success		200		{object}	AddPostResponse	"ok"
//	@Failure		400		{string}	string			"bad request"
//	@Failure		500		{string}	string			"internal error"
//	@Router			/posts [post]
func (p *PostsHandler) HandleAddPost(c *gin.Context) {
	newPost := c.MustGet("newPost").(models.Post)

	id, err := p.rep.AddPost(c.Request.Context(), newPost)
	if err != nil {
		p.l.Print(err)
		c.Error(ErrInternalError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
