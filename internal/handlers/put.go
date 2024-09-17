package handlers

import (
	"net/http"
	"strconv"

	"github.com/NeGat1FF/blog-api/internal/models"
	"github.com/gin-gonic/gin"
)

// Update body request can have only one field or all fields
type UpdatePostRequest struct {
	Title    *string   `json:"title"`
	Tags     *[]string `json:"tags"`
	Category *string   `json:"category"`
	Content  *string   `json:"content"`
} //	@name	UpdatePostRequest

// HandleUpdatePost godoc
//
//	@Summary		Update post
//	@Description	Update post by id
//	@ID				update-post
//	@Accept			json
//	@Param			id		path		int					true	"Post ID"
//	@Param			post	body		UpdatePostRequest	true	"Post"
//	@Success		200		{string}	string				"ok"
//	@Failure		400		{string}	string				"bad request"
//	@Failure		500		{string}	string				"internal error"
//	@Router			/posts/{id} [put]
func (p *PostsHandler) HandleUpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.l.Print(err)
		c.Error(ErrBadRequest)
		return
	}

	updatePost := c.MustGet("updatePost").(models.Post)

	updatePost.ID = id
	err = p.rep.UpdatePost(c.Request.Context(), updatePost)
	if err != nil {
		p.l.Print(err)
		c.Error(ErrInternalError)
		return
	}

	c.Status(http.StatusOK)
}
