package repository

import (
	"context"

	"github.com/NeGat1FF/blog-api/internal/models"
	"github.com/uptrace/bun"
)

type PostRepository struct {
	db *bun.DB
}

func NewPostRepository(db *bun.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) GetAllPosts(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.NewSelect().Model(&posts).Scan(ctx)
	return posts, err
}

func (r *PostRepository) GetPostByID(ctx context.Context, id int) (models.Post, error) {
	var post models.Post
	err := r.db.NewSelect().Model(&post).Where("? = ?", bun.Ident("id"), id).Scan(ctx)
	return post, err
}

func (r *PostRepository) GetPostsByTerm(ctx context.Context, term string) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.NewSelect().Model(&posts).Where("?0 ILIKE ?3 OR ?1 ILIKE ?3 OR ?2 ILIKE ?3", bun.Ident("title"), bun.Ident("content"), bun.Ident("category"), "%"+term+"%").Scan(ctx)
	return posts, err
}

func (r *PostRepository) UpdatePost(ctx context.Context, post models.Post) error {
	_, err := r.db.NewUpdate().Model(&post).OmitZero().WherePK().Exec(ctx)
	return err
}

func (r *PostRepository) DeletePost(ctx context.Context, id int) error {
	post := models.Post{ID: id}
	_, err := r.db.NewDelete().Model(&post).WherePK().Exec(ctx)
	return err
}

func (r *PostRepository) AddPost(ctx context.Context, post models.Post) (int, error) {
	var id int
	err := r.db.NewInsert().Model(&post).Returning("id").Scan(ctx, &id)
	return id, err
}
