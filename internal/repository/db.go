package repository

import (
	"database/sql"
	"time"

	"github.com/NeGat1FF/blog-api/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func InitDB(cfg *config.Config) *bun.DB {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(cfg.DB.Host+":"+cfg.DB.Port), // Load from config
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithUser(cfg.DB.User),         // Load from config
		pgdriver.WithPassword(cfg.DB.Password), // Load from config
		pgdriver.WithDatabase(cfg.DB.Name),     // Load from config
		pgdriver.WithApplicationName("blog-api"),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	// Create a Bun database instance
	db := bun.NewDB(sql.OpenDB(pgconn), pgdialect.New())

	// Return the database instance
	return db

}
