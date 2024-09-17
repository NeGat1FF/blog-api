package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"table:post" swaggerignore:"true"`

	ID         int       `bun:"id,pk,autoincrement" json:"id"`
	Title      string    `bun:"title,notnull" json:"title"`
	Tags       []string  `bun:"tags,array" json:"tags"`
	Category   string    `bun:"category" json:"category"`
	Content    string    `bun:"content,notnull" json:"content"`
	Created_at time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at,omitempty"`
	Updated_at time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at,omitempty"`
}
