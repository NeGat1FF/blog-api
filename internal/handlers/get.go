package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleGet
//	@Summary		Get all posts
//	@Description	Get all posts
//	@ID				get-all-posts
//	@Accept			json
//	@Produce		json
//	@Param			term	query		string		false	"Search term"
//	@Success		200		{array}		models.Post	"ok"
//	@Failure		500		{string}	string		"internal error"
//	@Router			/posts [get]
func (p *PostsHandler) HandleGetAllPosts(c *gin.Context) {
	val, exists := c.GetQuery("term")
	if !exists {
		posts, err := p.rep.GetAllPosts(c.Request.Context())
		if err != nil {
			p.l.Print(err)
			c.Error(ErrInternalError)
			return
		}
		c.JSON(http.StatusOK, posts)
	} else {
		posts, err := p.rep.GetPostsByTerm(c.Request.Context(), val)
		if err != nil {
			p.l.Print(err)
			c.Error(ErrInternalError)
			return
		}
		c.JSON(http.StatusOK, posts)
	}
}

// HandleGetPostByID
//	@Summary		Get post by id
//	@Description	Get post by id
//	@ID				get-post-by-id
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int			true	"Post ID"
//	@Success		200	{object}	models.Post	"ok"
//	@Failure		400	{string}	string		"bad request"
//	@Failure		500	{string}	string		"internal error"
//	@Router			/posts/{id} [get]
func (p *PostsHandler) HandleGetPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.l.Print(err)
		c.Error(ErrBadRequest)
		return
	}
	post, err := p.rep.GetPostByID(c.Request.Context(), id)
	if err != nil {
		p.l.Print(err)
		c.Error(ErrInternalError)
		return
	}
	c.JSON(http.StatusOK, post)
}
