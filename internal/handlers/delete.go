package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleDeletePost godoc
//
//	@Summary		Delete post
//	@Description	Delete post by id
//	@ID				delete-post
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Post ID"
//	@Success		200	{string}	string	"ok"
//	@Failure		400	{string}	string	"bad request"
//	@Failure		500	{string}	string	"internal error"
//	@Router			/posts/{id} [delete]
func (p *PostsHandler) HandleDeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.l.Print(err)
		c.Error(ErrBadRequest)
		return
	}
	err = p.rep.DeletePost(c.Request.Context(), id)
	if err != nil {
		p.l.Print(err)
		c.Error(ErrInternalError)
		return
	}
	c.Status(http.StatusOK)
}
