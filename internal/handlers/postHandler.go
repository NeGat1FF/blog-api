package handlers

import (
	"errors"
	"log"

	"github.com/NeGat1FF/blog-api/internal/repository"
)

type PostsHandler struct {
	l   *log.Logger
	rep *repository.PostRepository
}

func NewPostHandler(l *log.Logger, r *repository.PostRepository) *PostsHandler {
	return &PostsHandler{l, r}
}

var ErrInternalError = errors.New("internal error")
var ErrBadRequest = errors.New("bad request")
